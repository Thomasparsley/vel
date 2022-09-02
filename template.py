from datetime import datetime
from os import PathLike
from typing import Any, Mapping

from starlette.types import Receive, Scope, Send
from starlette.background import BackgroundTask
from starlette.responses import Response
from fastapi import Request

import minify_html
import jinja2

Context = dict[str, Any]


class _TemplateResponse(Response):
    media_type = "text/html"

    def __init__(
        self,
        template: jinja2.Template,
        context: Context,
        status_code: int = 200,
        headers: Mapping[str, str] | None = None,
        media_type: str | None = None,
        background: BackgroundTask | None = None,
    ):
        self.template = template
        self.context = context
        content = minify_html.minify(template.render(context))

        super().__init__(content, status_code, headers, media_type, background)

    async def __call__(self, scope: Scope, receive: Receive, send: Send) -> None:
        request = self.context.get("request", {})
        extensions = request.get("extensions", {})  # type: ignore
        if "http.response.template" in extensions:
            await send(
                {
                    "type": "http.response.template",
                    "template": self.template,
                    "context": self.context,
                }
            )
        await super().__call__(scope, receive, send)


class Jinja2Templating:
    def __init__(
        self,
        directory: str | PathLike,  # type: ignore
        **env_options: Any,
    ) -> None:
        assert jinja2 is not None, "jinja2 must be installed to use Jinja2Templates"
        self.env = self.__create_env(directory, **env_options)  # type: ignore

    def __create_env(
        self,
        directory: str | PathLike,  # type: ignore
        **env_options: Any,
    ) -> jinja2.Environment:
        @jinja2.pass_context
        def url_for(context: Context, name: str, **path_params: Any) -> str:
            request = context["request"]
            return request.url_for(name, **path_params)

        @jinja2.pass_context
        def datetime_now(context: Context):
            return datetime.now()

        loader = jinja2.FileSystemLoader(directory)
        env_options.setdefault("loader", loader)
        env_options.setdefault("autoescape", True)

        env = jinja2.Environment(**env_options)
        env.globals["url_for"] = url_for  # type: ignore
        env.globals["datetime"] = datetime_now  # type: ignore

        return env

    def get_template(self, name: str) -> jinja2.Template:
        return self.env.get_template(name)

    def RenderResponse(
        self,
        request: Request,
        name: str,
        context: Context,
        status_code: int = 200,
        headers: Mapping[str, str] | None = None,
        media_type: str | None = None,
        background: BackgroundTask | None = None,
    ) -> _TemplateResponse:
        context["request"] = request

        return _TemplateResponse(
            self.get_template(name),
            context,
            status_code=status_code,
            headers=headers,
            media_type=media_type,
            background=background,
        )

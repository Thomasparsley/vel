from typing import Any
from hashlib import sha256
from functools import wraps
from datetime import timedelta
from fastapi import BackgroundTasks

from fastapi.responses import HTMLResponse

from .redis import Redis
from .config import Config
from .request import Request
from .exceptions import InvalidAuthenticationError


def calculate_html_key(request: Request[Any]):
    path = request.url.path
    dark_mode = True if "dark_mode" in request.cookies else False

    key = f"html-{path}-{dark_mode}-{request.is_htmx_boosted}"
    key = sha256(bytes(key, "utf-8")).hexdigest()

    return key


async def set_html_cache(key: str, content: bytes, exp: timedelta):
    await Redis.set(key, str(content, "utf-8"), exp)


def html_cache(
    exp: timedelta = timedelta(minutes=15),
    skip_user: bool = False,
):

    DEV: bool = Config.get("DEV")

    def inner(func: Any):
        @wraps(func)
        async def wrapper(request: Request[Any], *args: Any, **kwargs: Any):

            key = calculate_html_key(request)

            if not DEV or (skip_user and request.user):
                cached_content = await Redis.get(key)
                if cached_content:
                    return HTMLResponse(cached_content)

            kwargs["request"] = request
            response: HTMLResponse = await func(*args, **kwargs)

            if response.status_code >= 400:
                return response

            tasks = BackgroundTasks(
                [response.background] if response.background else []
            )
            tasks.add_task(set_html_cache, key, response.body, exp)
            response.background = tasks

            return response

        return wrapper

    return inner


def login_required(func: Any):
    @wraps(func)
    async def wrapper(request: Request[Any], *args: Any, **kwargs: Any):
        if not request.user:
            raise InvalidAuthenticationError(request.url)

        kwargs["request"] = request
        return await func(*args, **kwargs)

    return wrapper

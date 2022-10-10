from typing import Generic, TypeVar

from fastapi import Request as _Request
from fastapi.routing import APIRoute

from .htmx import is_boosted


T = TypeVar("T")


class Request(_Request, Generic[T]):
    @property
    def user(self) -> T | None:
        if "user" not in self.scope:
            return None

        return self.scope["user"]

    @property
    def is_htmx_boosted(self) -> bool:
        return is_boosted(self)

    @classmethod
    def maker_router(cls):
        class RequestRoute(APIRoute):
            def get_route_handler(self):
                original_route_handler = super().get_route_handler()

                async def custom_route_handler(request: _Request):
                    request = cls(request.scope, request.receive)
                    return await original_route_handler(request)

                return custom_route_handler

        return RequestRoute

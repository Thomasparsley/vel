from typing import Any

from fastapi import Request

from ..config import Config
from ..identity import process_jwt_token


class AddUserIntoRequestMiddleware:
    async def __call__(self, request: Request, call_next: Any):
        new_token = None

        try:
            token = request.cookies[Config.get("JWT_TOKEN_NAME")]
            user, new_token = await process_jwt_token(token)
            request.scope["user"] = user
        except KeyError:
            request.scope["user"] = None

        response = await call_next(request)

        if new_token:
            response.set_cookie(Config.get("JWT_TOKEN_NAME"), new_token)

        return response

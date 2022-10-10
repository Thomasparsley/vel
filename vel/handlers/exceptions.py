from fastapi import Request
from fastapi.responses import RedirectResponse

from ..config import Config
from ..exceptions import InvalidAuthenticationError


login_path: str = Config.get("LOGIN_PATH")


async def invalid_authentication_handler(
    request: Request, exc: InvalidAuthenticationError
):
    return RedirectResponse(
        url=f"{request.base_url}{login_path}?after={exc.current_path}",
        status_code=exc.status_code,
    )

from typing import Final

from fastapi import Request


HX_RETARGET: Final = "HX-Retarget"
HX_RETARGET_MAIN: Final = {HX_RETARGET: "main"}
HX_REDIRECT: Final = "HX-Redirect"
HX_REFRESH: Final = "HX-Refresh"
HX_REFRESH_HEADER: Final = {HX_REFRESH: "true"}


def HX_REDIRECT_HEADER(url: str):
    return {HX_REDIRECT: url}


def is_boosted(request: Request) -> bool:
    headers = request.headers

    if "hx-boosted" in headers:
        return headers["hx-boosted"] == "true"

    return False

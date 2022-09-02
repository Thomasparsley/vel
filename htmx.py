from typing import Final


HX_REDIRECT: Final = "HX-Redirect"
HX_REFRESH: Final = "HX-Refresh"
HX_REFRESH_HEADER: Final = {HX_REFRESH: "true"}


def HX_REDIRECT_HEADER(url: str):
    return {HX_REDIRECT: url}

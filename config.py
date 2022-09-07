from typing import Any


class Config:
    dev: bool = False

    allow_origins: list[str] = []
    allow_credentials: bool = False
    allow_methods: list[str] = []
    allow_headers: list[str] = []

    secret_key: str = ""
    hashids_salt: str = ""

    database: Any = None
    database_models: dict[str, list[str]] = {}

    login_path: str = ""
    user_class: Any = None
    jwt_token_name: str = ""

from typing import Any
from types import ModuleType
from importlib import import_module


class Config:
    _module: ModuleType | None = None

    @staticmethod
    def _load_module():
        """path = os.environ.get("VEL_CONFIG")
        if not path:
            raise ValueError"""

        Config._module = import_module("web.config")

    @staticmethod
    def get(attribute_name: str) -> Any:
        if not Config._module:
            Config._load_module()

        return getattr(Config._module, attribute_name)

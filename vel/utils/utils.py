import importlib
from typing import Any


def get_attribute_from_module(module_path: str, attribute_name: str) -> Any:
    config_module = importlib.import_module(module_path)
    return getattr(config_module, attribute_name)

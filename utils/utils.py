import pathlib
import importlib
from typing import Any


def exist_file(path: pathlib.Path):
    return path.is_file()


def get_attribute_from_module(module_path: str, attribute_name: str) -> Any:
    config_module = importlib.import_module(module_path)
    return getattr(config_module, attribute_name)

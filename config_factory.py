import os
import importlib
from types import ModuleType

from . import constants
from .utils import Singleton


class ConfigFactory(Singleton):
    __mod = None

    def get(self) -> ModuleType:
        if self.__mod is not None:
            return self.__mod

        module_path = os.environ.get(constants.ENVIRONMENT_VARIABLE)
        if module_path is None:
            raise ValueError

        self.__mod = importlib.import_module(module_path)
        return self.get()

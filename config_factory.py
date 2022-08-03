import os
import importlib

from . import constants


class ConfigFactory:
    __MOD = None

    @staticmethod
    def get():
        if ConfigFactory.__MOD is not None:
            return ConfigFactory.__MOD

        module_path = os.environ.get(constants.ENVIRONMENT_VARIABLE)
        if module_path is None:
            raise ValueError

        ConfigFactory.__MOD = importlib.import_module(module_path)
        return ConfigFactory.__MOD

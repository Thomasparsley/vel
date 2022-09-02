import pathlib


def exist_file(path: pathlib.Path):
    return path.is_file()

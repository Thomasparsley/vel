import hashlib
import pathlib


def exist_file(path: pathlib.Path):
    return path.is_file()


def calculate_file_hash(file_path: pathlib.Path):
    with open(file_path, "rb") as file:
        bytes = file.read()
        return hashlib.sha256(bytes).hexdigest()

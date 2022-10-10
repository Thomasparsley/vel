from fastapi import HTTPException
from starlette.datastructures import URL


class NoFilesUploadedException(HTTPException):
    def __init__(self):
        super().__init__(status_code=400, detail="No files uploaded")


class FileNotFoundException(HTTPException):
    def __init__(self):
        super().__init__(status_code=404, detail="File not found")


class FileDoesNotExistException(HTTPException):
    def __init__(self):
        super().__init__(status_code=404, detail="File does not exist")


class InvalidAuthenticationError(HTTPException):
    def __init__(self, current_path: URL):
        super().__init__(302)
        self.current_path = current_path.path

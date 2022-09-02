from fastapi import HTTPException

from vel.config_factory import ConfigFactory


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
    def __init__(self):
        login_path: str = ConfigFactory().get().LOGIN_PATH
        super().__init__(
            status_code=302,
            detail="Not authorized",
            headers={"Location": login_path},
        )

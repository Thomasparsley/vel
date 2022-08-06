from fastapi import HTTPException


class NoFilesUploadedException(HTTPException):
    def __init__(self):
        super().__init__(status_code=400, detail="No files uploaded")


class FileNotFoundException(HTTPException):
    def __init__(self):
        super().__init__(status_code=404, detail="File not found")


class FileDoesNotExistException(HTTPException):
    def __init__(self):
        super().__init__(status_code=404, detail="File does not exist")

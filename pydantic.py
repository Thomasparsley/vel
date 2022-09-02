from fastapi import Request
from pydantic import BaseModel as _BaseModel


class BaseModel(_BaseModel):
    @classmethod
    async def from_form(cls, request: Request):
        form = await request.form()
        return cls(**form)

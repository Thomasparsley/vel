from fastapi import Request
from pydantic import BaseModel as __BaseModel
from pydantic import ValidationError as _ValidationError

from .exceptions_pydantic import ValidationError


class BaseModel(__BaseModel):
    @classmethod
    async def from_form(cls, request: Request):
        form = await request.form()

        try:
            return cls(**form)
        except _ValidationError as e:
            raise ValidationError(e.raw_errors, e.model) from e

    def keys(self):
        keys: list[str] = []

        for key, value in self.dict():
            if value is not None:
                keys.append(key)

        return keys

from fastapi import Request
from pydantic import BaseModel as __BaseModel
from pydantic import ValidationError as _ValidationError


class ValidationError(_ValidationError):
    def to_dict(self):
        result: dict[str, str] = dict()

        for error in self.errors():
            location = str(error["loc"][0])
            result[location] = error["msg"]

        return result


class BaseModel(__BaseModel):
    @classmethod
    async def from_form(cls, request: Request):
        form = await request.form()

        try:
            return cls(**form)
        except _ValidationError as e:
            raise ValidationError(e.raw_errors, e.model) from e

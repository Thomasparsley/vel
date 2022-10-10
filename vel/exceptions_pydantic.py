from pydantic import ValidationError as _ValidationError


class ValidationError(_ValidationError):
    def to_dict(self):
        result: dict[str, str] = dict()

        for error in self.errors():
            location = str(error["loc"][0])
            result[location] = error["msg"]

        return result

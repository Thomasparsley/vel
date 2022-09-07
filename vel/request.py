from typing import Generic, TypeVar

from fastapi import Request as _Request


T = TypeVar("T")


class Request(_Request, Generic[T]):
    @property
    def user(self) -> T | None:
        if "user" not in self.scope:
            return None

        return self.scope["user"]

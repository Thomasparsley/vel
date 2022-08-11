from hashids import Hashids as __Hashids  # type: ignore

from .config_factory import ConfigFactory
from .singleton import Singleton


class HashidsSingleton(Singleton):
    __hashids = __Hashids(salt=ConfigFactory().get().HASHED_ID_SALT)

    def decode(self, ids: str) -> tuple[int]:
        return self.__hashids.decode(ids)  # type: ignore

    def decode_single(self, id: str) -> int | None:
        ids = self.decode(id)

        try:
            return ids[0]
        except IndexError:
            return None

    def encode(self, ids: list[int]) -> str:
        return self.__hashids.encode(ids)  # type: ignore


class HashidsMixin:
    id: int = 0
    __hashed_id: str | None = None

    @property
    def hashed_id(self) -> str:
        if self.__hashed_id is not None:
            return self.__hashed_id
        else:
            self.__hashed_id = HashidsSingleton().encode([self.id])
            return self.hashed_id

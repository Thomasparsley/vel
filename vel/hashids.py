from hashids import Hashids as _Hashids  # type: ignore

from .config import Config


class Hashids:
    hashid = _Hashids(salt=Config.hashids_salt, min_length=6)

    @staticmethod
    def encode(id: int) -> str:
        return Hashids.hashid.encode(  # type: ignore
            id,
        )

    @staticmethod
    def decode(ids: str) -> tuple[int]:
        return Hashids.hashid.decode(ids)  # type: ignore

    @staticmethod
    def decode_single(id: str) -> int | None:
        ids = Hashids.decode(id)

        try:
            return ids[0]
        except IndexError:
            return None


class HashidsMixin:
    id: int = 0
    _hashed_id: str | None = None

    @property
    def hashed_id(self) -> str:
        if self._hashed_id is not None:
            return self._hashed_id
        else:
            self._hashed_id = Hashids.encode(self.id)
            return self.hashed_id

from pytz import UTC
from typing import TypeVar, Generic
from datetime import timedelta, datetime


from .models.cache import CacheModel


T = TypeVar("T")


class Cache(Generic[T]):
    @staticmethod
    async def set(key: str, data: T, delta: timedelta):
        new = CacheModel(key=key, data=data, exp=datetime.utcnow() + delta)
        await new.save()

    @staticmethod
    async def get(key: str) -> T | None:
        cached = await CacheModel.get_or_none(key=key)

        if not cached:
            return None

        if UTC.localize(datetime.utcnow()) > cached.exp:
            await Cache.delete(key)
            return None

        data: T = cached.data  # type: ignore
        return data

    @staticmethod
    async def delete(key: str):
        data = await CacheModel.get_or_none(key=key)
        if data:
            await data.delete()

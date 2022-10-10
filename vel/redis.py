from datetime import timedelta

from redis import asyncio as aioredis


class Redis:
    @staticmethod
    def init(host: str, port: int):
        Redis._backend: aioredis.Redis[str] = aioredis.Redis(
            host=host,
            port=port,
        )

    @staticmethod
    async def set(key: str, data: str, delta: timedelta):
        await Redis._backend.set(name=key, value=data, ex=delta)

    @staticmethod
    async def get(key: str):
        return await Redis._backend.get(key)

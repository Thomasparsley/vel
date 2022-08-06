from argparse import ArgumentError
from dataclasses import dataclass
from typing import Optional
from enum import IntEnum

from tortoise import Tortoise


class DatabaseType(IntEnum):
    POSTGRE_SQL = 0


@dataclass
class Database:
    user: str
    password: str
    server_address: str
    database_name: str
    server_port: int
    database_type: DatabaseType
    __is_init: bool = False

    async def init(self, models: dict[str, list[str]]):
        connection_string = self.get_connection_str()

        await Tortoise.init(  # type: ignore
            db_url=connection_string,
            modules=models,  # type: ignore
        )

        self.__is_init = True

    @property
    def is_init(self):
        return self.__is_init

    async def generate_schemas(self, models: Optional[dict[str, list[str]]] = None):
        if not self.is_init:
            if models is None:
                raise ArgumentError(
                    models,
                    "Argument `models` cannot be None unless a database connection is established",
                )

            await self.init(models)

        if self.is_init:
            await Tortoise.generate_schemas()

    def get_connection_str(self) -> str:
        match self.database_type:
            case DatabaseType.POSTGRE_SQL:
                return self.__get_postgres_connection()

    def __get_postgres_connection(self) -> str:
        return f"postgres://{self.user}:{self.password}@{self.server_address}:{self.server_port}/{self.database_name}"

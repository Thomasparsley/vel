import enum
from typing import Final

from tortoise import fields


class Visibility(enum.IntEnum):
    PUBLIC = 1
    DRAFT = 2
    PRIVATE = 3
    DELETED = 4


VISIBILITY_FIELD: Final = fields.IntEnumField(Visibility)

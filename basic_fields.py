from typing import Final

from tortoise import fields

from .visibility import VISIBILITY_FIELD as __VISIBILITY_FIELD


VISIBILITY_FIELD = __VISIBILITY_FIELD

ID_FIELD: Final = fields.BigIntField(pk=True)

CREATED_AT_FIELD: Final = fields.DatetimeField(auto_now_add=True)
UPDATED_AT_FIELD: Final = fields.DatetimeField(auto_now=True)

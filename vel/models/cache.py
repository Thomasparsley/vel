from tortoise import fields, models

from .. import basic_fields


class CacheModel(models.Model):
    key = fields.CharField(256, pk=True)
    data = fields.JSONField()
    exp = fields.DatetimeField()
    created_at = basic_fields.CREATED_AT_FIELD
    updated_at = basic_fields.UPDATED_AT_FIELD

    class Meta:  # type: ignore
        table = "cache"

from tortoise import fields, models

from .. import basic_fields


class Page(models.Model):
    id = basic_fields.ID_FIELD
    visibility = basic_fields.VISIBILITY_FIELD
    title = fields.CharField(max_length=255, null=False)
    slug = fields.CharField(max_length=255, null=False, unique=True, index=True)
    body = fields.TextField()

    created_at = basic_fields.CREATED_AT_FIELD
    updated_at = basic_fields.UPDATED_AT_FIELD

    class Meta:  # type: ignore
        abstract = True

    def render(self) -> str:
        return ""

from tortoise import fields, models
from tortoise.exceptions import DoesNotExist

from .. import basic_fields
from ..visibility import Visibility
from ..hashids import HashidsMixin, Hashids


class FileModel(models.Model, HashidsMixin):
    id = basic_fields.ID_FIELD
    visibility = basic_fields.VISIBILITY_FIELD
    filename = fields.CharField(max_length=2048, null=False, index=True)
    size = fields.IntField()
    content_type = fields.CharField(max_length=256, null=False)

    class Meta:  # type: ignore
        abstract = True

    def __init__(self, size: int, filename: str, content_type: str):
        super().__init__()
        self.visibility = Visibility.PRIVATE

        self.size = size
        self.filename = filename
        self.content_type = content_type

    @classmethod
    async def get_all(cls):
        return await cls.all().order_by("created_at")

    @classmethod
    async def get_by_hashed_id(cls, hashed_id: str):
        id = Hashids().decode_single(hashed_id)
        if not id:
            return None

        try:
            return await cls.get(id=id)
        except DoesNotExist:
            return None

    def is_image(self) -> bool:
        return "image/" in self.content_type

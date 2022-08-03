import pathlib
from typing import Final

from tortoise import fields, models
from tortoise.exceptions import DoesNotExist

from vel.hashids import HashidsMixin, hashids

from . import basic_fields


IMAGES_EXT: Final = [
    "jpg",  # JPEG data in JFIF or Exif format
    "png",  # Portable Network Graphics
    "gif",  # GIF 87a and 89a File
    "webp",  # WebP file
    "rgb",  # SGI ImgLib File
    "pbm",  # Portable Bitmap File
    "pgm",  # Portable Graymap File
    "ppm",  # Portable Pixmap File
    "tiff",  # TIFF File
    "rast",  # Sun Raster File
    "xbm",  # X Bitmap File
    "bmp",  # BMP file
    "exr",  # OpenEXR File
]


class File(models.Model, HashidsMixin):
    id = basic_fields.ID_FIELD
    visibility = basic_fields.VISIBILITY_FIELD
    filename = fields.CharField(max_length=2048, null=False, index=True)
    size = fields.IntField()
    content_type = fields.CharField(max_length=256, null=False)
    uploaded_by = fields.ForeignKeyField("vel.User", null=True)

    created_at = basic_fields.CREATED_AT_FIELD
    updated_at = basic_fields.UPDATED_AT_FIELD

    class Meta:
        table = "vel_files"

    @classmethod
    async def get_by_hashed_id(cls, hashed_id: str):
        try:
            id: tuple = hashids.decode(hashed_id)[0]
        except IndexError:
            return None

        try:
            return await cls.get(id=id)
        except DoesNotExist:
            return None

    def is_image(self) -> bool:
        return "image/" in self.content_type

def exist_file(path: pathlib.Path):
    return path.is_file()
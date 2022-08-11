import pathlib
from enum import Enum
from typing import Final


from tortoise import fields, models
from tortoise.exceptions import DoesNotExist

from vel.hashids import HashidsSingleton, HashidsMixin
from vel.visibility import Visibility

from .. import basic_fields
from .user import User


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


class File(HashidsMixin, models.Model):
    id = basic_fields.ID_FIELD
    visibility = basic_fields.VISIBILITY_FIELD
    filename = fields.CharField(max_length=2048, null=False, index=True)
    size = fields.IntField()
    content_type = fields.CharField(max_length=256, null=False)
    uploaded_by: fields.ForeignKeyRelation[User] = fields.ForeignKeyField(
        "vel.User", null=True
    )

    created_at = basic_fields.CREATED_AT_FIELD
    updated_at = basic_fields.UPDATED_AT_FIELD

    class Meta:  # type: ignore
        table = "vel__files"

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
        id = HashidsSingleton().decode_single(hashed_id)
        if not id:
            return None

        try:
            return await cls.get(id=id)
        except DoesNotExist:
            return None

    def is_image(self) -> bool:
        return "image/" in self.content_type


class ImageSize(str, Enum):
    DEFAULT = ""
    XS = "xs"
    S = "s"
    M = "m"
    L = "l"
    XL = "xl"
    XXL = "xxl"

    def get_px(self) -> int:
        match self:
            case ImageSize.XS:
                return 100

            case ImageSize.S:
                return 200

            case ImageSize.M:
                return 400

            case ImageSize.L:
                return 800

            case ImageSize.XL:
                return 1200

            case ImageSize.XXL:
                return 1600

            case _:
                return 0


class ImageType(str, Enum):
    DEFAULT = ""
    WEBP = "webp"


def exist_file(path: pathlib.Path):
    return path.is_file()

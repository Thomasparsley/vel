from typing import Final
from enum import Enum


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

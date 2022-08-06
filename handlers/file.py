import pathlib
import aiofiles

from PIL import Image
from fastapi import UploadFile
from fastapi.responses import FileResponse

from ..models.file import (
    File,
    ImageSize,
    ImageType,
    exist_file,
)


def image_path(file: File, size: ImageSize, type: ImageType):
    match (size, type):
        case (ImageSize.DEFAULT, ImageType.DEFAULT):
            return pathlib.Path(f"./files/{file.filename}")

        case (size, ImageType.DEFAULT):
            return pathlib.Path(f"./files/{file.filename}-{size}")

        case (ImageSize.DEFAULT, type):
            return pathlib.Path(f"./files/{file.filename}-{type}")

        case _:
            return pathlib.Path(f"./files/{file.filename}-{size}-{type}")


async def generate_image(
    file: File,
    save_to: pathlib.Path,
    size: ImageSize,
    type: ImageType,
):
    image = Image.open(pathlib.Path(f"./files/{file.filename}"))

    px = size.get_px()
    if px:
        image.thumbnail((px, px))

    match type:
        case ImageType.DEFAULT:
            image.save(save_to)

        case ImageType.WEBP:
            image.save(save_to, format="webp")


async def image_handler(file: File, size: ImageSize, type: ImageType):
    file_path = image_path(file, size, type)

    if not exist_file(file_path):
        await generate_image(file, file_path, size, type)

    return FileResponse(file_path)


async def uploaded_file_handler(file: UploadFile):
    content = await file.read()
    db_file = File(len(content), file.filename, file.content_type)
    await db_file.save()

    async with aiofiles.open(f"./files/{file.filename}", "wb") as writer:
        await writer.write(content)

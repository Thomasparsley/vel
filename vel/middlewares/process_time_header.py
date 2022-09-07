from time import time
from typing import Any

from fastapi import Request


class ProcessTimeHeaderMiddleware:
    async def __call__(self, request: Request, call_next: Any):
        start_time = time()
        response = await call_next(request)
        process_time = time() - start_time

        if process_time <= 0.009:
            process_time = f"{process_time/1000}ms"

        response.headers["X-Process-Time"] = str(process_time)
        return response

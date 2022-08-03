import time

from fastapi import Request


class ProcessTimeHeader:
    async def __call__(self, request: Request, call_next):
        start_time = time.time()
        response = await call_next(request)
        process_time = time.time() - start_time

        if process_time <= 0.009:
            process_time = f"{process_time/1000}ms"

        response.headers["X-Process-Time"] = str(process_time)
        return response

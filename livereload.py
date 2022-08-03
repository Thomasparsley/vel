import asyncio
import pathlib
from typing import Any
from datetime import datetime

from websockets.exceptions import ConnectionClosed
from fastapi import FastAPI, WebSocket, WebSocketDisconnect
from watchdog.observers import Observer  # type: ignore
from watchdog.events import LoggingEventHandler  # type: ignore


class LiveReload:
    class WSHub:
        def __init__(self):
            self.connections: list[WebSocket] = []

        async def connect(self, connection: WebSocket):
            await connection.accept()
            self.connections.append(connection)

        def disconnect(self, connection: WebSocket):
            self.connections.remove(connection)

        async def broadcast_json(self, json_message: dict[str, Any]):
            for connection in self.connections:
                await connection.send_json(json_message)

        async def send_personal_json(
            self, websocket: WebSocket, json_message: dict[str, Any]
        ):
            await websocket.send_json(json_message)

    class EventHandler(LoggingEventHandler):
        def __init__(self, manager: "LiveReload.WSHub", logger: Any = None):
            super().__init__(logger)

            self.manager = manager

        def on_modified(self, event: Any):
            super().on_modified(event)  # type: ignore
            asyncio.run(self.broadcast_reload())

        async def broadcast_reload(self):
            return await self.manager.broadcast_json(
                {"command": "reload", "time": datetime.now().isoformat()}
            )

    def __init__(self, path: pathlib.Path):
        self.__manager = LiveReload.WSHub()
        self.__event_handler = LiveReload.EventHandler(self.__manager)
        self.__observer = Observer()
        self.__observer.schedule(self.__event_handler, path, recursive=True)  # type: ignore

    def add_endpoint(self, app: FastAPI, path: str):
        @app.websocket(path)
        async def livereload_endpoint(connection: WebSocket):  # type: ignore
            await self.__manager.connect(connection)

            try:
                while True:
                    await self.__manager.send_personal_json(
                        connection, {"ping": datetime.now().isoformat()}
                    )
                    await asyncio.sleep(3)
            except WebSocketDisconnect:
                pass
            except ConnectionClosed:
                pass
            finally:
                self.__manager.disconnect(connection)

    def start(self):
        self.__observer.start()

    def stop(self):
        self.__observer.stop()
        self.__observer.join()

    def jinja_filter(self, addr: str):
        return f"""
            <script>
            let fullReaload = false;

            const livereload = (path) => {{
                let socket = new WebSocket(path);

                socket.onopen = () => {{
                    console.log("Live reload connected");

                    if (fullReaload) {{
                        location.reload();
                    }}
                }};

                socket.onclose = () => {{
                    console.log("Trying to reconnect to live reload");

                    setTimeout(() => {{
                        livereload(path);
                    }}, 1000);
                }};

                socket.onerror = (err) => {{
                    console.error("Live reload error: ", err);
                    fullReaload = true;
                }};

                socket.onmessage = (message) => {{
                    const data = JSON.parse(message.data);

                    if (data.command === "reload") {{
                        location.reload();
                    }}
                }};
            }};

            livereload("{addr}");
            </script>
            """

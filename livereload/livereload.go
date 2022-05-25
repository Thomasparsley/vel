package livereload

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"github.com/Thomasparsley/vel/ws"
)

type Livereload struct {
	watcher   *fsnotify.Watcher
	directory string
	done      chan bool
}

func New(directory string) (*Livereload, error) {
	watcherm, err := fsnotify.NewWatcher()

	return &Livereload{
		watcher:   watcherm,
		directory: directory,
		done:      make(chan bool),
	}, err
}

func (l *Livereload) Watch() error {
	err := filepath.Walk(l.directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsDir() {
			return l.watcher.Add(path)
		}

		return nil
	})

	go l.hub(ws.NewHub())

	return err
}

func (l *Livereload) hub(hub *ws.Hub) {
	for {
		select {
		// Register new client
		case client := <-hub.ChanRegister():
			hub.AddClient(client)

		// Unregister client
		case connection := <-hub.ChanUnregister():
			if has := hub.HasClient(connection); has {
				hub.RemoveClient(connection)
			}

		case event, ok := <-l.watcher.Events:
			if !ok {
				continue
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)

				hub.SendBroadcast(ws.NewJsonMessage(
					map[string]any{
						"command": "reload",
						"time":    time.Now().Format(time.RFC3339),
					},
					"*",
				))
			}

		case err, ok := <-l.watcher.Errors:
			if !ok {
				continue
			}

			log.Println("error:", err)

		case <-l.done:
			return
		}
	}
}

func (l *Livereload) HttpControlLer(app *fiber.App, path string) {
	hub := ws.NewHub()

	app.Use(path, func(c *fiber.Ctx) error {
		// Returns true if the client requested upgrade to the WebSocket protocol
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}

		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Get(path, websocket.New(func(c *websocket.Conn) {
		defer func() {
			hub.SendUnregister(c)
			c.Close()
		}()

		client := ws.NewClient(c, "*")
		hub.SendRegister(client)

		for {
			err := client.SendJSON(map[string]any{
				"command": "ping",
				"time":    time.Now().Format(time.RFC3339),
			})
			if err != nil {
				return
			}

			time.Sleep(time.Second * 5)
		}
	}))
}

func (l *Livereload) Close() error {
	return l.watcher.Close()
}

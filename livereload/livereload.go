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
	hub       *ws.Hub
}

func New(directory string) (*Livereload, error) {
	watcherm, err := fsnotify.NewWatcher()

	return &Livereload{
		watcher:   watcherm,
		directory: directory,
		done:      make(chan bool),
		hub:       ws.NewHub(),
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

	go l.hubHandler()

	return err
}

func (l *Livereload) hubHandler() {
	for {
		select {
		// Register new client
		case client := <-l.hub.ChanRegister():
			l.hub.AddClient(client)

		// Unregister client
		case connection := <-l.hub.ChanUnregister():
			if has := l.hub.HasClient(connection); has {
				l.hub.RemoveClient(connection)
			}

		case event, ok := <-l.watcher.Events:
			if !ok {
				continue
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)

				l.hub.SendBroadcast(ws.NewJsonMessage(
					fiber.Map{
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

func (l *Livereload) HttpController(app *fiber.App, path string) {

	app.Use(path, ws.AutoUpgrade)
	app.Get(path, websocket.New(func(c *websocket.Conn) {
		defer func() {
			l.hub.SendUnregister(c)
			c.Close()
		}()

		client := ws.NewClient(c, "*")
		l.hub.SendRegister(client)

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

package src

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Application struct {
	up websocket.Upgrader
}

func NewApplication() *Application {
	var up = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &Application{
		up: up,
	}
}

func (a *Application) echo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	c, err := a.up.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (a *Application) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (a *Application) Run() error {
	http.HandleFunc("/echo", a.echo)
	http.HandleFunc("/", a.test)
	http.ListenAndServe("0.0.0.0:8080", nil)
	return nil
}

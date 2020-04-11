package router

import (
	// "../middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"html/template"
	"fmt"
)

var homeTemplate = template.Must(template.ParseFiles("index.html"))
var upgrader = websocket.Upgrader{
	// TO-DO: actually check origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
  	fmt.Println("Failed to set websocket upgrade: %+v", err)
    return
  }
	defer c.Close()

  for {
  	t, msg, err := c.ReadMessage()
  	if err != nil {
    	break
    }
    c.WriteMessage(t, msg)
  }
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/ws", echo)
	// router.HandleFunc("/api/movement", middleware.Movement).Methods("GET", "OPTIONS")

	return router
}

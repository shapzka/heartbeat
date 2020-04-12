package main

import (
	"fmt"
	/* Uncomment for Recent WebSocket
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	*/
	"net/http"
    "log"
    "./router"
)
/*
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
*/
func main() {
/* Uncomment for Websocket
  	r := gin.Default()
  	r.LoadHTMLFiles("index.html")

  	r.GET("/", func(c *gin.Context) {
      c.HTML(200, "index.html", nil)
  	})
  	r.GET("/ping", func(c *gin.Context) {
  		c.JSON(200, gin.H{
  			"message": "pong",
  		})
  	})
  	r.GET("/ws", func(c *gin.Context) {
  		echo(c.Writer, c.Request)
  	})
  	r.Run() // listen and serve on 0.0.0.0:8080
      */
    r := router.Router()
    // fs := http.FileServer(http.Dir("build"))
    // http.Handle("/", fs)
    fmt.Println("Starting server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))

}

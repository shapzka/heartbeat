package main

import (
	"fmt"
	"net/http"
  "log"
  "./router"
)

func main() {
    r := router.Router()

    fmt.Println("Starting Go server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))

}

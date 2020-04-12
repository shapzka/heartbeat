package main

import (
	"fmt"
	"net/http"
	"log"
	"flag"

	"./router"
	"./middleware"
)

func main() {
		var dbAddr string
		flag.StringVar(&dbAddr, "db", "mongodb://localhost:27017", "MongoDB instance address")
		flag.Parse()

		middleware.Init(dbAddr)

    r := router.Router()

    fmt.Println("Starting Go server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))

}

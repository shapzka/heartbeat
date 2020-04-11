package router

import (
	"../middleware"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

var homeTemplate = template.Must(template.ParseFiles("index.html"))

// Serve the basic HTML
func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/api/websocket")
}

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/api/movements", middleware.Movements).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/movements/{id}/vote", middleware.Vote).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/api/websocket", middleware.RegisterConnection)

	go middleware.Broadcast()

	return router
}

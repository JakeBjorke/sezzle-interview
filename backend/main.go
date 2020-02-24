package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jakebjorke/sezzle-interview/endpoints"
	"github.com/jakebjorke/sezzle-interview/websockets"
)

func main() {
	pool := websockets.NewPool()
	go pool.Start()

	endpoints.SetBroadcastChan(pool.Broadcast)

	router := mux.NewRouter()

	router.HandleFunc("/expression/", endpoints.Expression).Methods("POST")
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		endpoints.ServeWebSockets(pool, w, r)
	}).Methods("GET")

	port := ":8080"

	log.Println("Listening and serving on port ", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"}),
		handlers.AllowedMethods([]string{"POST", "GET"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}

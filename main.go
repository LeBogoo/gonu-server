package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"

	"gonu-server/onu"
)

// define the upgrade websocket function
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var PORT = flag.Int("port", 3000, "port to run the server on")

func main() {
	flag.Parse()

	// create a new router
	router := http.NewServeMux()

	// add CORS middleware to the router
	corsHandler := handlers.CORS()(router)

	// define the handler function for the /ping route
	pingHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}

	// add CORS middleware to the pingHandler function
	pingCorsHandler := handlers.CORS()(http.HandlerFunc(pingHandler))

	games := make(map[string]*onu.Game)

	// define the WebSocket handler for the root route
	wsHandler := func(w http.ResponseWriter, r *http.Request) {
		// upgrade the HTTP connection to a WebSocket connection
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			// log.Println(err)
			return
		}

		// create a new player
		onu.NewPlayer(conn, &games)

	}

	// add the routes to the router
	router.HandleFunc("/", wsHandler)
	router.Handle("/ping", pingCorsHandler)

	log.Printf("Server started on port %d", *PORT)
	// start the server with the CORS middleware
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *PORT), corsHandler))

}

// Command server runs the Shrink! backend API.
package main

import (
	"log"
	"net/http"

	"github.com/yaremam/shrink/backend/internal/httpserver"
)

func main() {
	const addr = ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, httpserver.New()))
}

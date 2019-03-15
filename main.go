package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := "127.0.0.1:8080"
	log.Println(`listening on port` + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil), `listeing on port:`, PORT)
}

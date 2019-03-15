package main

import (
	"log"
	"net/http"
)

func main() {
	// create a port for the server to connect
	PORT := "127.0.0.1:8080"
	// prints the server is listening on port 8080
	log.Println(`listening on port` + PORT)
	// net/http method that handles servering statice content
	http.Handle("/", http.FileServer(http.Dir("public")))
	// connects to  port 8080
	log.Fatal(http.ListenAndServe(PORT, nil))
}

//CompleteTaskFunc will get the taks id and send it to the client
// func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		id := r.URL.Path[len("/tasks"):]
// 		w.Write([]byte("Get the task " + id))
// 	}
// }

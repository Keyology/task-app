package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := "127.0.0.1:8080"
	log.Println(`listening on port` + PORT)
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(PORT, nil))
}

//CompleteTaskFunc will get the taks id and send it to the client
// func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		id := r.URL.Path[len("/tasks"):]
// 		w.Write([]byte("Get the task " + id))
// 	}
// }

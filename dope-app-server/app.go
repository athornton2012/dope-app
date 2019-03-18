package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func reviewHandler(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Path[len("/review/"):]
// 	a, _ := album.fetch(id)
// }

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/album/{id}", GetAlbum).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode( /*encode some stuff here!*/ )
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// GreetHandler returns the greetings along with browser user agent
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ua := r.UserAgent()
	fmt.Fprintf(w, "Greetings user of %s", ua)
}

// IndexHandler returns the REST API introduction message
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Simple REST API which greets users")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", IndexHandler).Methods("GET")

	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("/greet", GreetHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8383", router))

}

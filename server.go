package main

import (
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"strings"

	"github.com/gorilla/mux"
)

const port string = ":8080"

func getName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Fprintln(w, "Hello "+name+".")
}

type textHandler struct {
	h http.HandlerFunc
}

func (t textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	t.h(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "friend"
	}
	fmt.Fprintln(w, "Hello, "+name+"!")
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	name, err := ioutil.ReadAll(r.Body)
	stringName := string(name)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else if stringName == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		fmt.Fprintln(w, "Hello, "+strings.Trim(stringName, " ")+"!")
	}
}

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/hello/{name}", getName)
	// http.Handle("/", r)
	http.Handle("/hello", textHandler{bodyHandler})
	fmt.Println("Server running on " + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

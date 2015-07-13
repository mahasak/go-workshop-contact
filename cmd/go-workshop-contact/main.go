package main

import (
	"encoding/json"
	"net/http"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/contact/save", saveContact)
	mux.HandleFunc("/api/contact/get", getContact)
	http.ListenAndServe(":"+port, mux)

}

//Contact contact struct
type Contact struct {
	Tel   string
	Email string
}

var contacts = make(map[string]Contact)

func saveContact(w http.ResponseWriter, r *http.Request) {
	var res map[string]string
	//&contact --> Pass by reference
	json.NewDecoder(r.Body).Decode(&res)
	c := Contact{
		Tel:   res["tel"],
		Email: res["email"],
	}
	contacts[res["name"]] = c
}

func getContact(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	//c, found := contacts[name]
	c, _ := contacts[name]
	json.NewEncoder(w).Encode(c)
}

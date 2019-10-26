package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Animal struct {
	Id      int    `json:"id"`
	Species string `json:"species"`
	Sound   string `json:"sound"`
}

func jsonAnimal(w http.ResponseWriter, req *http.Request) {
	// what is w
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	sound := req.URL.Query()["sound"]
	animal := Animal{
		Id:      1,
		Species: "Cat",
		Sound:   sound[0],
	}
	// fmt.Println(&w)
	// fmt.Fprintf(w, "%v", req.Header)
	json.NewEncoder(w).Encode(animal)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	user := User{
		Id:    1,
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
		Phone: "000099999",
	}

	json.NewEncoder(w).Encode(user)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/higo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/json", jsonHandler)

	http.HandleFunc("/json/animal", jsonAnimal)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

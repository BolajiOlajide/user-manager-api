package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["status"] = "success"
	response["message"] = "Hello World"

	fmt.Fprint(w, "Hello World")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go - User Manager Tutorial")

	handleRequests()
}

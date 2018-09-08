package main

import (
	"fmt"
	"net/http"
)

// AllUsers - method for getting all users
func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users Endpoint Hit")
}

// NewUser - method for creating a new user
func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New user endpoint hit")
}

// DeleteUser - method for deleting a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete endpoint hit")
}

// UpdateUser - method for updating a new user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update endpoint hit")
}

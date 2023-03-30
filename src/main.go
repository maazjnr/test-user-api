package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var userInfo []UserData

func runAllUserInfo() {
	users1 := UserData{
		Name:  "John",
		Email: "john@example.com",
		Age:   29,
	}

	userInfo = append(userInfo, users1)

	users2 := UserData{
		Name:  "Abdul",
		Email: "abdul@example.com",
		Age:   25,
	}

	userInfo = append(userInfo, users2)

	users3 := UserData{
		Name:  "Aminco",
		Email: "aminco@example.com",
		Age:   28,
	}

	userInfo = append(userInfo, users3)

	users4 := UserData{
		Name:  "Maaz",
		Email: "mz@example.com",
		Age:   21,
	}

	userInfo = append(userInfo, users4)
}

func getUserDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
}

func handleRoute() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUserDataHandler).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func main() {
	runAllUserInfo()
	handleRoute()
}

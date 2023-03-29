package main

import (
	"encoding/json"
	"fmt"
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

func mapArr() {

	var mapArrays = make(map[string]int)
	mapArrays["Tokyo event for flutter"] = 0
	mapArrays["London Conference event"] = 1
	mapArrays["Rwanda vacation"] = 3

	fmt.Printf("this the list of maps we have %v\n", mapArrays)

	var arraySlice = [5]int{1, 2, 3, 4, 5}
	var arraySlice2 = [5]int{6, 7, 8, 9, 10}

	slice1 := arraySlice[:4]
	slice2 := arraySlice2[:3]

	result := append(slice1, slice2...)

	fmt.Printf("Printing out %v ", result)
}

func main() {
	fmt.Println("Creating user details...")
	runAllUserInfo()
	handleRoute()
}

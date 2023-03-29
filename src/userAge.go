package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserAge struct {
	Name string
	Age  int
}

var usesAgeDis = []UserAge{
	{
		Name: "John",
		Age:  18,
	},
	{
		Name: "Jane",
		Age:  19,
	},
	{
		Name: "Mary",
		Age:  20,
	},
}

func AccessHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for _, user := range usesAgeDis {
		if user.Name == name {
			if user.Age >= 18 {
				fmt.Fprintf(w, "%s has access", name)
			} else {
				fmt.Fprintf(w, "%s does not have access", name)
			}
			return
		}
	}

	fmt.Fprintf(w, "User %s not found", name)
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/access/{name}", AccessHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

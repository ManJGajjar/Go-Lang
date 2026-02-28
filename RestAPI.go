package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var users []User

func RestAPIEx() {
    router := mux.NewRouter()
    router.HandleFunc("/users", getUsers).Methods("GET")
    router.HandleFunc("/users/{id}", getUser).Methods("GET")
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", router)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range users {
        if item.ID == params["id"] {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.NotFound(w, r)
}
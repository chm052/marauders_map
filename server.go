package main

import (
  "github.com/gorilla/mux"
  "net/http"
)

func main() {
  mx := mux.NewRouter()

  mx.HandleFunc("/", SayHelloWorld)
  mx.HandleFunc("/trucks", GetFoodTrucks)

  http.ListenAndServe(":8080", mx)
}

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
}

func GetFoodTrucks() {
  
}

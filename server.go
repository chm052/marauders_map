package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
)

func main() {
  mx := mux.NewRouter()

  mx.HandleFunc("/", SayHelloWorld)
  mx.HandleFunc("/trucks", GetFoodTrucks)
  mx.HandleFunc("/truck/open", OpenFoodTruck)
  mx.HandleFunc("/truck/close", OpenFoodTruck)

  fmt.Printf("Serving on port %i", 8080)
  http.ListenAndServe(":8080", mx)
}

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
}

func GetFoodTrucks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf=8")
  myItems := []string{"item1", "item2", "item3"}
  a, _:= json.Marshal(myItems)
  w.Write(a)
  return
}
func OpenFoodTruck(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Opening food truck!"))
}

func CloseFoodTruck(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Closing food truck :("))
}

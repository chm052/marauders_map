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
  mx.HandleFunc("/truck/close", CloseFoodTruck)

  fmt.Printf("Serving on port %i", 8080)
  http.ListenAndServe(":8080", mx)
}

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
}

func GetFoodTrucks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf=8")

  truck1 := Truck{1, "Greek Food Truck", "-41.292489", "174.778656"}
  truck2 := Truck{2, "Beat Kitchen", "-41.287022", "174.778667"}
  truck3 := Truck{3, "Nanny's Food Truck", "-41.290425", "174.779272"}
  myItems := []Truck{truck1, truck2, truck3}

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

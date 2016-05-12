package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
)

var truck1 = Truck{1, "Greek Food Truck", -41.292489, 174.778656}
var truck2 = Truck{2, "Beat Kitchen", -41.287022, 174.778667}
var truck3 = Truck{3, "Nanny's Food Truck", -41.290425, 174.779272}
var allTrucks = []Truck{truck1, truck2, truck3}
var db = initDb()

func main() {
  mx := mux.NewRouter()

  mx.HandleFunc("/", SayHelloWorld)
  mx.HandleFunc("/api/trucks/create", CreateFoodTruck).Methods("POST")
  mx.HandleFunc("/api/trucks/delete/{id}", DeleteFoodTruck).Methods("DELETE")
  mx.HandleFunc("/api/trucks/open/{id}", OpenFoodTruck).Methods("POST")
  mx.HandleFunc("/api/trucks/close/{id}", CloseFoodTruck).Methods("POST")
  mx.HandleFunc("/api/trucks/location/{id}", GetFoodTruckLocation).Methods("GET")
  mx.HandleFunc("/api/trucks/location/{id}", PostFoodTruckLocation).Methods("POST")
  mx.HandleFunc("/api/trucks/location", GetFoodTrucks).Methods("GET")
  mx.HandleFunc("/api/trucksdb", GetTrucksFromDb)

  fmt.Printf("Serving on port %i", 9001)
  http.ListenAndServe(":9001", mx)
}

func GetTrucksFromDb(w http.ResponseWriter, r *http.Request) {
  trucks := []Truck{}
  err := db.Select(&trucks, "SELECT id, name, lat, lng FROM FoodTrucks")
  fmt.Println(trucks)
  if (err != nil){
    fmt.Println(err)
    return
  }
  a, _:= json.Marshal(trucks)
  w.Write(a)
  return
}

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
}

func CreateFoodTruck(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Creating food truck!"))
}

func DeleteFoodTruck(w http.ResponseWriter, r *http.Request) {
  foodTruckId := mux.Vars(r)["id"]
  w.Write([]byte("Deleting food truck :( " + foodTruckId))
}

func OpenFoodTruck(w http.ResponseWriter, r *http.Request) {
  foodTruckId := mux.Vars(r)["id"]
  w.Write([]byte("Opening food truck! " + foodTruckId))
}

func CloseFoodTruck(w http.ResponseWriter, r *http.Request) {
  foodTruckId := mux.Vars(r)["id"]
  w.Write([]byte("Closing food truck :( " + foodTruckId))
}

func GetFoodTruckLocation(w http.ResponseWriter, r *http.Request) {
  foodTruckId := mux.Vars(r)["id"]
  w.Write([]byte("Here is food truck location! " + foodTruckId))
}

func PostFoodTruckLocation(w http.ResponseWriter, r *http.Request) {
  queryParameters := r.URL.Query()
  latitude := queryParameters.Get("lat")
  longitude := queryParameters.Get("lon")
  foodTruckId := mux.Vars(r)["id"]
  w.Write([]byte(fmt.Sprintf("Posting food truck location! %s at %s, %s",
                              foodTruckId, latitude, longitude)))
}

func GetFoodTrucks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf=8")
  a, _:= json.Marshal(allTrucks)
  w.Write(a)
  return
}

package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "strconv"
)

// var truck1  = Truck{1, "Greek Food Truck", 1, -41.292489, 174.778656, ""}
// var truck2 = Truck{2, "Beat Kitchen", 2, -41.287022, 174.778667, ""}
// var truck3 = Truck{3, "Nanny's Food Truck", 3, -41.290425, 174.779272, ""}
// var allTrucks = []Truck{truck1, truck2, truck3}
var db = initDb()

func main() {
  mx := mux.NewRouter()

  mx.HandleFunc("/api/trucks/create", CreateFoodTruck).Methods("POST")
  mx.HandleFunc("/api/trucks/delete/{id}", DeleteFoodTruck).Methods("DELETE")
  mx.HandleFunc("/api/trucks/open/{id}", OpenFoodTruck).Methods("POST")
  mx.HandleFunc("/api/trucks/close/{id}", CloseFoodTruck).Methods("POST")
  mx.HandleFunc("/api/trucks/location/{id}", GetFoodTruckLocation).Methods("GET")
  mx.HandleFunc("/api/trucks/location/{id}", PostFoodTruckLocation).Methods("POST")
  mx.HandleFunc("/api/trucks/location/all", GetFoodTrucks).Methods("GET")
  mx.HandleFunc("/api/trucks/test", GetTestFoodTrucks).Methods("GET")

  fmt.Printf("Serving on port 9001")
  http.ListenAndServe(":9001", mx)
}

func GetFoodTrucks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf=8")
  trucks := []Truck{}
  err := db.Select(&trucks, "SELECT id, name, owner_id, lat, lng, url FROM FoodTrucks")
  fmt.Println(trucks)
  if (err != nil){
    fmt.Println(err)
    return
  }
  a, _:= json.Marshal(trucks)
  w.Write(a)
  return
}

func CreateFoodTruck(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Creating food truck!\n"))

  queryParameters := r.URL.Query()

  // surely there's a better way to do this
  name := queryParameters.Get("name")
  ownerid, err1 := strconv.Atoi(queryParameters.Get("ownerid"))
  latitude, err2 := strconv.ParseFloat(queryParameters.Get("lat"), 64)
  longitude, err3 := strconv.ParseFloat(queryParameters.Get("lon"), 64)
  url := queryParameters.Get("url")

  if err1 != nil || err2 != nil || err3 != nil {
    w.Write([]byte(fmt.Sprintf("BAD INPUT :( %s %s %s", err1, err2, err3)))
    return
  }
  addTruck := `INSERT INTO foodtrucks (id, name, owner_id, lat, lng, url) VALUES ($1, $2, $3, $4, $5, $6)`
  var truckId = len(allTrucks)+1
  tx, err := db.Begin()
  _, err = tx.Exec(addTruck, truckId, name, int(ownerid), latitude, longitude, url)
  err = tx.Commit()

  if (err != nil){
    fmt.Println(err)
    return
  }
  // newTruck := Truck{FoodTruckId: len(allTrucks)+1,
  //                  Name: name,
  //                  OwnerId: int(ownerid),
  //                  Latitude: latitude,
  //                  Longitude: longitude,
  //                  Url: url}
  //
  // allTrucks = append(allTrucks, newTruck)
  // a, _:= json.Marshal(allTrucks)
  // w.Write(a)
}

func DeleteFoodTruck(w http.ResponseWriter, r *http.Request) {
  foodTruckId := mux.Vars(r)["id"]

  deleteTruck := `DELETE FROM foodtrucks WHERE id = $1`

  tx, err := db.Begin()
  _, err = tx.Exec(deleteTruck, foodTruckId)
  err = tx.Commit()

  if (err != nil){
    fmt.Println(err)
    return
  }

  w.Write([]byte("Deleting food truck :( " + foodTruckId))
}

func OpenFoodTruck(w http.ResponseWriter, r *http.Request) {
  queryParameters := r.URL.Query()
  foodTruckId := mux.Vars(r)["id"]
  latitude, err2 := strconv.ParseFloat(queryParameters.Get("lat"), 64)
  longitude, err3 := strconv.ParseFloat(queryParameters.Get("lon"), 64)

  // TODO Update the Food Truck status/open timestamp

  w.Write([]byte(fmt.Sprintf("Opening food truck %s at %s,%s ", foodTruckId, latitude, longitude)))
}

func CloseFoodTruck(w http.ResponseWriter, r *http.Request) {
  foodTruckId := mux.Vars(r)["id"]
  // TODO Update the Food Truck status/open timestamp
  w.Write([]byte("Closing food truck :( " + foodTruckId))
}

func GetFoodTruckLocation(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf=8")
  foodTruckId := mux.Vars(r)["id"]

  var truck Truck
  err := db.QueryRowx("SELECT id, name, lat, lng FROM FoodTrucks WHERE id = $1", foodTruckId).StructScan(&truck)
  fmt.Println("Here is food truck location! " + foodTruckId)
  if (err != nil){
    fmt.Println(err)
    return
  }
  a, _:= json.Marshal(truck)
  w.Write(a)
}

func PostFoodTruckLocation(w http.ResponseWriter, r *http.Request) {
  queryParameters := r.URL.Query()
  foodTruckId := mux.Vars(r)["id"]
  latitude, err2 := strconv.ParseFloat(queryParameters.Get("lat"), 64)
  longitude, err3 := strconv.ParseFloat(queryParameters.Get("lon"), 64)

  // TODO Update the Food Truck status/open timestamp

  w.Write([]byte(fmt.Sprintf("Posting food truck location! %s at %s, %s",
                              foodTruckId, latitude, longitude)))
}

func GetTestFoodTrucks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf=8")
  a, _:= json.Marshal(allTrucks)
  w.Write(a)
  return
}

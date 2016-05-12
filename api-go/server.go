package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "strconv"
)

var truck1  = Truck{"1", "Greek Food Truck", -41.292489, 174.778656, true}
var truck2 = Truck{"2", "Beat Kitchen", -41.287022, 174.778667, false}
var truck3 = Truck{"3", "Nanny's Food Truck", -41.290425, 174.779272, true}
var allTrucks = []Truck{truck1, truck2, truck3}
var db = initDb()

func main() {
  mx := mux.NewRouter()

  mx.HandleFunc("/api/trucks/create", CreateFoodTruck).Methods("POST")
  mx.HandleFunc("/api/trucks/delete/{id}", DeleteFoodTruck).Methods("DELETE")
  mx.HandleFunc("/api/trucks/open/{id}", OpenFoodTruck).Methods("POST")
  mx.HandleFunc("/api/trucks/close/{id}", CloseFoodTruck).Methods("POST")
  mx.HandleFunc("/api/trucks/location/all", GetFoodTrucks).Methods("GET")
  mx.HandleFunc("/api/trucks/location/{id}", GetFoodTruckLocation).Methods("GET")
  mx.HandleFunc("/api/trucks/location/{id}", PostFoodTruckLocation).Methods("POST")
  mx.HandleFunc("/api/trucks/test", GetTestFoodTrucks).Methods("GET")

  fmt.Printf("Serving on port 9001")
  http.ListenAndServe(":9001", mx)
}

func WriteAllowOriginHeader(w http.ResponseWriter) {
  w.Header().Set("Content-Type", "application/json; charset=utf=8")
  w.Header().Add("Access-Control-Allow-Origin", "http://whatsontheme.nu")
}

func GetNearbyFoodTrucks(w http.ResponseWriter, r *http.Request) {
  WriteAllowOriginHeader(w)

  queryParameters := r.URL.Query()
  latitude, err2 := strconv.ParseFloat(queryParameters.Get("lat"), 64)
  longitude, err3 := strconv.ParseFloat(queryParameters.Get("lon"), 64)


  if err2 != nil || err3 != nil {
    w.Write([]byte(fmt.Sprintf("BAD INPUT :( %s %s", err2, err3)))
    return
  }

  trucks := []Truck{}
  // TODO needs id, name etc?
  err := db.Select(&trucks, "SELECT id, name, lat, lng, url FROM FoodTrucks WHERE ST_DWithin(geom,  ST_GeomFromText('POINT($1 $2)', 4326),1000,false)", longitude,latitude)
  fmt.Println(trucks)
  if (err != nil){
    fmt.Println(err)
    return
  }
  a, _:= json.Marshal(trucks)
  w.Write(a)
  return
}

func GetFoodTrucks(w http.ResponseWriter, r *http.Request) {
  WriteAllowOriginHeader(w)
  trucks := []Truck{}
  err := db.Select(&trucks, "SELECT id, name, lat, lng, is_open FROM FoodTrucks")
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
  latitude, err2 := strconv.ParseFloat(queryParameters.Get("lat"), 64)
  longitude, err3 := strconv.ParseFloat(queryParameters.Get("lon"), 64)
  url := queryParameters.Get("url")

  if err2 != nil || err3 != nil {
    w.Write([]byte(fmt.Sprintf("BAD INPUT :( %s %s", err2, err3)))
    return
  }
  addTruck := `INSERT INTO foodtrucks (id, name, lat, lng, url, geom) VALUES ($1, $2, $3, $4, $5, ST_GeomFromText('POINT($4 $3)', 4326))`
  var truckId = len(allTrucks)+1
  tx, err := db.Begin()
  _, err = tx.Exec(addTruck, truckId, name, latitude, longitude, url)
  err = tx.Commit()

  if (err != nil){
    fmt.Println(err)
    return
  }

  // hack
  newTruck := Truck{FoodTruckId: string(len(allTrucks)+1),
                   Name: name,
                   Latitude: latitude,
                   Longitude: longitude,
                   IsOpen: false}
  allTrucks = append(allTrucks, newTruck)
  a, _:= json.Marshal(allTrucks)
  w.Write(a)
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

  // hack
  var newAllTrucks = []Truck{}
  for i := 0; i < len(allTrucks); i++ {
    if allTrucks[i].FoodTruckId != foodTruckId {
      newAllTrucks = append(newAllTrucks, allTrucks[i])
    }
  }
  allTrucks = newAllTrucks
}

func OpenFoodTruck(w http.ResponseWriter, r *http.Request) {
  queryParameters := r.URL.Query()
  foodTruckId := mux.Vars(r)["id"]
  latitude, err2 := strconv.ParseFloat(queryParameters.Get("lat"), 64)
  longitude, err3 := strconv.ParseFloat(queryParameters.Get("lon"), 64)

  if err2 != nil || err3 != nil {
    w.Write([]byte(fmt.Sprintf("BAD INPUT :( %s %s", err2, err3)))
    return
  }

  update := `UPDATE foodtrucks SET (lat, lng, is_open) = ($2, $3, true) WHERE id = $1`
  tx, err := db.Begin()
  _, err = tx.Exec(update, foodTruckId, latitude, longitude)
  err = tx.Commit()

  if (err != nil){
    fmt.Println(err)
    return
  }

  // hack
  for i := 0; i < len(allTrucks); i++ {
    if allTrucks[i].FoodTruckId == foodTruckId {
      allTrucks[i].IsOpen = true
      allTrucks[i].Latitude = latitude
      allTrucks[i].Longitude = longitude
    }
  }

  w.Write([]byte(fmt.Sprintf("Opening food truck %s at %s,%s ", foodTruckId, latitude, longitude)))
}

func CloseFoodTruck(w http.ResponseWriter, r *http.Request) {
  foodTruckId := mux.Vars(r)["id"]

  update := `UPDATE foodtrucks SET is_open = false WHERE id = $1`
  tx, err := db.Begin()
  _, err = tx.Exec(update, foodTruckId)
  err = tx.Commit()

  if (err != nil){
    fmt.Println(err)
    return
  }

  // hack
  for i := 0; i < len(allTrucks); i++ {
    if allTrucks[i].FoodTruckId == foodTruckId {
      allTrucks[i].IsOpen = false
    }
  }

  w.Write([]byte("Closing food truck :( " + foodTruckId))
}

func GetFoodTruckLocation(w http.ResponseWriter, r *http.Request) {
  WriteAllowOriginHeader(w)
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

  if err2 != nil || err3 != nil {
    w.Write([]byte(fmt.Sprintf("BAD INPUT :( %s %s", err2, err3)))
    return
  }
  update := `UPDATE foodtrucks SET (lat, lng) = ($2, $3) WHERE id = $1`
  tx, err := db.Begin()
  _, err = tx.Exec(update, foodTruckId, latitude, longitude)
  err = tx.Commit()

  if (err != nil){
    fmt.Println(err)
    return
  }

  // hack
  for i := 0; i < len(allTrucks); i++ {
    if allTrucks[i].FoodTruckId != foodTruckId {
      allTrucks[i].Latitude = latitude
      allTrucks[i].Longitude = longitude
    }
  }

  w.Write([]byte(fmt.Sprintf("Posting food truck location! %s at %s, %s",
                              foodTruckId, latitude, longitude)))
}

func GetTestFoodTrucks(w http.ResponseWriter, r *http.Request) {
  WriteAllowOriginHeader(w)
  a, _:= json.Marshal(allTrucks)
  w.Write(a)
  return
}

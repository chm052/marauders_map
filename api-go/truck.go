package main

type Truck struct {
  FoodTruckId int `db:"id"`
  Name string
  OwnerId int `db:"owner_id"`
  Latitude float64 `db:"lat"`
  Longitude float64 `db:"lng"`
  Url string
}

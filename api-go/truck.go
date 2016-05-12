package main

type Truck struct {
  FoodTruckId string `db:"id"`
  Name string
  Latitude float64 `db:"lat"`
  Longitude float64 `db:"lng"`
  IsOpen bool `db:"is_open"`
}

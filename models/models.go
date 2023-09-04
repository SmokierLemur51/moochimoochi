package models 

import (
    "time"
)



// data for rendering regular pages
type PageData struct {
	Page 	string
	Title	string
	CoilSelection []GutterCoil
}



type Order struct {
    ID                  int
    OrderDate           time.Time
    RequestedDelivery   time.Time
    CustomerContact     []CustomerContactInformation
    GutterCoil          []GutterCoil
    GutterRuns          []GutterRun
    DeliveryCharge      []DeliveryCharge
    Total               float32
}


type CustomerContactInformation struct {
    Name    string
    Phone   string
    Address []Address
}


type Address struct {
    ID int

}


type GutterCoil struct {
    ID                      int
    Size                    string
    Color                   string
    OnHand                  float32
    PoundsFeetMulti         float32
    DisplayInformation      string
    Cost                    float32
    Margin                  float32
    Selling                 float32
}



type GutterRun struct {}



type DeliveryCharge struct {
    ID      int
    Miles   float32
    TollFee float32
    Total   float32
}


type User struct {}


type AdministrativeFeed struct {
	Page string
	Title string
}







func CostToSelling(cost, margin float32) float32 {
    pointsMargin := 1.00 - (margin/100)
    // need to implement rounding ... 
    return (cost / pointsMargin)
}
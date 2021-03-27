/*
Copyright 2021 © The Onestop Authors

All Rights Reserved.

NOTICE: All information contained herein is, and remains the property of
The Onestop Authors. The intellectual and technical concepts contained
herein are proprietary to The Onestop Authors. Dissemination of this
information or reproduction of this material is strictly forbidden unless
prior written permission is obtained from The Onestop Authors.

Authors: Manish Sahani          <rec.manish.sahani@gmail.com>
         Priyadarshan Singh	    <singhpd75@gmail.com>

*/

package models

//import (
//"time"
//)

type TripSearch struct {
	Source      Location `json:"source" binding:"required"`
	Destination Location `json:"destination" binding:"required"`
	FromDate    string
	ToDate      string
	//FromDate    time.Time `json:"from_date" binding:"required"`
	//ToDate      time.Time `json:"to_date" binding:"required"`
}

type TequilaData struct {
	Duration struct {
		Departure uint
		Return    uint
		Total     uint
	}
	FlyFrom      string
	CityFrom     string
	CityCodeFrom string
	CountryFrom  struct {
		Code string
		Name string
	}
	FlyTo      string
	CityTo     string
	CityCodeTo string
	CountryTo  struct {
		Code string
		Name string
	}
	Distance       uint
	Price          uint     `json:"price"`
	Route          []Flight `json:"route"`
	LocalArrival   string   `json:"local_arrival"`
	UTCArrival     string   `json:"utc_arrival"`
	LocalDeparture string   `json:"local_departure"`
	UTCDeparture   string   `json:"utc_departure"`
}
type Tequila struct {
	Data []TequilaData `json:"data"`
}

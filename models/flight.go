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

import "gorm.io/gorm"

type Flight struct {
	gorm.Model
	CityTo         string   `json:"cityTo"`
	CityFrom       string   `json:"cityFrom"`
	CityCodeFrom   string   `json:"cityCodeFrom"`
	CityCodeTo     string   `json:"cityCodeTo"`
	FlyTo          string   `json:"flyTo"`
	FlyFrom        string   `json:"flyFrom"`
	Airline        string   `json:"airline"`
	FlightNo       string   `json:"flight_no"`
	LocalArrival   string   `json:"local_arrival"`
	UTCArrival     string   `json:"utc_arrival"`
	LocalDeparture string   `json:"local_departure"`
	UTCDeparture   string   `json:"utc_departure"`
	Segment       Segment `gorm:"polymorphic:Vendor"`
}

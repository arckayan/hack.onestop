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

import (
	"time"

	"gorm.io/gorm"
)

type Cab struct {
	gorm.Model
	FromLat       float64
	FromLng       float64
	ToLat         float64
	ToLng         float64
	ArrivalTime   time.Time
	DepartureTime time.Time
	ExpectedPrice float64
	Distance      float64
	ExpectedTime  float64
	Segment       Segment `gorm:"polymorphic:Vendor"`
}

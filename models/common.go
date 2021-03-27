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
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

var Type = make(map[string]reflect.Type)

type Location struct {
	gorm.Model
	Lat      string `binding:"required"`
	Lng      string `binding:"required"`
	City     string `binding:"required"`
	State    string
	TripID   uint  `gorm:"default:null"`
	Airports []int `gorm:"type:text"`
}

func RegisterTypes() {
	for _, t := range []interface{}{Cab{}, Flight{}} {
		Type[fmt.Sprintf("%T", t)] = reflect.TypeOf(t)
	}
}

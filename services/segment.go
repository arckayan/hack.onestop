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

package services

import (
	"errors"

	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/models"
)

type Segment struct {
}

// Find the segment from id
func (s *Segment) Find(ID uint) (*models.Segment, interface{}, error) {
	var segment models.Segment

	if err := core.K.DB.Engine.Where("id = ?", ID).First(&segment).Error; err != nil {
		return nil, nil, errors.New("Segment does not exist.")
	}

	//v := reflect.New(models.Type["models."+strings.Title(segment.VendorType[:len(segment.VendorType)-1])]).Elem().Interface()
	if segment.VendorType == "flights" {
		var v models.Flight
		core.K.DB.Engine.Where("ID = ?", segment.VendorID).First(&v)
		return &segment, v, nil
	} else {
		var v models.Cab
		core.K.DB.Engine.Where("ID = ?", segment.VendorID).First(&v)
		return &segment, v, nil
	}
}

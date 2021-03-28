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
	"time"

	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/models"
)

type Trip struct{}

// Create a trip from the search :
func (s *Trip) Create(from *models.Location, to *models.Location, fromDate string, user *models.User) (*models.Trip, error) {
	// Process the arguments and create a trip object
	trip := models.Trip{
		Source:      models.Location{City: from.City, State: from.State, Lat: from.Lat, Lng: from.Lng},
		Destination: models.Location{City: to.City, State: to.State, Lat: to.Lat, Lng: to.Lng},
		UserID:      user.ID,
		Date:        fromDate,
		Expires:     time.Now().Add(time.Minute * 30),
	}

	// Create a flight for the trip and save in the database
	core.K.DB.Engine.Create(&trip)
	if trip.ID == 0 {
		return nil, errors.New("Could not create a flight.")
	}

	return &trip, nil
}

type Res struct {
	Segment models.Segment
	Vendor  interface{}
}

// Find the trip from uuid
func (t *Trip) Find(UUID string) (*models.Trip, []Res, error) {
	var trip models.Trip
	var segments []models.Segment

	// find the trip from the uuid
	if err := core.K.DB.Engine.Model(&trip).Where("uuid = ?", UUID).First(&trip).Association("Segments").Find(&segments).Error; err != nil {
		return nil, nil, errors.New("Trip with uuid does not exist")
	}

	core.K.DB.Engine.Model(&trip).Association("Segments").Find(&segments)
	var result []Res
	for _, s := range segments {
		_, v, _ := new(Segment).Find(s.ID)
		r := Res{Segment: s, Vendor: v}
		result = append(result, r)
	}

	return &trip, result, nil
}

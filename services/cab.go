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

type Cab struct{}

// Create a resource from the input data
func (s *Cab) Create(from *models.Coordinate, to *models.Coordinate, tripID uint) (*models.Cab, error) {
	// Process the arguments and create a cab object
	cab := models.Cab{
		FromLat:       from.Lat,
		FromLng:       from.Lng,
		ToLat:         to.Lat,
		ToLng:         to.Lng,
		ExpectedPrice: s.Price(distance(from.Lat, from.Lng, to.Lat, to.Lng)),
		Segment:       models.Segment{TripID: tripID},
	}

	// Create a cab for the trip and save in the database
	core.K.DB.Engine.Create(&cab)
	if cab.ID == 0 {
		return nil, errors.New("Could not create a cab ride.")
	}

	return &cab, nil
}

// Price return the current cab rate for the location
func (s *Cab) Price(distance float64) float64 {
	return distance*8 + 45
}

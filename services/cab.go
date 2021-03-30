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

type Cab struct{}

// Create a resource from the input data
func (s *Cab) Create(from *models.Coordinate, to *models.Coordinate, ref time.Time, forward bool, tripID uint) (*models.Cab, error) {
	// Process the arguments and create a cab object
	dist := distance(from.Lat, from.Lng, to.Lat, to.Lng)
	cab := models.Cab{
		FromLat:       from.Lat,
		FromLng:       from.Lng,
		ToLat:         to.Lat,
		ToLng:         to.Lng,
		Distance:      dist,
		ArrivalTime:   arrivaltime(ref, int(s.Time(dist)), forward),
		DepartureTime: departuretime(ref, int(s.Time(dist)), forward),
		ExpectedTime:  s.Time(dist),
		ExpectedPrice: s.Price(dist),
		Segment:       models.Segment{TripID: tripID},
	}

	// Create a cab for the trip and save in the database
	core.K.DB.Engine.Create(&cab)
	if cab.ID == 0 {
		return nil, errors.New("Could not create a cab ride.")
	}

	return &cab, nil
}

func departuretime(ref time.Time, dur int, forward bool) time.Time {
	t := ref
	if forward {
		t = t.Add(time.Minute * 90)
		t = t.Add(time.Minute * time.Duration(dur))
		return t
	}

	t = t.Add(-time.Minute * 90)
	return t
}

func arrivaltime(ref time.Time, dur int, forward bool) time.Time {
	t := ref
	if forward {
		t = t.Add(time.Minute * 90)
		return t
	}

	t = t.Add(-time.Minute * 90)
	t = t.Add(-time.Minute * time.Duration(dur))
	return t
}

// Price return the current cab rate for the location
func (s *Cab) Price(distance float64) float64 {
	return (distance/1000)*8 + 45
}

func (s *Cab) Time(distance float64) float64 {
	return ((distance/1000)/50)*60 + 15
}

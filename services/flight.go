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
	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/models"
)

type Flight struct{}

// Search Flight
func (s *Flight) Create(f *models.Flight, tripID uint) (*models.Flight, error) {
	flight := models.Flight{
		CityTo:         f.CityTo,
		CityFrom:       f.CityFrom,
		CityCodeTo:     f.CityCodeTo,
		CityCodeFrom:   f.CityCodeFrom,
		FlyTo:          f.FlyTo,
		FlyFrom:        f.FlyFrom,
		Airline:        f.Airline,
		FlightNo:       f.FlightNo,
		LocalArrival:   f.LocalArrival,
		UTCArrival:     f.UTCArrival,
		LocalDeparture: f.LocalDeparture,
		UTCDeparture:   f.UTCDeparture,
		Duration:       int(f.UTCArrival.Sub(f.UTCDeparture).Minutes()),
		Segment:        models.Segment{TripID: tripID},
	}

	if err := core.K.DB.Engine.Create(&flight).Error; err != nil {
		return nil, err
	}

	return &flight, nil
}

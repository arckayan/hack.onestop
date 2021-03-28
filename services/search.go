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
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/models"
)

type Search struct{}

func (s *Search) SearchAiportsInCity(l *models.Location) []models.Airport {
	// This list will be populated with the search
	var airports []models.Airport

	// TODO: better search query
	core.K.DB.Engine.Where("city LIKE ? OR state LIKE ?", "%"+l.City+"%", "%"+l.City+"%").Find(&airports)

	// TODO: sort the search results based on the user's locations

	return airports
}

func (s *Search) SearchEndToEndTrips(ts *models.TripSearch, user *models.User) (*[]models.Trip, error) {
	//
	from := models.Coordinate{Lat: ts.Source.Lat, Lng: ts.Source.Lng}
	to := models.Coordinate{Lat: ts.Destination.Lat, Lng: ts.Destination.Lng}
	//
	fromAirport, _ := new(Airport).Find(ts.Source.Airports[0])
	toAiport, _ := new(Airport).Find(ts.Destination.Airports[0])

	// Call flight api for the best available offers
	tequila := SearchFlightOptions(fromAirport, toAiport, ts.FromDate, ts.ToDate)

	// User's coordinates
	var trips []models.Trip

	for _, shot := range tequila {
		// Create a trip
		trip, _ := new(Trip).Create(&ts.Source, &ts.Destination, ts.FromDate, user)

		// segments table
		var s []models.Segment

		// Create a cab from location to airports
		c1, _ := new(Cab).Create(&from, &models.Coordinate{Lat: fromAirport.Lat, Lng: fromAirport.Lng}, trip.ID)
		s = append(s, c1.Segment)

		for _, flight := range shot.Route {
			f, _ := new(Flight).Create(&flight, trip.ID)
			s = append(s, f.Segment)
		}

		c2, _ := new(Cab).Create(&models.Coordinate{Lat: toAiport.Lat, Lng: toAiport.Lng}, &to, trip.ID)
		s = append(s, c2.Segment)

		trip.EstimatedPrice = shot.Price + c1.ExpectedPrice + c2.ExpectedPrice
		core.K.DB.Engine.Save(&trip)

		core.K.DB.Engine.Model(&trip).Association("Segments").Append(s)
		trips = append(trips, *trip)
	}

	return &trips, nil
}

func SearchFlightOptions(from *models.Airport,
	to *models.Airport, ft string, tt string) []models.TequilaData {

	flightAPI := fmt.Sprintf(
		"%s/search?fly_from=%s&fly_to=%s&dateFrom=%s&dateTo=%s&curr=%s",
		core.Config("kiwiURL"),
		from.Code,
		to.Code,
		ft,
		tt,
		"INR",
	)

	println(flightAPI)

	req, _ := http.NewRequest("GET", flightAPI, nil)
	req.Header.Add("apikey", core.Config("kiwiAPIKey"))

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var t models.Tequila
	json.Unmarshal([]byte(body), &t)

	return t.Data
}

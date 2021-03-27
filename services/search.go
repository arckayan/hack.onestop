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
	//sort.StringSlice(airports, func (a, b models.Airport) bool {
	//return distance()
	//})

	return airports
}

func (s *Search) SearchEndToEndTrips(ts *models.TripSearch, user *models.User) []models.Trip {

	// Call flight api for the best available offers
	//tequila := SearchFlightOptions(&ts.Source, &ts.Destination, ts.FromDate, ts.ToDate)

	// Create a trip
	trips := []models.Trip{
		{
			Source: models.Location{
				Lat:  ts.Source.Lat,
				Lng:  ts.Source.Lng,
				City: ts.Source.City,
			},
			Destination: models.Location{
				Lat:  ts.Destination.Lat,
				Lng:  ts.Destination.Lng,
				City: ts.Destination.City,
			},
			UserID: user.ID,
		},
	}
	core.K.DB.Engine.Create(&trips)

	f := models.Flight{
		Segment: models.Segment{TripID: trips[0].ID},
	}
	core.K.DB.Engine.Create(&f)
	core.K.DB.Engine.Model(&trips[0]).Association("Segments").Append([]models.Segment{
		f.Segment,
	})

	return trips

	/*
		for _, f := range tequila {
			source := ts.Source
			destination := ts.Destination
			trip := models.Trip{
				Persist:     false,
				Source:      source,
				Destination: destination,
				Date:        ts.FromDate,
				Expires:     time.Now().Add(time.Minute * 5),
			}
			core.K.DB.Engine.Create(&trip)
			fmt.Println(f)

			// flight
			//f := flight.Route[0]
			//s := models.Segment{TripID: trip.ID}
			//f.Segment = s
			//core.K.DB.Engine.Model(&trip).Association("Segments").Append([]models.Segment{s})

			trips = append(trips, trip)
		}

	*/
	// Compare distance from user's pickup to source airport

	// Compare distance from destination airport to user's drop off

	// Generate an Itenary

	return trips
}

func SearchFlightOptions(source *models.Location,
	destination *models.Location, ft string, tt string) []models.TequilaData {
	// Get Airports from Airport's Ids
	var from models.Airport
	var to models.Airport

	core.K.DB.Engine.First(&from, source.Airports[0])
	core.K.DB.Engine.First(&to, destination.Airports[0])

	flightAPI := fmt.Sprintf(
		"%s/search?fly_from=%s&fly_to=%s&dateFrom=%s&dateTo=%s&curr=%s",
		core.Config("kiwiURL"),
		from.Code,
		to.Code,
		ft,
		tt,
		"INR",
	)

	req, _ := http.NewRequest("GET", flightAPI, nil)
	req.Header.Add("apikey", core.Config("kiwiAPIKey"))

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var t models.Tequila
	json.Unmarshal([]byte(body), &t)

	return t.Data
}

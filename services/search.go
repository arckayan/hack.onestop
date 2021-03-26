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

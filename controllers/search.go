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

package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kalkayan/onestop/models"
	"github.com/kalkayan/onestop/services"
)

type Search struct{ Controller }

// EndToEndTrip searches all the trip routes for the user -- including flight,
// train, and buses all accomodated by cabs.
func (c *Search) EndToEndTrips(ctx *gin.Context) {
	var ts models.TripSearch

	// Validate the request for the handle
	if err := ctx.ShouldBind(&ts); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		fmt.Println(err.Error())
		return
	}

	// searching and optimizing
	user, _ := ctx.Get("user")
	trips, _ := new(services.Search).SearchEndToEndTrips(&ts, user.(*models.User))

	c.OK(ctx, gin.H{"search": trips})
}

// AirportInCity is used for setting the airports (source and destination) for
// End-to-End search engine.
func (c *Search) AirportInCity(ctx *gin.Context) {
	var location models.Location

	// Validate the request for the handle
	if err := ctx.ShouldBind(&location); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}

	// actual searching and sorting
	airports := new(services.Search).SearchAiportsInCity(&location)

	c.OK(ctx, gin.H{"search": airports})
}

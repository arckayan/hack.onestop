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
	"github.com/gin-gonic/gin"
	"github.com/kalkayan/onestop/models"
	"github.com/kalkayan/onestop/services"
)

type Trip struct{ Controller }

func (c *Trip) All(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	trips := new(services.Trip).All(user.(*models.User))
	c.OK(ctx, gin.H{"trips": trips})
}

// Find retrieve the resource from the params
func (c *Trip) Find(ctx *gin.Context) {
	var t ParamUUID

	// Validate the request for the handle
	if err := ctx.ShouldBindUri(&t); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}

	// find the trip and segements
	trip, segments, err := new(services.Trip).FindWithSegments(t.UUID)
	if err != nil {
		c.NotFound(ctx, err.Error())
		return
	}

	c.OK(ctx, gin.H{"trip": trip, "segments": segments})
}

func (c *Trip) Update(ctx *gin.Context) {
	var d models.Trip
	var t ParamUUID

	// Validate the request for the handle
	if err := ctx.ShouldBindUri(&t); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}
	// Validate the repuest binding
	if err := ctx.ShouldBindJSON(&d); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}

	c.OK(ctx, "The func will be soon implemented.")
}

func (c *Trip) Book(ctx *gin.Context) {
	var t ParamUUID

	// Validate the request for the handle
	if err := ctx.ShouldBindUri(&t); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}

	trip, err := new(services.Trip).Find(t.UUID)
	if err != nil {
		c.NotFound(ctx, err.Error())
	}
	if err := new(services.Trip).Book(trip); err != nil {
		c.NotAcceptable(ctx, err.Error())
	}
	c.OK(ctx, "The trip has been booked")
}

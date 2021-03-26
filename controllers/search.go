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
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalkayan/onestop/models"
	"github.com/kalkayan/onestop/services"
)

type Search struct{ Controller }

func (c *Search) EndToEndTrip(ctx *gin.Context) {

}

func (c *Search) AirportInCity(ctx *gin.Context) {
	var location models.Location

	// Validate the request for the handle
	if err := ctx.ShouldBind(&location); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// actual searching and sorting
	airports := new(services.Search).SearchAiportsInCity(&location)

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"search": airports,
		},
	})
}

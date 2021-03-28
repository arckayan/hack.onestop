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
	"github.com/kalkayan/onestop/services"
)

type Trip struct{ Controller }

// Find retrieve the resource from the params
func (c *Trip) Find(ctx *gin.Context) {
	var t ParamUUID

	// Validate the request for the handle
	if err := ctx.ShouldBindUri(&t); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}

	// find the trip and segements
	trip, segments, _ := new(services.Trip).Find(t.UUID)

	c.OK(ctx, gin.H{"trip": trip, "segments": segments})
}

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

type Segment struct {
	Controller
	Service services.Segment
}

// Find retrieve the resource from the params
func (c *Segment) Find(ctx *gin.Context) {
	var s ParamID

	// Validate the request for the handle
	if err := ctx.ShouldBindUri(&s); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}

	segment, v, err := new(services.Segment).Find(s.ID)
	if err != nil {
		c.NotFound(ctx, err.Error())
		return
	}

	c.OK(ctx, gin.H{"search": segment, "vendor": v})
}

// Update a particular resource
func (c *Segment) Update(ctx *gin.Context) {
	var s ParamID

	// Validate the request for the handle
	if err := ctx.ShouldBindUri(&s); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}

	segment, v, err := new(services.Segment).Find(s.ID)
	if err != nil {
		c.NotFound(ctx, err.Error())
		return
	}

	c.OK(ctx, gin.H{"search": segment, "vendor": v})
}

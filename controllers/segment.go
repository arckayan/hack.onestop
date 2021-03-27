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
	"github.com/kalkayan/onestop/services"
)

type Segment struct{ Controller }

func (c *Segment) Find(ctx *gin.Context) {
	type SegmentURI struct {
		UUID string `uri:"uuid" binding:"required"`
	}
	var s SegmentURI

	if err := ctx.ShouldBindUri(&s); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
		return
	}

	segment, err := new(services.Segment).Find(s.UUID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"search": segment,
		},
	})

}

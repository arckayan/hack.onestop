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
)

type Controller struct{}

func (c Controller) ValidateBindings(ctx *gin.Context, schema interface{}) bool {
	// Validate the request for the handle
	err := ctx.ShouldBindJSON(&schema)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return false
	}

	return true
}

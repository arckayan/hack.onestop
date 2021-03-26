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
)

type Trip struct{ Controller }

func (c *Trip) Search(ctx *gin.Context) {
	var trip models.Trip

	//if !c.ValidateBindings(ctx, &trip) {
	//return
	//}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"trip":    trip,
			"message": "Hello this is the search",
		},
	})
}

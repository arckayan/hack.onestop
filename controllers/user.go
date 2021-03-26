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

type User struct{ Controller }

func (c *User) Me(ctx *gin.Context) {
	user, set := ctx.Get("user")

	if set {
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"user": user,
			},
		})

		return
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{
		"message": "User has been logged out, please login again",
	})
}

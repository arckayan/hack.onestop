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

type AuthController struct{ Controller }

// Register creates and return a new user
func (c AuthController) Register(ctx *gin.Context) {
	var user models.User

	if !c.ValidateBindings(ctx, &user) {
		return
	}

	if err := new(services.UserService).Create(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user.Transform(),
		},
	})
}

// Login authenticates a user from credentials
func (c AuthController) Login(ctx *gin.Context) {
	var creds models.Credentials

	if !c.ValidateBindings(ctx, &creds) {
		return
	}

	token, err := new(services.AuthService).Authenticate(&creds)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"token": token,
		},
	})
}

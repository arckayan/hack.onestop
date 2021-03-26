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

type Auth struct{ Controller }

// Register creates and return a new user
func (c Auth) Register(ctx *gin.Context) {
	var user models.User

	// Validate the request for the handle
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// Create the User in the database
	if err := new(services.User).Create(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Authenticate the user and create a token
	token, err := new(services.Auth).CreateAccessToken(user.UUID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the token and the newly created user info
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":  user.Transform(),
			"token": token,
		},
	})
}

// Login authenticates a user from credentials
func (c Auth) Login(ctx *gin.Context) {
	var creds models.Credentials

	// Validate the request for the handle
	if err := ctx.ShouldBind(&creds); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// Authenticate the user and create a token
	token, err := new(services.Auth).Authenticate(&creds)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the access and refresh token to user
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"token": token,
		},
	})
}

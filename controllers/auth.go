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

type Auth struct{ Controller }

// Register creates and return a new user
func (c Auth) Register(ctx *gin.Context) {
	var user models.User

	// Validate the request for the handle
	if err := ctx.ShouldBind(&user); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}

	// Create the User in the database
	if err := new(services.User).Create(&user); err != nil {
		c.NotAcceptable(ctx, err.Error())
		return
	}

	// Authenticate the user and create a token
	token, err := new(services.Auth).CreateAccessToken(user.UUID)
	if err != nil {
		c.NotAcceptable(ctx, err.Error())
		return
	}

	c.Created(ctx, gin.H{"user": user.Transform(), "token": token})
}

// Login authenticates a user from credentials
func (c Auth) Login(ctx *gin.Context) {
	var creds models.Credentials

	// Validate the request for the handle
	if err := ctx.ShouldBind(&creds); err != nil {
		c.UnprocessableEntity(ctx, err.Error())
		return
	}

	// Authenticate the user and create a token
	token, err := new(services.Auth).Authenticate(&creds)
	if err != nil {
		c.NotAcceptable(ctx, err.Error())
		return
	}

	c.OK(ctx, gin.H{"token": token})
}

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
)

type User struct{ Controller }

// Me return the current authenticated user
func (c *User) Me(ctx *gin.Context) {
	user, set := ctx.Get("user")

	// Abort the request if the user is not set
	if !set {
		c.Unauthorized(ctx, "User has been logged out, please login again")
		return
	}

	c.OK(ctx, gin.H{"user": user})
}

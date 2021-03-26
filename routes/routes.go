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

package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kalkayan/onestop/controllers"
	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/services"
)

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(bearer) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token entry",
			})
			return
		}

		user, err := new(services.Auth).ValidateTokenAndAuthenticate(bearer[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("user", user.Transform())
		c.Next()
	}
}

func Register(r *core.Router) {

	auth := new(controllers.Auth)
	user := new(controllers.User)
	search := new(controllers.Search)

	v1 := r.Engine.Group("v1")
	{
		v1.POST("/register", auth.Register)
		v1.POST("/login", auth.Login)

		private := v1.Group("app")
		private.Use(AuthGuard())
		{
			private.GET("/me", user.Me)

			// Search endpoints
			private.POST("/search/trip", search.EndToEndTrip)
			private.POST("/search/airports", search.AirportInCity)
		}
	}
}

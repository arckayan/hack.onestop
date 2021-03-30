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

// AuthGuard checks and validates the incoming requests for authorised private
// routes. Also sets a global user to the requests.
func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Authorization token from the requests headers and split
		// according to the jwt specs.
		bearer := strings.Split(c.Request.Header.Get("Authorization"), " ")

		// Abort the requests if the bearer token is not present in the request
		// headers.
		if len(bearer) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token entry",
			})
			return
		}

		// Verify the token, find the user and Authenticate the request
		user, err := new(services.Auth).ValidateTokenAndAuthenticate(bearer[1])

		// Abort the request if the user is not found
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Set the user to the global request context. Process the request,
		// move to the particular handle for the request
		c.Set("user", user)
		c.Next()
	}
}

// Register the application's routes to the kernel's router
func Register(r *core.Router) {
	// Controllers for the routes
	var auth controllers.Auth
	var user controllers.User
	var search controllers.Search
	var trip controllers.Trip
	var segment controllers.Segment

	// Register the routes and the routes groups in the kernel's router
	v1 := r.Engine.Group("v1")
	{
		v1.GET("/", func(c *gin.Context) { c.JSON(200, "hi, this is onestop.") })

		// Authentication endpoints
		v1.POST("/register", auth.Register)
		v1.POST("/login", auth.Login)

		// Private routes (requires token)
		private := v1.Group("app")
		private.Use(AuthGuard())
		{
			// User's profile endpoints
			private.GET("/me", user.Me)

			// Search endpoints
			private.POST("/search/trip", search.EndToEndTrips)
			private.POST("/search/airports", search.AirportInCity)

			// Trip endpoints
			private.GET("/trips", trip.All)
			private.GET("/trip/:uuid", trip.Find)
			private.PUT("/trip/:uuid", trip.Update)

			// Segment endpoints
			private.GET("/segment/:uuid", segment.Find)
			private.PUT("/segment/:uuid", segment.Update)
		}
	}
}

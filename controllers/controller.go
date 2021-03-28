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

// struct for binding request params
type ParamUUID struct {
	UUID string `uri:"uuid" binding:"required"`
}

// struct for binding request params
type ParamID struct {
	ID uint `uri:"uuid" binding:"required"`
}

// OK writes the content to the request's context
func (c *Controller) OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// NotFound aborts the request and with status Unauthorized
func (c *Controller) NotFound(ctx *gin.Context, err interface{}) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"errors": err,
	})
}

// Unauthorized aborts the request and with status Unauthorized
func (c *Controller) Unauthorized(ctx *gin.Context, err interface{}) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"errors": err,
	})
}

// UnprocessableEntity aborts the request with status UnprocessableEntity
func (c *Controller) UnprocessableEntity(ctx *gin.Context, err interface{}) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"errors": err,
	})
}

// Created writes the content to the request's content with status Created(201)
func (c *Controller) Created(ctx *gin.Context, res interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"data": res,
	})
}

// NotAcceptable aborts  the request with status UnprocessableEntity
func (c *Controller) NotAcceptable(ctx *gin.Context, err interface{}) {
	ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
		"errors": err,
	})
}

// ValidateBindings checks
func (c Controller) ValidateBindings(ctx *gin.Context, schema interface{}) {
	// Validate the request for the handle
	if err := ctx.ShouldBind(&schema); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
	}
}

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

package models

type Credentials struct {
	Email    string `binding:"required" form:"email"`
	Password string `binding:"required" form:"password"`
}

type AuthenticationToken struct {
	AccessToken    string
	RefreshToken   string
	AccessUUID     string
	RefreshUUID    string
	AccessExpires  int64
	RefreshExpires int64
}

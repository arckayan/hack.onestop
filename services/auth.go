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

package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/models"
)

// Auth handles application's all authentication related operations
type Auth struct{}

// Authenticate is used to login a user from credentials
func (s *Auth) Authenticate(creds *models.Credentials) (*models.Token, error) {
	var user models.User

	// Check if the user with given email exist or not
	if err := core.K.DB.Engine.Where("email = ? ", creds.Email).First(&user).Error; err != nil {
		return nil, errors.New("User does not exist.")
	}

	// Check if the given password matches the original user's password
	if !user.Authenticate(creds.Password) {
		return nil, errors.New("Invalid credentials provided")
	}

	// Create a token which will be further used for authentication (stateless)
	token, err := s.CreateAccessToken(user.UUID)
	if err != nil {
		return nil, errors.New("Unable to create a token")
	}

	return token, nil
}

// CreateToken creates a authentication token for the user
func (s *Auth) CreateAccessToken(userUUID uuid.UUID) (
	*models.Token,
	error,
) {
	var err error

	token := models.Token{}
	token.Expires = time.Now().Add(time.Minute * 60).Unix()

	// Define Access claims and expiry
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = userUUID.String()
	claims["expires"] = token.Expires

	// Create Access claims and singed token
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token.Token, err = access.SignedString([]byte(core.Config("secret"))); err != nil {
		return nil, err
	}

	return &token, nil
}

// ValidateTokenAndAuthenticate is the helper for auth guard middleware, it first
// verifies the token and then check if the claims are valid for the authentication.
func (s *Auth) ValidateTokenAndAuthenticate(tokenstr string) (*models.User, error) {
	// Verify token's signing method
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method of the token")
		}

		return []byte(core.Config("secret")), nil
	})

	// if the token is not valid abort the request with err
	if err != nil {
		return nil, err
	}

	// Extract Access UUID from the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("Token not valid")
	}

	userUUID, ok := claims["user"].(string)
	if !ok {
		return nil, errors.New("login expired, please re-login")
	}

	// Verfiy if the authentication record exist in cache or not
	var user models.User
	if err := core.K.DB.Engine.Where("uuid = ? ", userUUID).First(&user).Error; err != nil {
		return nil, errors.New("User does not exist.")
	}

	return &user, nil
}

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

// AuthService handles application's all authentication related operations
type AuthService struct{}

// Authenticate is used to login a user from credentials
func (s *AuthService) Authenticate(creds *models.Credentials) (*models.AuthenticationToken, error) {
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
	token, err := s.CreateToken(user.UUID)
	if err != nil {
		return nil, errors.New("Unable to create a token")
	}

	// Start an authentication session (stateless) for the user
	//if err = s.CreateAuthSession(user.UUID, token); err != nil {
	//return nil, errors.New("Unable to authenticate")
	//}

	return token, nil
}

// CreateToken creates a authentication token for the user
func (s *AuthService) CreateToken(userUUID uuid.UUID) (
	*models.AuthenticationToken,
	error,
) {
	var err error
	token := models.AuthenticationToken{}

	// Define Access claims and expiry
	token.AccessExpires = time.Now().Add(time.Minute * 60).Unix()
	token.AccessUUID = uuid.Must(uuid.NewRandom()).String()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uuid"] = token.AccessUUID
	claims["user"] = userUUID.String()
	claims["expires"] = token.AccessExpires
	// Create Access claims and singed token
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.AccessToken, err = access.SignedString([]byte(core.Config("secret")))
	if err != nil {
		return nil, err
	}

	// Define Refresh claims and expiry
	token.RefreshExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	token.RefreshUUID = uuid.Must(uuid.NewRandom()).String()
	claims = jwt.MapClaims{}
	claims["uuid"] = token.RefreshUUID
	claims["expires"] = token.RefreshExpires
	claims["user"] = userUUID.String()
	// Create Refresh claims and singed token
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.RefreshToken, err = refresh.SignedString([]byte(core.Config("secret")))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// ValidateTokenAndAuthenticate is the helper for auth guard middleware, it first
// verifies the token and then check if the claims are valid for the authentication.
func (s *AuthService) ValidateTokenAndAuthenticate(tokenstr string) (*models.User, error) {
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

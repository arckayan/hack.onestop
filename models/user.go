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

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"type:varchar(36)"`
	Name     string    `gorm:"default:null" form:"name" binding:"required"`
	Email    string    `gorm:"unique;" form:"email" binding:"required"`
	Password string    `form:"password" binding:"required"`

	// Relationships
}

// BeforeCreate is a event hook provided by gorm, all the operations specified
// below are performed before creating a new user.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UUID for the model
	u.UUID = uuid.Must(uuid.NewRandom())

	// Hash user's password before saving
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hashed)

	return err
}

// Transform is the formatter for user model, the fields specified below are
// returned to the client as JSON
func (u *User) Transform() map[string]interface{} {
	return map[string]interface{}{
		"uuid":       u.UUID,
		"name":       u.Name,
		"email":      u.Email,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}

// Authenticate matches the input password with the uesr's original password
func (u *User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

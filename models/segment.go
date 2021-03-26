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
	"gorm.io/gorm"
)

type Segement struct {
	gorm.Model
	UUID   uuid.UUID `gorm:"type:varchar(36)"`
	TripID uint
}

// BeforeCreate is a event hook provided by gorm, all the operations specified
// below are performed before creating a new user.
func (s *Segement) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UUID for the model
	s.UUID = uuid.Must(uuid.NewRandom())

	return nil
}

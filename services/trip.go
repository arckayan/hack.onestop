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

	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/models"
)

type TripService struct{}

func (s *TripService) Create(t *models.Trip) error {
	core.K.DB.Engine.Create(t)

	if t.ID == 0 {
		return errors.New("Could not create a user.")
	}

	return nil
}

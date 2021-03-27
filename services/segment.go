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

type Segment struct {
}

func (s *Segment) Find(UUID string) (*models.Segment, error) {
	var segment models.Segment

	if err := core.K.DB.Engine.Where("uuid = ?", UUID).First(&segment).Error; err != nil {
		return nil, errors.New("Segment does not exist.")
	}

	return &segment, nil
}

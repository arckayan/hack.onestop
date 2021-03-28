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
	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/models"
)

type Airport struct{}

func (s *Airport) Find(ID int) (*models.Airport, error) {
	var airport models.Airport

	if err := core.K.DB.Engine.Model(&airport).Where("ID = ?", ID).First(&airport).Error; err != nil {
		return nil, err
	}

	return &airport, nil
}

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

package seeds

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/models"
)

type Airport struct{}

func (s *Airport) Run() {
	resp, err := http.Get(core.Config("airportURL"))
	if err != nil {
		panic("Error while reading the airport database file.")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	airports := []models.Airport{}
	json.Unmarshal([]byte(body), &airports)

	core.K.DB.Engine.Create(&airports)
}

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

package core

import (
	"github.com/kalkayan/onestop/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Engine *gorm.DB
}

func (d *Database) Register() {
	engine, err := gorm.Open(mysql.Open(Config("dbDNS")), &gorm.Config{})
	if err != nil {
		panic("Database connection failed")
	}
	d.Engine = engine

	if Conf["debug"].(bool) {
		println("Migrating models")
		d.Engine.AutoMigrate(&models.User{}, &models.Airport{})
	}
}

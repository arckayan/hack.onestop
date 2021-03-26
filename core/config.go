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
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Conf map[string]interface{}

// Configure loads the environment variables into the application
func Configure() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file, please check for .env file in the root directory.")
	}

	// Initiate the map
	Conf = make(map[string]interface{})

	// This value is the name of this appliation, the value will be used when
	// the server needs to place appliation's name.
	Conf["name"] = os.Getenv("APP_NAME")

	Conf["secret"] = os.Getenv("APP_SECRET")

	// This value is used to check if the application should do the logging of
	// events and errors.
	Conf["debug"], _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))

	// database string
	Conf["dbDNS"] = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	// Redis
	Conf["redisDNS"] = fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	Conf["redisPassword"] = os.Getenv("REDIS_PASSWORD")

	Conf["airportURL"] = os.Getenv("URL_AIRPORT")
}

func Config(key string) string {
	return Conf[key].(string)
}

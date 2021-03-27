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

var C map[string]interface{}

// Configure loads the environment variables into the application
func Configure() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file, please check for .env file in the root directory.")
	}

	// Initiate the map
	C = make(map[string]interface{})

	// This value is the name of this appliation, the value will be used when
	// the server needs to place appliation's name.
	C["name"] = os.Getenv("APP_NAME")

	// Secret used to sign stuff
	C["secret"] = os.Getenv("APP_SECRET")

	// This value is used to check if the application should do the logging of
	// events and errors.
	C["debug"], _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))

	// database string
	C["dbDNS"] = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	// Redis
	C["redisDNS"] = fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	C["redisPassword"] = os.Getenv("REDIS_PASSWORD")

	C["airportURL"] = os.Getenv("URL_AIRPORT")

	C["kiwiURL"] = os.Getenv("KIWI_URL")
	C["kiwiAPIKey"] = os.Getenv("KIWI_KEY")
}

func Config(key string) string {
	return C[key].(string)
}

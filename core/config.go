/*
Copyright 2021 © The Onestop Authors

All Rights Reserved.

NOTICE: All information contained herein is, and remains the property of
The Onestop Authors. The intellectual and technical concepts contained
herein are proprietary to The Onestop Authors. Dissemination of this
information or reproduction of this material is strictly forbidden unless
prior written permission is obtained from The Onestop Authors.

Authors: Manish Sahani			<rec.manish.sahani@gmail.com>
		 Priyadarshan Singh		<singhpd75@gmail.com>

*/

package core

import "github.com/joho/godotenv"

// Configure loads the environment variables into the application
func Configure() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file, please check for .env file in the root directory.")
	}
}

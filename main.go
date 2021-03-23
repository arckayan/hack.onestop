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

package main

import (
	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/routes"
)

/*
                        _
                       | |
   ___  _ __   ___  ___| |_ ___  _ __
  / _ \| '_ \ / _ \/ __| __/ _ \| '_ \
 | (_) | | | |  __/\__ \ || (_) | |_) |
  \___/|_| |_|\___||___/\__\___/| .__/
                                | |
                                |_|

Reinventing traveling with one-stop traveling
*/

func main() {
	// Configure the core of the application -- this includes loading the
	// application environment variables, setting some flags, etc.
	core.Configure()

	// We need kernel instance to listen the requests, kernel is the component
	// that binds all the services and providers in one place.
	kernel := core.Boot()

	// This registers all the providers of the kernel
	kernel.Register()

	// routes are registered to the application router. the further addition of
	// routes  should be done in Register method under "routes.go" and not here.
	routes.Register(&kernel.Router)

	// Once the kernel is booted and all the services are initiated we run the
	// kernel at a defined/default port and listen for the incoming requests.
	kernel.Run()
}

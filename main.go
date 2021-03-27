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
	"flag"

	"github.com/kalkayan/onestop/core"
	"github.com/kalkayan/onestop/models"
	"github.com/kalkayan/onestop/routes"
	"github.com/kalkayan/onestop/seeds"
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
	// options for the servers
	seed := flag.Bool("seed", false, "Seed the database")
	migrate := flag.Bool("migrate", false, "Enables auto-migration of the database schemas")

	flag.Parse()

	// Configure the core of the application -- this includes loading the
	// application environment variables, setting some flags, etc.
	core.Configure()

	// We need kernel instance to listen the requests, kernel is the component
	// that binds all the services and providers in one place.
	kernel := core.Boot()

	// This registers all the providers of the kernel
	kernel.Register()

	// Migrate the model schemas to the database, can only be runned if the
	// application is in develpment mode and `--migrate` option is provided
	if *migrate && core.C["debug"].(bool) {
		kernel.DB.Engine.AutoMigrate(
			&models.User{},
			&models.Airport{},
			&models.Trip{},
			&models.Location{},
			&models.Segment{},
			&models.Flight{},
			// Add Schema Models to the list
		)
	}

	// Seed the database tables by running the seeder, (This should be done at-
	// least one time for filling airports table)
	if *seed {
		// Custom seeders are defined under the `seeds` folder and are called
		// in the `seeds.go` file, if you need to create a new seeder, create
		// a file with Run implemented and call it from `seeds.go` (NOT HERE).
		seeds.Run()
	}

	// routes are registered to the application router. the further addition of
	// routes  should be done in Register method under "routes.go" and not here.
	routes.Register(&kernel.Router)

	// Once the kernel is booted and all the services are initiated we run the
	// kernel at a defined/default port and listen for the incoming requests.
	kernel.Run()
}

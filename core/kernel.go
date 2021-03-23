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

// Kernel binds all the services and providers in one place
type Kernel struct {
	Router Router
	DB     Database
}

var K Kernel

// Boot binds all the providers of the application
func Boot() *Kernel {
	//
	router := Router{}

	//
	DB := Database{}

	//
	K = Kernel{router, DB}

	return &K
}

// Register initiates the kernel's services and binds them with the kernel
func (k *Kernel) Register() {
	//
	k.Router.Register()

	//
	k.DB.Register()
}

// Run listen for the incoming requests and process them
func (k *Kernel) Run() {
	k.Router.Run()
}

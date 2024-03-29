// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func Initialize() string {
	myFooer := provideMyFooer()
	string2 := provideBar(myFooer)
	return string2
}

// wire.go:

var PSet = wire.NewSet(
	provideBar, wire.Bind(new(Fooer), new(*MyFooer)), provideMyFooer,
)

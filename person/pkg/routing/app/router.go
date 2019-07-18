package app

import (
	routing "github.com/go-ozzo/ozzo-routing"
)

// Register .
func Register(router *routing.RouteGroup) {
	{
		handler := NewPersonHandler()
		router.Post("/person", handler.CreatePerson)
	}

}

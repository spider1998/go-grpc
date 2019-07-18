package app

import (
	`git.sdkeji.top/share/sdlib/api`
	routing `github.com/go-ozzo/ozzo-routing`
	`sdkeji/person/pkg/code`
	`sdkeji/person/pkg/entity`
	`sdkeji/person/pkg/service`
)

func NewPersonHandler() PersonHandler {
	return  PersonHandler{}
}

type  PersonHandler struct{}


func (p PersonHandler)CreatePerson(c *routing.Context) error {
	var req entity.CreatePersonRequest
	err := c.Read(&req)
	if err != nil {
		return code.Error(api.InvalidData).WithDetails(err)
	}
	person,err := service.Person.CreatePerson(req)
	if err != nil {
		return err
	}
	return c.Write(person)
}

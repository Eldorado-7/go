package Employees

import (
	"go-microservices/Lib/Controllers"
	"go-microservices/Lib/Providers/Factory"
)

type Employee struct {
	Controllers.Controllers
}

func (this Employee) DoRun(params []string) string {
	Factory.CreateProvider().List()

	//Create a JSON encoded result
	return ""
}

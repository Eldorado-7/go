package Models

import (
	"go-microservices/Lib/Controllers"
	_ "go-microservices/Models/Employees"
)

type EmployeeList struct {
	Controllers.Controller
}

func (this EmployeeList) DoRun(params []string) string {
	//model := EmployeeModel{}
	//model
	//Creaye a new provider and inject into controller provider using Factory class pattern
	//this.Provider =
	//empMOdel :=
	//Fetch the list of Data from data provider
	//list := this.Provider.List()
	//Create a JSON encoded result
	return ""
}

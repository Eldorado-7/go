package Models

import (
	"go-microservices/Lib/Controllers"
)

type EmployeeList struct {
	Controllers.Controller
}

func (this EmployeeList) DoRun(params []string) string {
	//Creaye a new provider and inject into controller provider using Factory class pattern
	//this.Provider =

	//Fetch the list of Data from data provider
	//list := this.Provider.List()
	//Create a JSON encoded result
	return ""
}

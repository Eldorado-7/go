package Employees

import (
	Entities "go-microservices/Entities/Employees"
	"go-microservices/Lib/Providers/Interface"
)

type EmployeeModel struct {
	Provider Interface.ProviderInterface
}

func (this EmployeeModel) List() []Entities.Employee {
	//Stub method. must return list from its provder
	return make([]Entities.Employee, 0)
}

func (this EmployeeModel) Insert() bool {
	//Stub method. must return list from its provder
	return true
}

func (this EmployeeModel) Delete() bool {
	//Stub method. must return list from its provder
	return true
}

func (this EmployeeModel) Update() bool {
	//Stub method. must return list from its provder
	return true
}

func (this EmployeeModel) Find() Entities.Employee {
	//Stub method. must return list from its provder
	return Entities.Employee{}
}

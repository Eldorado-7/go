package Employees

import (
	"go-microservices/Controllers/Employees"
	"go-microservices/Lib/Providers/Interface"
)

type EmployeeListProvider struct {
	Interface.ProviderInterface
}

func (this EmployeeListProvider) List() []Employees.Employee {
	//Stub method. must return list from its provder
	return make([]Employees.Employee, 0)
}

func (this EmployeeListProvider) Insert() bool {
	//Stub method. must return list from its provder
	return true
}

func (this EmployeeListProvider) Delete() bool {
	//Stub method. must return list from its provder
	return true
}

func (this EmployeeListProvider) Update() bool {
	//Stub method. must return list from its provder
	return true
}

func (this EmployeeListProvider) Find() Employees.Employee {
	//Stub method. must return list from its provder
	return Employees.Employee{}
}

package main

import "go-microservices/Controllers/Employees"

func testAppEngine() {
	//TODO: here we have to ascertain about the object creational of the whole framework
	var engin = Employees.EmployeeList{}
	engin.Run(make([]string, 0))
}
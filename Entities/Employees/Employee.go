package Entities

import (
	"go-microservices/Lib/Entities"
)

type Employee struct {
	Entities.Entity
	Name   string
	Family string
	Email  string
	Phone  string
}

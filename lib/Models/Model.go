package Models

import (
	"go-microservices/lib/Providers/Factory"
	"go-microservices/lib/Providers/Interface"
)

type Model struct {
	Provider *Interface.ProviderInterface
}

func (this Model) Init() {
	provder := Factory.CreateProvider()
	this.Provider = &provder
	// this.Provider = Factory.CreateProvider()
}

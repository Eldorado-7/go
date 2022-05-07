package Controllers

import (
	"go-microservices/lib/Controllers/ControllerInterfaces"
	"go-microservices/lib/Controllers/Executors"
	"go-microservices/lib/Controllers/Executors/ExecutorInterfaces"
	"net/http"
)

type Controller struct {
	ControllerInterfaces.ControllerInterface
	executor ExecutorInterfaces.ExecutorInterface
}

//Create a generic http response helper for all incoming requests
func (this Controller) Run(params map[string]string) (http.HandlerFunc, error) {

	//Checking for the configuration of the execution engine
	if this.executor == nil {
		//TODO: check for the context of execution engine {Web, CLI, Microservice}
		this.executor = Executors.Executor{
			target = this
		}
	}
	
	// return func(rw http.ResponseWriter, r *http.Request) {

	// 	rw.Header().Add("content-type", "application/json")
	// 	rw.WriteHeader(http.StatusFound)

	// 	//Wrap the response into the http header as sequence of bytes
	// 	//rw.Write([]byte(this.DoRun(params)))
	// }, nil
	return nil, nil
}

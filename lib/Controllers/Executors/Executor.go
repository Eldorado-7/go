/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
 */

package Executors

import (
	"go-microservices/lib/Controllers/ControllerInterfaces"
	"go-microservices/lib/Controllers/Executors/ExecutorInterfaces"
	"net/http"
)

type Executor struct {
	ExecutorInterfaces.ExecutorInterface
	target      *ControllerInterfaces.ControllerInterface
	ContextType string
}

func (this *Executor) SetTarget(target ControllerInterfaces.ControllerInterface) {
	this.target = &target
}

func (this *Executor) GetTarget() ControllerInterfaces.ControllerInterface {
	return *this.target
}

func (this *Executor) Execute(params map[string]string) (bool, error) {
	//TODO: exectue context gateway action
	//this.target.Run(params)
	return true, nil
}
func (this *Executor) ExecuteAsHandler(params map[string]string) (http.HandlerFunc, error) {
	//TODO: exectue context gateway action
	var er error = nil
	return func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		res, err := this.GetTarget().ProcessMe(params)
		er = err
		//Wrap the response into the http header as sequence of bytes
		rw.Write([]byte(res))

	}, er
}
func (this *Executor) ExecuteAsync(params map[string]string) (bool, error) {
	//TODO: exectue context gateway action
	//this.target.Run(params)
	return true, nil
}
func (this *Executor) initiateEngine(context ControllerInterfaces.ControllerInterface) {
	this.target = &context
}

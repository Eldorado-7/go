package controllers

type runner func(params []string) struct{}

type Controller struct {
}

func (this Controller) Run(params []string) {
	this.DoRun(params)
}

func (this Controller) DoRun(params []string) {

}

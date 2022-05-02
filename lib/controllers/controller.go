package Controllers

type DoRun func([]string)

type Controller struct {
	DoRun
}

func (this Controller) Run(params []string) {
	this.DoRun(params)
}

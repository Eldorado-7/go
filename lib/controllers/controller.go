package Controllers

type DRun func([]string)

ype Controller struct {
	DoRun
}

func (this Controller) Run(params []string) {
	this.DoRun(params)
}

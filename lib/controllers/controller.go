package Controllers

type DoRun func([]string)

type Controllers struct {
	DoRun
}

func (this Controllers) Run(params []string) {
	this.DoRun(params)
}

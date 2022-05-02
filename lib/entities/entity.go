package entities

type Entity struct {
	Id int
}

func (this Entity) getId() int {
	return this.Id
}

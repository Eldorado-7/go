package MySql

type MySqlProvider struct {
	host       string
	user       string
	password   string
	datasource string
}

//TODO: find a way to implement all of Provider Interface methods

func (this MySqlProvider) Find() {
}

func (this MySqlProvider) List() {
}

func (this MySqlProvider) Insert() bool {
	return true
}

func (this MySqlProvider) Update() bool {
	return true
}

func (this MySqlProvider) Delete() bool {
	return true
}

func (this MySqlProvider) Connect() bool {
	return true
}

//Setters
func (this MySqlProvider) SetUser(value string) {
	this.user = value
}
func (this MySqlProvider) SetPassword(value string) {
	this.password = value
}

func (this MySqlProvider) SetHost(value string) {
	this.password = value
}

func (this MySqlProvider) SetDatasource(value string) {
	this.password = value
}

//Getters
func (this MySqlProvider) GetUser() string {
	return this.user
}
func (this MySqlProvider) GetPassword() string {
	return this.password
}

func (this MySqlProvider) GetHost() string {
	return this.host
}

func (this MySqlProvider) GetDatasource() string {
	return this.datasource
}

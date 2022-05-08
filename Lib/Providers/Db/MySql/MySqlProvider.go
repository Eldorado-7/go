package MySql

import (
	"database/sql"
)

type MySqlProvider struct {
	host       string
	user       string
	password   string
	datasource string
	connection *sql.DB
}

//TODO: find a way to implement all of Provider Interface methods

func (this MySqlProvider) Query(params map[string]string) {

}

func (this MySqlProvider) Connect() bool {
	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/test")
	this.connection = db
	if err != nil {
		return false
	} else {
		return true
	}
}

//Setters
func (this *MySqlProvider) SetUser(value string) {
	this.user = value
}
func (this *MySqlProvider) SetPassword(value string) {
	this.password = value
}

func (this *MySqlProvider) SetHost(value string) {
	this.password = value
}

func (this *MySqlProvider) SetDatasource(value string) {
	this.password = value
}

//Getters
func (this *MySqlProvider) GetUser() string {
	return this.user
}
func (this *MySqlProvider) GetPassword() string {
	return this.password
}

func (this *MySqlProvider) GetHost() string {
	return this.host
}

func (this *MySqlProvider) GetDatasource() string {
	return this.datasource
}

package postgre

import (
	"github.com/jinzhu/gorm"
	"orange/db/dber"
)

type postgreDB struct {
	dber.Parameter
	db *gorm.DB
}

func New(parameter dber.Parameter)dber.DB{
	return &postgreDB{
		Parameter:dber.Parameter{
			IP: parameter.IP,
			Port: parameter.Port,
			User: parameter.User,
			Password: parameter.Password,
			DBName: parameter.DBName,
		},
	}
}
func (this *postgreDB)Init(){

}
func (this *postgreDB)Connect()(err error){
	dataSourceName := "host=" + this.IP + "  user=" + this.User + " dbname=" +this.DBName + "  sslmode=disable  password=" + this.Password
	this.db, err = gorm.Open("postgres", dataSourceName)
	if err != nil {
		return
	}
	this.db.DB().SetMaxIdleConns(10)
	this.db.DB().SetMaxOpenConns(100)
	this.db.SingularTable(true) //全局禁用表名复数,使用TableName设置的表名不受影响
	return nil
}
func (this *postgreDB)Close()error{
	return nil
}
func (this *postgreDB)DB()*gorm.DB{
	return this.db
}
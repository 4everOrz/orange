package mysql

import (
	"github.com/jinzhu/gorm"
	"orange/db/dber"
)

type mysqlDB struct {
    dber.Parameter
	db *gorm.DB
}

func New(parameter dber.Parameter)dber.DB{
	return &mysqlDB{
		Parameter:dber.Parameter{
			IP: parameter.IP,
			Port: parameter.Port,
			User: parameter.User,
			Password: parameter.Password,
			DBName: parameter.DBName,
	},
	}
}
func (this *mysqlDB)Init(){

}
//创建连接
func (this *mysqlDB)Connect()(err error){
	dataSourceName := this.User + ":" + this.Password + "@tcp(" +
		this.IP + ":" + this.Port + ")/" +
		this.DBName + "?charset=utf8"
	this.db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		return
	}
	this.db.DB().SetMaxIdleConns(10)
	this.db.DB().SetMaxOpenConns(100)
	this.db.SingularTable(true) //全局禁用表名复数
	return nil
}
//关闭连接
func (this *mysqlDB)Close()error{
	return 	this.db.DB().Close()
}
//获取db对象
func (this *mysqlDB)DB()*gorm.DB{
	return this.db
}
package postgre

import (
	"github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"orange/db/param"
	"time"
)

type postgreDB struct {
	param.Parameter
	db *gorm.DB
}

func New(parameter param.Parameter)param.DB{
	return &postgreDB{
		Parameter:param.Parameter{
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
	go func (){
		for{
			if err:=this.db.DB().Ping();err!=nil{
				log4go.Error(err.Error())
			}
			time.Sleep(1*time.Minute)
		}
	}()
	return nil
}
func (this *postgreDB)Close()error{
	return nil
}
func (this *postgreDB)DB()*gorm.DB{
	return this.db
}
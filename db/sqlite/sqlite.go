package sqlite

import (
	"github.com/jinzhu/gorm"
	"orange/db/dber"
)

type sqliteDB struct {
	dber.Parameter
	db *gorm.DB
}
func New(parameter dber.Parameter)dber.DB{
	return &sqliteDB{
		Parameter:dber.Parameter{
			DBName: parameter.DBName,
		},
	}
}
func (this *sqliteDB)Init(){

}
func (this *sqliteDB)Connect()(err error){
	this.db, err = gorm.Open("sqlite3", this.DBName)
    if err!=nil{
     	return
    }
	return nil
}
func (this *sqliteDB)Close ()error{
	this.db.Close()
	return nil
}
func (this *sqliteDB)DB()*gorm.DB{
	return this.db
}
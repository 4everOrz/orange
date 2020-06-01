package sqlite

import (
	"github.com/jinzhu/gorm"
	"orange/db/param"
)

type sqliteDB struct {
	param.Parameter
	db *gorm.DB
}
func New(parameter param.Parameter)param.DB{
	return &sqliteDB{
		Parameter:param.Parameter{
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
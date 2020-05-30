package dber

import "github.com/jinzhu/gorm"

type DB interface {
	Init()
	Connect()error
	DB()*gorm.DB
	Close()error
}
type Parameter struct {
	IP string
	Port string
	User string
	Password string
	DBName  string
}
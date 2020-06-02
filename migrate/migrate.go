package migrate

import (
	"github.com/jinzhu/gorm"
	"orange/controller/article"
	"orange/controller/user"
)
type Migrate interface {
	AutoMigrate()error
}
func New(db *gorm.DB)Migrate{
	return &migrate{
		db: db,
	}
}
type migrate struct {
	db  *gorm.DB
}

func (this *migrate)AutoMigrate()error{
	models:=[]interface{}{
		&user.User{},
		&article.Article{},
	}
	this.db.AutoMigrate(models...)
	return nil
}
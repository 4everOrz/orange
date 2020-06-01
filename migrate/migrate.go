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
	this.db.AutoMigrate(&user.User{},&article.Article{})
	return nil
}
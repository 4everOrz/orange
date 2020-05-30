package runtime

import (
	"github.com/jinzhu/gorm"
	"orange/controller/article"
	"orange/controller/user"
)

type Runtime interface{
	Init()
	DB()*gorm.DB
	User()user.Controller
	Article()article.Controller
}
type runtime struct {
	db *gorm.DB
	user  user.Controller
	article article.Controller

}
func New(db *gorm.DB)Runtime{
	return &runtime{
		db: db,
	}
}
func (this *runtime)Init()  {
	this.user=user.NewUser(this.db)
	this.article=article.NewArticle(this.db)
}
func (this *runtime)DB()*gorm.DB{
	return this.db
}
func (this *runtime)User()user.Controller{
	return this.user
}
func (this *runtime)Article()article.Controller{
	return this.article
}
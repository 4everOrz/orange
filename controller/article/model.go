package article

import "github.com/jinzhu/gorm"

type Article struct {
	ID    string    `gorm:"column:id;type:varchar(20);primary_key;not null" json:"id"`
	Title string    `gorm:"column:title;type:varchar(100);index:idx_articles_title;not null" json:"title"`
	Content string  `gorm:"column:content;type:text" json:"content"`
	Writer  string  `gorm:"column:writer;type:varchar(50);index:idx_articles_writer" json:"writer"`
	Creator string  `gorm:"column:creator;type:varchar(20);index:idx_articles_creator;not null" json:"creator"`
	CreatedAt int64 `gorm:"column:created_at;type:int8;index:idx_articles_created_at;not null" json:"createdAt"`
	UpdatedAt int64 `gorm:"column:updated_at;type:int8;index:idx_articles_updated_at;not null" json:"updatedAt"`
}
func (this *Article)TableName()string{
	return "articles"
}
type ArticleList struct {
	List  []Article `json:"list"`
	Total  int   `json:"total"`
	Count  int   `json:"count"`
}
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//数据库操作对象
type Model interface {
	//新增
	Add(db *gorm.DB,cmd *AddCmd)error
//	Add(db *gorm.DB,cmd *AddCmd)error
	//查找
	//Get(db *gorm.DB,id,loginName string)(userEntity *User,err error)
	//记录是否存在
	//Existed(db *gorm.DB,id,loginName string) bool
	//删除
	//Delete(db *gorm.DB,ids []string) error
	//修改
	//Update(db *gorm.DB,cmd *UpdateCmd)error
	//查询列表
	//Query(db *gorm.DB,cmd *QueryCmd)(UserList,error)

}
func NewModel()Model{
	return &model{}
}
type model struct {

}
func (this *model)Add(db *gorm.DB,cmd *AddCmd)error{
	article:=&Article{
		ID: cmd.ID,
		Title: cmd.Title,
		Writer: cmd.Writer,
		Creator: cmd.Creator,
		Content: cmd.Content,
	}
  return  db.Create(article).Error
}

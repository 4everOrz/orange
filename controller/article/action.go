package article

import (
	"errors"
	"orange/common/uuid"
	"strings"
)

type AddCmd struct {
	ID    string `json:"-"`
	Writer string `json:"writer"`
	Title string `json:"title"`
	Creator string `json:"-"`
	Content string `json:"content"`
}
func (this *AddCmd)Normalize(){
	this.ID=uuid.New()
	this.Title=strings.TrimSpace(this.Title)
	this.Creator=strings.TrimSpace(this.Creator)
}
func (this *AddCmd)Validate()error{
	this.Normalize()
	if len(this.Title)==0{
		return errors.New("缺少标题")
	}
	if len(this.Creator)==0{
		return errors.New("创建者信息异常")
	}

	return nil
}
type QueryCmd struct {
	Writer   string
	UpdatedFrom int64
	UpdatedTo  int64
	CreatedFrom   int64
	CreatedTo    int64
	CreatorName  string
}
func (this *QueryCmd)Normalize(){

}
func (this *QueryCmd)Validate()error{
	this.Normalize()

	return nil
}
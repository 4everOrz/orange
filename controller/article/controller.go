package article

import (
	//"github.com/alecthomas/log4go"
	"github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
)

type Controller interface{
	//新增
	Add(cmd *AddCmd)error
	//获取
	//Get(id,loginName string)(*User ,error)
	//删除
	//Delete(ids []string)error
	//修改
	//Update(cmd *UpdateCmd)error
	//查询列表
	//Query(cmd *QueryCmd)(UserList,error)
}
func NewArticle(db *gorm.DB)Controller{
	return &article{
		Model: NewModel(),
		DB: db,
	}
}

type article struct {
	Model Model
	DB  *gorm.DB
}


func (this *article)Add(cmd *AddCmd)error{
	if err := cmd.Validate(); err != nil {
		log4go.Error(err.Error())
		return err
	}
	return  this.Model.Add(this.DB,cmd)
}
/*
func (this *user)Get(id,loginName string)(userEntity *User,err error){
	return this.Model.Get(this.DB,id,loginName)
}

func (this *user)Delete(ids []string)error{
	return	this.Model.Delete(this.DB,ids)
}

func (this *user)Update(cmd *UpdateCmd)error{
	if err:=cmd.Validate();err!=nil{
		return err
	}
	if cmd.LoginName!=""{
		if this.Model.Existed(this.DB,"",cmd.LoginName){
			return errors.New("用户名已存在")
		}
	}
	return this.Model.Update(this.DB,cmd)
}
func (this *user)Query(cmd *QueryCmd)(UserList,error){
	if err:=cmd.Validate();err!=nil{
		return UserList{},err
	}
	return this.Model.Query(this.DB,cmd)
}*/
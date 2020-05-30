package user

import (
	"errors"
	"orange/common/uuid"
	"orange/controller"
	"security"
	"strings"
)

//	Normalize()
//	Validate() error
type AddCmd struct {
	ID          string  `json:"-"`
	LoginName    string `json:"loginName"`
	Password     string `json:"password"`
	Email        string `json:"email" `
	Role         int    `json:"role"` //1:管理员2:操作员
	//CreatedAt    int64  `json:"createdAt"`
	//UpdatedAt    int64  `json:"updatedAt"`
	Enable        int   `json:"enable"` //1:正常  2:禁用
}

//参数归化
func(this *AddCmd)Normalize(){
     this.ID=uuid.New()
     this.LoginName=strings.TrimSpace(this.LoginName)
     this.Password=strings.TrimSpace(this.Password)
}
//参数检验
func(this *AddCmd)Validate()error{
	this.Normalize()
	if len(this.LoginName)<5{
		return errors.New("登录名长度非法")
	}
	if len(this.Password)<6{
		return errors.New("密码长度非法")
	}
	if strings.TrimSpace(this.Password)!=""{
		this.Password=security.Md5(this.Password)
	}else {
		return errors.New("密码格式错误")
	}
	switch this.Role {
	case 1,2:
		break
	case 0:
		this.Role=2
		break
	default:
		return errors.New("非法的用户角色")
	}

	switch this.Enable {
	case 0:
		this.Enable=1
		break
	case 1,2:
		break
	default:
		return errors.New("非法的使能参数")
	}
	return nil
}

type QueryCmd struct {
	CreatedFrom int64 `json:"createdFrom"`
	CreatedTo   int64  `json:"createdTo"`
	UpdatedFrom int64 `json:"updatedFrom"`
	UpdatedTo   int64 `json:"updatedTo"`
	Role       int    `json:"role"`
	Enable    int    `json:"enable"`
	controller.QueryCmd
}
func (this *QueryCmd)Normalize(){

}

func (this *QueryCmd)Validate()error{
	this.Normalize()
	sortMap:=map[string]string{
		"updatedAt":"updated_at",
		"createdAt":"created_at",
	}
	if this.Sort!=""{
	   	sort,ok:=sortMap[this.Sort]
	   	if !ok{
	   		return errors.New("不支持的排序字段")
		}
		this.Sort=sort
	}else {
		this.Sort,_=sortMap["updatedAt"]
	}
	if this.Order!=""{
		switch this.Order {
		case "desc","asc":
			break
		default:
			return errors.New("非法的排序规则")
		}
	}else {
		this.Order="asc"
	}
	return nil
}

type UpdateCmd struct {
     ID       string
	LoginName string  `json:"loginName"`
	Email    string   `json:"email"`
}
func (this *UpdateCmd)Normalize(){
	this.LoginName=strings.TrimSpace(this.LoginName)
	this.Email=strings.TrimSpace(this.Email)
}

func (this *UpdateCmd)Validate()error{
	this.Normalize()
	return nil
}
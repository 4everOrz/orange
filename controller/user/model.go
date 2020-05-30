package user

import (
	"github.com/jinzhu/gorm"
	"orange/controller/article"
)


//数据库模型
type User struct {
	ID           string `gorm:"column:id;type:varchar(20);primary_key;not null" json:"id"`
	LoginName    string `gorm:"column:login_name;type:varchar(20);index:idx_users_login_name" json:"loginName"`
	Password     string `gorm:"column:password;type:varchar(50);index:idx_users_password" json:"-"`
	Email        string `gorm:"column:email;type:varchar(100)" json:"email" `
	Role         int    `gorm:"column:role;type:int2;index:idx_users_role" json:"role"` //1:管理员2:操作员
	Token        string `gorm:"column:token;type:varchar(255)" json:"-"`
	Region       string `gorm:"column:region;type:varchar(100)" json:"region"`
	CreatedAt     int64  `gorm:"column:created_at;type:int8;index:idx_users_created_at" json:"createdAt"`
	UpdatedAt     int64 `gorm:"column:updated_at;type:int8;index:idx_users_updated_at" json:"updatedAt"`
	Enable        int   `gorm:"column:enable;type:int2;index:idx_users_enable" json:"enable"` //1:正常  2:禁用
	Articles     []article.Article `gorm:"foreignkey:creator;association_foreignkey:id" json:"articles,omitempty"`//一对多关系
}
//表名
func (this *User)TableName()string{
	return "users"
}
type UserList struct {
	List  []User `json:"list"`
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
	//查找
	Get(db *gorm.DB,id,loginName string)(userEntity *User,err error)
	//记录是否存在
	Existed(db *gorm.DB,id,loginName string) bool
	//删除
	Delete(db *gorm.DB,ids []string) error
	//修改
	Update(db *gorm.DB,cmd *UpdateCmd)error
	//查询列表
	Query(db *gorm.DB,cmd *QueryCmd)(UserList,error)

}
func NewModel()Model{
	return &model{}
}

type model struct {

}

func (this *model) Add(db *gorm.DB,cmd *AddCmd) error {
	user:=&User{
		ID: cmd.ID,
		LoginName:cmd.LoginName,
		Password: cmd.Password,
		Email: cmd.Email,
		Role: cmd.Role,
		//CreatedAt: cmd.CreatedAt,
		//UpdatedAt: cmd.UpdatedAt,
		Enable: cmd.Enable,
	}
	return db.Create(user).Error
}

func (this *model)Get(db *gorm.DB,id,loginName string)( *User,error){
	var (
	user User
	err  error
		)
	ql:=db.Model(&User{})

	if id!=""{
		ql=ql.Where("id = ?",id)
	}
	if loginName!=""{
		ql=ql.Where("login_name = ?",loginName)
	}
	err=ql.Select("*").First(&user).Error
	if err!=nil&&err==gorm.ErrRecordNotFound{
		//查询无记录
		return nil,nil
	}else if err!=nil{
		return nil,err
	}
	return &user,err
}
func (this *model)Existed(db *gorm.DB,id,loginName string)bool{
	ql:=db.Model(&User{})
	if id!=""{
		ql=ql.Where("id = ?",id)
	}
	if loginName!=""{
		ql=ql.Where("login_name = ?",loginName)
	}
    return !ql.First(&User{}).RecordNotFound()
}

func (this *model) Delete(db *gorm.DB,ids []string) error {
	return db.Model(&User{}).Where("id IN (?)", ids).Delete(User{}).Error
}

func (this *model)Query(db *gorm.DB,cmd *QueryCmd)(UserList,error){
	var  list UserList
	ql:=db.Model(&User{})
	if cmd.Enable!=0{
		ql=ql.Where("enable = ?",cmd.Enable)
	}
	if cmd.Role!=0{
		ql=ql.Where("role = ?",cmd.Role)
	}
	if cmd.CreatedFrom!=0||cmd.CreatedTo!=0{
		ql=ql.Where("created_at BETWEEN ? AND ?",cmd.CreatedFrom,cmd.CreatedTo)
	}
	if cmd.UpdatedFrom!=0||cmd.UpdatedTo!=0{
		ql=ql.Where("updated_at BETWEEN ? AND ?",cmd.UpdatedFrom,cmd.UpdatedTo)
	}
	ql=ql.Count(&list.Total)
	if cmd.Limit!=0{
		ql=ql.Limit(cmd.Limit).Offset(cmd.Offset)
	}
	err:=ql.Order(cmd.Sort+" "+cmd.Order).Find(&list.List).Error
	if err!=nil{
		return list,err
	}
	list.Count=len(list.List)
    return list,nil
}
func (this *model)Update(db *gorm.DB,cmd *UpdateCmd)error{
	ql:=db.Model(&User{}).Where("id = ?",cmd.ID)
	paramers := make(map[string]interface{})
	if cmd.LoginName!=""{
	   paramers["loginName"]=cmd.LoginName
	}
	if cmd.Email!=""{
		paramers["email"]=cmd.Email
	}
	return ql.Update(paramers).Error
}
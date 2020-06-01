package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"orange/common/config"
	"orange/db/mysql"
	"orange/db/param"
	"orange/db/postgre"
	"orange/db/sqlite"
	"time"
)

const (
	mysqlDirver ="mysql"
	postgreDirver="postgre"
	sqliteDirver="sqlite"
)


func New()param.DB{
	return &database{}
}
type database struct {
	dbtype  string
	param.Parameter
	db *gorm.DB
}

func (this *database)Init(){
	this.dbtype,_=config.GetString("db","DBType")
	this.IP, _ = config.GetString("db", "IP")
	this.Port,_=config.GetString("db","Port")
	this.User, _ = config.GetString("db", "User")
	this.Password, _ = config.GetString("db", "Password")
	this.DBName, _ = config.GetString("db", "DBName")
}
func (this *database)Connect()(err error){
	var dbEntity param.DB
	//创建数据库对象
	switch this.dbtype {
	case mysqlDirver:
		dbEntity=mysql.New(param.Parameter{IP: this.IP,Port: this.Port,User: this.User,Password: this.Password,DBName: this.DBName})
	case postgreDirver:
		dbEntity=postgre.New(param.Parameter{IP: this.IP,Port: this.Port,User: this.User,Password: this.Password,DBName: this.DBName})
	case sqliteDirver:
		dbEntity=sqlite.New(param.Parameter{DBName: this.DBName})
	default:
		return errors.New("不受支持的数据库类型")
	}
	//初始化
	dbEntity.Init()
	//创建连接
	if err=dbEntity.Connect();err!=nil{
		return
	}
	//获取数据库操作对象
	this.db=dbEntity.DB()
	//设置数据库回调函数
	CallBackReplace(this.db)
	return nil
}

func (this *database)DB()*gorm.DB{
	return this.db
}
//关闭连接
func (this *database)Close()error{
   return    this.db.Close()
}
//**************************************************************************************************************/
//替换掉原有库钩子回调,用于解决原生库updatedAt,createdAt默认存时间戳而不是int64的问题
func CallBackReplace(db *gorm.DB){
	db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete",deleteCallback)
}
//回调钩子-新建操作
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}
//回调钩子-更新操作
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}
//回调钩子-删除操作
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
package mysql

import (
	"GinWeb/common/helper"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age int
	Hobby string
}
/**
desc：添加数据
说明：可接收interface数据。如果为User，单条数据添加；[]User，批量添加
*/
func AddUser(user interface{})  {
	switch user.(type) {
	case User:
		//单条添加
		userFinal := user.(User)
		create := MysqlDb.Create(&userFinal)
		if nil != create.Error{
			helper.PolarLog.Warning(fmt.Sprintf("add user fail %v",create.Error))
		}
	case []User:
		//批量添加
		batchUserFinal := user.([]User)
		create := MysqlDb.Create(&batchUserFinal)
		if nil != create.Error{
			helper.PolarLog.Warning(fmt.Sprintf("batch add user fail %v",create.Error))
		}
	default:
		helper.PolarLog.Warning("Invalid data for user !")
	}


}
/**
通过主键id更新数据
 */
func UpdateUserById(id int,user User)  {
	update := MysqlDb.Model(&User{}).Where("id",id).Updates(user)
	if nil != update.Error{
		helper.PolarLog.Warning(fmt.Sprintf("update user fail %v",update.Error))
	}
}
/**
创建用户数据表
 */
func CreateUser()  {
	MysqlDb.Table("users")
	err := MysqlDb.AutoMigrate(&User{})
	if nil != err{
		helper.PolarLog.Warning(fmt.Sprintf("user table create fail %v",err))
		return
	}
	helper.PolarLog.Info("user table create success !")
}
/**
desc：根据主键id删除，软删除
说明：支持单条&多条删除
 */
func DeleteUserById(id interface{})  {
	switch id.(type) {
	case int:
		//单条删除
		tx := MysqlDb.Delete(&User{}, id.(int))
		if nil != tx.Error{
			helper.PolarLog.Warning(fmt.Sprintf("delete user fail %v",tx.Error))
			return
		}
	case []int:
		tx := MysqlDb.Delete(&User{}, id.([]int))
		if nil != tx.Error{
			helper.PolarLog.Warning(fmt.Sprintf("batch delete user fail %v",tx.Error))
			return
		}
	default:
		helper.PolarLog.Warning("Invalid data for user !")
	}
}
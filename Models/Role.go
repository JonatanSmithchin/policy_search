package Models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	Mysql "police_search/Databases"
	"strings"
)

//StringArray 角色部分
type StringArray []string

//Role 角色
type Role struct {
	Id          int          `gorm:"column:id" form:"id" json:"id" comment:"自增id" sql:"int(11),PRI"`
	Name        string       `gorm:"column:name" form:"name" json:"name" comment:"角色名" sql:"varchar(255)"`
	Description string       `gorm:"column:description" form:"description" json:"description" comment:"描述" sql:"varchar(255)"`
	Permission  *StringArray `gorm:"type:json;column:permission" form:"permission" json:"permission" comment:"权限"`
}

func (data *StringArray) Scan(val interface{}) (err error) {
	if val == nil {
		return nil
	}
	if payload, ok := val.([]byte); ok {
		var value []string
		err = json.Unmarshal(payload, &value)
		if err == nil {
			*data = value
		}
	}
	return
}

func CheckRolePermission(roleId uint, permissionFunc string) bool {
	if roleId == 0 {
		return false
	}

	var myRole Role
	err := Mysql.DB.Where("id = ?", roleId).First(&myRole).Error
	if err != gorm.ErrRecordNotFound {
		fmt.Printf("%v", myRole)
		permissions := myRole.Permission
		permissions.Scan(permissions)
		for _, permission := range *permissions {
			fmt.Println("permissionFunc:", permissionFunc)
			fmt.Println("permission:", permission)
			if strings.HasPrefix(permissionFunc, permission) {
				return true
			}
		}
	}
	return false
}

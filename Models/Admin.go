package Models

import (
	Mysql "police_search/Databases"
)

type Admin struct {
	ID       int    `form:"id" json:"id" gorm:"primaryKey;column:uid"`
	Name     string `form:"name" json:"name" form:"name" gorm:"column:uname;type:varchar(32);default:(-)"`
	Password string `form:"password" json:"password" form:"password" gorm:"column:pwd;type:varchar(32);default:(-)"`
}

func GetAdminByName(name string) *Admin {

	admin := &Admin{}
	Mysql.DB.Where("uname=?", name).Find(admin)
	return admin

}

func GetAdmins() *[]*Admin {
	var admins []*Admin
	Mysql.DB.Find(&admins)
	return &admins
}

func (a *Admin) DeleteAdmin() error {
	result := Mysql.DB.Where("uname=?", a.Name).Delete(a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *Admin) AddAdmin() error {
	result := Mysql.DB.Create(a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

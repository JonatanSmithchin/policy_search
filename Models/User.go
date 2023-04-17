package Models

import "police_search/Databases"

type User struct {
	UserID   int    `json:"id" gorm:"primaryKey;column:UserID"`
	UserName string `form:"UserName" json:"UserName" gorm:"column:UserName;type:varchar(32);default:(-)"`
	UserPwd  string `form:"password" json:"password" gorm:"column:Password;type:varchar(32);default:(-)"`
	Age      int8   `json:"age" gorm:"column:Age;type:smallint"`
	Email    string `json:"email" gorm:"column:Email;type:varchar(64)"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) CreateUser() error {

	result := Mysql.DB.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) ChangeAge(age int) error {

	result := Mysql.DB.Model(&u).Update("Age", age)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) ChangeEmail(email string) error {
	result := Mysql.DB.Model(&u).Update("Email", email)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) ChangePassword(password string) error {
	result := Mysql.DB.Model(&u).Update("Password", password)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) DeleteUser() error {
	result := Mysql.DB.Delete(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByName(name string) *User {
	user := &User{}
	Mysql.DB.Where("UserName=?", name).Find(user)
	return user
}

func GetAllUsers() *[]*User {
	var users []*User
	Mysql.DB.Find(&users)
	return &users
}

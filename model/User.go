package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"labplatform/db"
	"labplatform/utils/errmsg"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchal(20);not null" json:"username,omitempty" validate:"required,min=2,max=12" label:"用户名"`
	password string `gorm:"type:varchal(100);not null" json:"password,omitempty" validate:"required,min=4,max=20" label:"密码"`
	Role     int    `gorm:"type:int;default 1;not null" json:"role,omitempty" validate:"required,gte=2" label:"角色"`
}

func CreateUser(data *User) int {
	//data.password = ScryptPw(data.password)
	err := db.DB.Create(data)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func CheckUser(name string) (code int) {
	var user User
	db.DB.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.DB.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCSE
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

func GetUser(id int) (User, int) {
	var user User
	err := db.DB.First(&user, id)
	if err != nil {
		return User{}, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.DB.Model(&user).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.DB.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	u.password = ScryptPw(u.password)
//	return nil
//}
//
//func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
//	u.password = ScryptPw(u.password)
//	return nil
//}

func ChangePassword(id int, data *User) int {
	err := db.DB.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func ScryptPw(password string) string {
	//const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}

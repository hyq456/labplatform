package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"labplatform/utils"
	"labplatform/utils/errmsg"
	"log"
	"math/rand"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username,omitempty" validate:"required,min=2,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(100);not null" json:"Password,omitempty" validate:"required,min=4,max=20" label:"密码"`
	Role     int    `gorm:"type:int;default 1;not null" json:"role,omitempty" validate:"required,gte=2" label:"角色"`
	Email    string `gorm:"type:varchar(20);not null" json:"email,omitempty" validate:"required,email" label:"邮箱"`
	status   string `gorm:"type:char(1);default 0;not null" json:"status"`
}

func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(data)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
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
	err := db.First(&user, id).Error
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
	maps["email"] = data.Email
	maps["status"] = data.status
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//func (u *User) BeforeCreate(tx *gorm.db) (err error) {
//	u.Password = ScryptPw(u.Password)
//	return nil
//}
//
//func (u *User) BeforeUpdate(tx *gorm.db) (err error) {
//	u.Password = ScryptPw(u.Password)
//	return nil
//}

func ChangePassword(id int, data *User) int {
	err := db.Select("Password").Where("id = ?", id).Updates(&data).Error
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

func CheckLogin(username string, password string) (User, int) {
	var user User
	//var PasswordErr error
	db.Where("username = ?", username).First(&user)

	//PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	//if PasswordErr != nil{
	if password != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCSE
}

func SendEmailValidate(email string) (string, int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	e := gomail.NewMessage()
	e.SetHeader("From", utils.Mailaddress)
	e.SetHeader("To", email)
	//e.SetAddressHeader("To",""
	e.SetHeader("Subject", "激活邮件")
	e.SetBody("text/html", vCode)

	d := gomail.NewDialer(utils.DialerAdress, 587, utils.DialerUser, utils.DialerPW)
	err := d.DialAndSend(e)
	if err != nil {
		log.Println(err)
		return "", errmsg.ERROR_SEND_FAIL
	}
	return vCode, errmsg.SUCCSE
}

//func ValidEmail(vCode string,email string) int {
//	vCodeRaw :=
//}

package v1

import (
	"github.com/gin-gonic/gin"
	"labplatform/model"
	"labplatform/utils/errmsg"
	"labplatform/utils/validator"
	"log"
	"net/http"
	"strconv"
	"time"
)

// AddUser 创建用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	vCode := c.Param("vCode")
	_ = c.BindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":  validCode,
			"message": msg,
		})
		c.Abort()
		return
	}

	codeKey := "VerityCode" + data.Email + ":Code"
	vCodeRaw, err := model.DbRedis.Get(model.Ctx, codeKey).Result()
	if err != nil || vCodeRaw != vCode {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_VCODE_NOT_EXIT,
			"message": errmsg.GetErrMsg(errmsg.ERROR_VCODE_NOT_EXIT),
		})
		c.Abort()
		return
	}
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCSE {
		model.EditUser(id, &data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// ChangeUserPassword 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.ChangePassword(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteUser(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func SendValidateCode(c *gin.Context) {
	em := c.Query("email")
	vCode, code := model.SendEmailValidate(em)
	if code != errmsg.SUCCSE {
		log.Println(errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, gin.H{
			"status": code,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}
	codeKey := "VerityCode" + em + ":Code"
	err := model.DbRedis.Set(model.Ctx, codeKey, vCode, time.Minute*5).Err()

	if err != nil {
		log.Println(errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, gin.H{
			"status": errmsg.ERROR_SAVE_VCODE,
			"msg":    errmsg.GetErrMsg(errmsg.ERROR_SAVE_VCODE),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "验证码发送成功",
		"status": 200,
		"vCode":  vCode,
	})
	return
}

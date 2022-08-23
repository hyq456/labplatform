package v1

import (
	"github.com/gin-gonic/gin"
	"labplatform/model"
	"labplatform/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var code int
	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrMsg(code),
			//"token":   token,
		})
	} else {
		//setToken(c, formData)
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrMsg(200),
			//"token":   token,
		})
	}

}

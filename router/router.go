package router

import (
	"github.com/gin-gonic/gin"
	"labplatform/api/v1"
	"labplatform/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	router := r.Group("api/v1")
	{
		//登录
		router.POST("login", v1.Login)
		//发送邮件
		router.GET("sendmail", v1.SendValidateCode)
		// 用户模块的路由接口
		router.POST("user/add/:vCode", v1.AddUser)
		//router.GET("admin/users", v1.GetUsers)
		router.GET("user/:id", v1.GetUserInfo)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//修改密码
		router.PUT("admin/changepw/:id", v1.ChangeUserPassword)
	}

	_ = r.Run(utils.HttpPort)

}

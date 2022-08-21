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
		router.POST("user/add", v1.AddUser)
		// 用户模块的路由接口
		//router.GET("admin/users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//修改密码
		router.PUT("admin/changepw/:id", v1.ChangeUserPassword)
	}
}

package router

import (
	v1 "yuki_book/controllers/v1"
	"yuki_book/middleware"
	"yuki_book/util/sign"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//全局 Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。
	r.Use(gin.Recovery())
	//全局 日志中间件
	//r.Use(middleware.LoggerToFile())
	//全局 跨域中间件
	r.Use(middleware.Cors())
	//加载模板文件
	//r.LoadHTMLGlob("yuki_book/templates/*")
	//加载静态文件
	//r.Static("/web", "yuki_book/static")
	//swagger文档
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//v1版本
	V1 := r.Group("con/v1")
	initAdminRouter(V1)
	initUserRouter(V1)

	return r
}

func initAdminRouter(V1 *gin.RouterGroup) {
	admin := V1.Group("/admin")
	{
		// 管理员注册
		admin.POST("/register", v1.AdminRegister)
		// 管理员登录
		admin.POST("/login", v1.AdminLogin)
		// 获取管理员信息
		admin.GET("/get", middleware.JWT(sign.AdminClaimsType), v1.AdminInfoGet)
		// 获取全部管理员信息
		admin.GET("/getAll", middleware.JWT(sign.AdminClaimsType), v1.AdminInfoGetAll)
		// 更改管理员密码
		admin.POST("/update", middleware.JWT(sign.AdminClaimsType), v1.AdminPasswordUpdate)
		// 更改管理员信息
		admin.POST("/updateInfo", middleware.JWT(sign.AdminClaimsType), v1.AdminInfoUpdate)
		// 注销管理员
		admin.POST("/delete", middleware.JWT(sign.AdminClaimsType), v1.AdminDelete)
	}
}

func initUserRouter(V1 *gin.RouterGroup) {
	user := V1.Group("/user")
	{
		// 用户注册
		user.POST("/register", v1.UserRegister)
		// 用户登录
		user.POST("/login", v1.UserLogin)
		// 获取用户信息
		user.GET("/get", middleware.JWT(sign.UserClaimsType), v1.UserInfoGet)
		// 更改用户密码
		user.POST("/update", middleware.JWT(sign.UserClaimsType), v1.UserPasswordUpdate)
		// 更改用户信息
		user.POST("/updateInfo", middleware.JWT(sign.UserClaimsType), v1.UserInfoUpdate)
		// 管理员更改用户类型
		user.POST("/updateUserType", middleware.JWT(sign.AdminClaimsType), v1.UserTypeUpdate)
		// 注销用户
		user.POST("/delete", middleware.JWT(sign.UserClaimsType), v1.UserDelete)
	}
}

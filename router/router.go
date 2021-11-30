package router

import (
	v1 "yuki_book/controller/v1"
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
	initUserTypeRouter(V1)
	initBookClassRouter(V1)
	initBookRouter(V1)
	initBookBorrowRouter(V1)

	return r
}

func initAdminRouter(V1 *gin.RouterGroup) {
	admin := V1.Group("/admin")
	{
		// 管理员注册
		admin.POST("/register", v1.AdminRegister)
		// 管理员登录
		admin.POST("/login", v1.AdminLogin)
		// 获取管理员部分信息
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
		// 管理员获取全部用户信息
		user.GET("/getAll", middleware.JWT(sign.AdminClaimsType), v1.UserInfoGetAll)
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

func initUserTypeRouter(V1 *gin.RouterGroup) {
	userType := V1.Group("/userType")
	{
		// 添加用户类型
		userType.POST("/new", v1.UserTypeNew)
		// 查看用户类型信息
		userType.GET("/get", v1.UserTypeGet)
		// 查看全部用户类型
		userType.GET("/getAll", v1.UserTypeGetAll)
		// 修改用户类型信息
		userType.POST("/update", v1.UserTypeUpdateInfo)
		// 删除用户类型
		userType.POST("/delete", v1.UserTypeDelete)

	}
}

func initBookClassRouter(V1 *gin.RouterGroup) {
	bookClass := V1.Group("/bookClass")
	{
		// 管理员添加书集
		bookClass.POST("/new", v1.BookClassNew)
		// 管理员删除书集
		bookClass.POST("/delete", v1.BookClassDelete)
		// 分页展示全部书集
		bookClass.GET("/getAll", v1.BookClassGetAll)
		// 通过id查找书集
		bookClass.GET("/getById", v1.BookClassGetById)
		// 管理员修改书集信息
		bookClass.POST("/updateInfo", v1.BookClassUpdateById)
	}
}

func initBookRouter(V1 *gin.RouterGroup) {
	book := V1.Group("/book")
	{
		// 管理员添加书本
		book.POST("/new", v1.BookNew)
		// 管理员删除书本
		book.POST("/delete", v1.BookDelete)
		// 通过id查找书本
		book.GET("/getById", v1.BookGetById)
		// 通过书集id查找书本
		book.GET("/getByBookClassId", v1.BookGetByClassId)
		// 管理员修改书本状态
		book.POST("/updateBookStatu", v1.BookUpdateStatu)
		// 管理员修改书本受损程度
		book.POST("/updateBookDamage", v1.BookUpdateDamage)
	}
}

func initBookBorrowRouter(V1 *gin.RouterGroup) {
	bookBorrow := V1.Group("/bookBorrow")
	{
		// 新建借书记录
		bookBorrow.POST("/new", v1.BookBorrowNew)
		// 查询借书记录
		bookBorrow.GET("/get", v1.BookBorrowGet)
		// 续借，修改借书记录
		bookBorrow.POST("/update", v1.BookBorrowUpdate)
		// 还书，修改借书记录
		bookBorrow.POST("/return", v1.BookBorrowReturnUpdate)
	}
}

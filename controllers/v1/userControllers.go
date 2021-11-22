package v1

import (
	"yuki_book/model/user_model"
	"yuki_book/service/user_service"
	"yuki_book/util/app"

	"github.com/gin-gonic/gin"
)

// @Tags 用户
// @Summary 用户注册
// @Description 用户注册
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /con/v1/user/register [post]
type UserRegisterBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserRegister(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body UserRegisterBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	user, _ := user_model.GetUserInfo(body.Username)
	if user != nil {
		appG.BadResponse("该账号已被注册")
		return
	}
	if appG.HasError(user_service.CreateUser(body.Username, body.Password)) {
		return
	}
	appG.SuccessResponse("注册成功")
}

// @Tags 用户
// @Summary 用户登录
// @Description 用户登录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/user/login [post]
type UserLoginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserLogin(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body AdminLoginBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	token, err := user_service.GenerateToken(body.Username, body.Password)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(token)

}

// @Tags 用户
// @Summary 获取用户信息
// @Description 获取用户信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/user/get [GET]
type UserInfoGetBody struct {
	Username string `json:"username" form:"username" validate:"required"`
}

func UserInfoGet(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body AdminInfoGetBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	user, err := user_model.GetUserInfoPart(body.Username)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(user)
}

// @Tags 用户
// @Summary 根据账号更改用户密码
// @Description 根据账号更改用户密码
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/user/update [post]
type UserPasswordUpdateBody struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	NewPassWord string `json:"newpassword" validate:"required"`
}

func UserPasswordUpdate(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body UserPasswordUpdateBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := user_model.GetUserInfo(body.Username)
	if err != nil {
		appG.BadResponse("未找到该账号")
		return
	}
	if appG.HasError(user_service.UpdateUserPassword(body.Username, body.Password, body.NewPassWord)) {
		return
	}
	appG.SuccessResponse("修改用户密码成功")
}

// @Tags 用户
// @Summary 根据账号修改户信息
// @Descrition 根据账号修用户信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @outer con/v1/user/updateInfo [post]
type UserInfoUpdateBody struct {
	Username        string  `json:"username" validate:"required"`
	Name            *string `json:"name"`
	UserPhonenumber *string `json:"user_phonenumb"`
	UserAddress     *string `json:"user_address"`
	UserClass       *string `json:"user_class"`
	UserEmail       *string `json:"user_Email"`
}

func UserInfoUpdate(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body UserInfoUpdateBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := user_model.GetUserInfo(body.Username)
	if err != nil {
		appG.BadResponse("未找到该用户账号")
		return
	}
	if appG.HasError(user_service.UpdateUserInfo(body.Username, body.Name, body.UserPhonenumber, body.UserAddress, body.UserClass, body.UserEmail)) {
		return
	}
	appG.SuccessResponse("修改用户信息成功")
}

// @Tags 用户
// @Summary 管理员根据账号修用户类型
// @Description 管理员根据账号修用户类型
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/user/updateUserType [post]
type UserTypeUpdateBody struct {
	Username string `json:"username" validate:"required"`
	Usertype string `json:"user_type"`
}

func UserTypeUpdate(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body UserTypeUpdateBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := user_model.GetUserInfo(body.Username)
	if err != nil {
		appG.BadResponse("未找到该用户账号")
		return
	}
	if appG.HasError(user_service.UpdateUserType(body.Username, body.Usertype)) {
		return
	}
	appG.SuccessResponse("修改用户信息成功")
}

// @Tags 用户
// @Summary 根据账号注销用户
// @Description 根据账号注销用户
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/admin/delete [post]
type UserDeleteBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserDelete(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body UserDeleteBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	if appG.HasError(user_service.DeleteUser(body.Username, body.Password)) {
		return
	}
	appG.SuccessResponse("注销用户成功")
}

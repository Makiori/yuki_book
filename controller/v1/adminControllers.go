package v1

import (
	"yuki_book/model/admin_model"
	"yuki_book/service/admin_service"
	"yuki_book/util/app"

	"github.com/gin-gonic/gin"
)

// @Tags 管理员
// @Summary 管理员注册
// @Description 管理员注册
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/admin/register [post]
type AdminRegisterBody struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

func AdminRegister(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body AdminRegisterBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	admin, _ := admin_model.GetAdminInfo(body.PhoneNumber)
	if admin != nil {
		appG.BadResponse("该电话已被注册")
		return
	}
	if appG.HasError(admin_service.CreateAdmin(body.PhoneNumber, body.Password)) {
		return
	}
	appG.SuccessResponse("管理员账号注册成功")
}

// @Tags 管理员
// @Summary 管理员登录
// @Description 管理员登录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/admin/login [post]
type AdminLoginBody struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

func AdminLogin(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body AdminLoginBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	token, err := admin_service.GenerateToken(body.PhoneNumber, body.Password)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(token)
}

// @Tags 管理员
// @Summary 获取管理员消息
// @Description 获取管理员消息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/admin/get [get]
type AdminInfoGetBody struct {
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
}

func AdminInfoGet(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body AdminInfoGetBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	admin, err := admin_model.GetAdminInfoPart(body.PhoneNumber)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(admin)
}

// @Tags 管理员
// @Summary 获取全部管理员消息
// @Description 获取全部管理员消息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/admin/getAll [get]
func AdminInfoGetAll(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	adminList, err := admin_model.GetAllAdminInfo()
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(map[string]interface{}{"adminList": adminList})
}

// @Tags 管理员
// @Summary 根据账号修改管理员密码
// @Description 根据账号修改管理员密码
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/admin/update [post]
type AdminPasswordUpdateBody struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
	NewPassWord string `json:"newpassword" validate:"required"`
}

func AdminPasswordUpdate(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body AdminPasswordUpdateBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := admin_model.GetAdminInfo(body.PhoneNumber)
	if err != nil {
		appG.BadResponse("未找到该管理员账号")
		return
	}
	if appG.HasError(admin_service.UpdateAdminPassword(body.PhoneNumber, body.Password, body.NewPassWord)) {
		return
	}
	appG.SuccessResponse("修改管理员密码成功")
}

// @Tags 管理员
// @Summary 根据账号修改管理员信息
// @Description 根据账号修改管理员信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/admin/updateInfo [post]
type AdminInfoBody struct {
	PhoneNumber  string `json:"phone_number" validate:"required"`
	Nickname     string `json:"nickname"`
	EmailAddress string `json:"email_address"`
}

func AdminInfoUpdate(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body AdminInfoBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := admin_model.GetAdminInfo(body.PhoneNumber)
	if err != nil {
		appG.BadResponse("未找到该管理员账号")
		return
	}
	if appG.HasError(admin_service.UpdateAdminInfo(body.PhoneNumber, body.Nickname, body.EmailAddress)) {
		return
	}
	appG.SuccessResponse("修改管理员信息成功")
}

// @Tags 管理员
// @Summary 根据账号注销管理员
// @Description 根据账号注销管理员
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/admin/delete [post]
type AdminDeleteBody struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

func AdminDelete(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body AdminDeleteBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	if appG.HasError(admin_service.DeleteAdmin(body.PhoneNumber, body.Password)) {
		return
	}
	appG.SuccessResponse("注销管理员成功")
}

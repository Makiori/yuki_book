package v1

import (
	"yuki_book/model/userType_model"
	"yuki_book/service/userType_service"
	"yuki_book/util/app"

	"github.com/gin-gonic/gin"
)

// @Tags 用户类型
// @Summary 新增用户类型
// @Description 新增用户类型
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/userType/new [post]
type UserTypeNewBody struct {
	TypeName    string `json:"type_name" validate:"required"`
	MaxBorNum   int    `json:"max_bor_Num" validate:"required"`
	MaxTime     int    `json:"max_time" validate:"required"`
	MaxBorCount int    `json:"max_bor_count" validate:"required"`
}

func UserTypeNew(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body UserTypeNewBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	userType, _ := userType_model.GetUserTypeInfoByUserType(body.TypeName)
	if userType != nil {
		appG.BadResponse("该用户类型已经存在，请勿重复添加")
		return
	}

	if appG.HasError(userType_service.CreateUserType(
		body.TypeName,
		body.MaxBorNum,
		body.MaxTime,
		body.MaxBorCount)) {
		return
	}
	appG.SuccessResponse("添加用户类型成功")
}

// @Tags 用户类型
// @Summary 查看用户类型
// @Description 查看用户类型
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/userType/get [GET]
type UserTypeGetBody struct {
	Id int `json:"id" validate:"required"`
}

func UserTypeGet(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body UserTypeGetBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	userType, err := userType_model.GetUserTypeInfoById(body.Id)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(userType)
}

// @Tags 用户类型
// @Summary 查看用户类型信息
// @Description 查看用户类型信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/userType/getAll [GET]
func UserTypeGetAll(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	userTypeList, err := userType_model.GetAllUserTypeInfo()
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(map[string]interface{}{"userTypeList": userTypeList})
}

// @Tags 用户类型
// @Summary 修改用户类型信息
// @Description 修改用户类型信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/userType/update [post]
type UserTypeUpdateInfoBody struct {
	Id          int    `json:"id" validate:"required"`
	TypeName    string `json:"type_name"`
	MaxBorNum   int    `json:"max_bor_Num"`
	MaxTime     int    `json:"max_time"`
	MaxBorCount int    `json:"max_bor_count"`
}

func UserTypeUpdateInfo(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body UserTypeUpdateInfoBody
	if !appG.ParseJSONRequest(&body) {
		return
	}

	_, err := userType_model.GetUserTypeInfoById(body.Id)
	if err != nil {
		appG.BadResponse("未找到该用户类型")
		return
	}
	if appG.HasError(userType_service.UpdateUserTypeInfo(body.Id, body.TypeName, body.MaxBorNum, body.MaxTime, body.MaxBorCount)) {
		return
	}
	appG.SuccessResponse("修改用户信息成功")
}

// @Tags 用户类型
// @Summary 删除用户类型
// @Description 删除用户类型
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/userType/delete [post]
type UserTypeDeleteBody struct {
	Id int `json:"id" validate:"required"`
}

func UserTypeDelete(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body UserTypeDeleteBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	if appG.HasError(userType_service.DeleteUserType(body.Id)) {
		return
	}
	appG.SuccessResponse("删除用户类型成功")
}

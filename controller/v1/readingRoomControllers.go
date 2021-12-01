package v1

import (
	"yuki_book/controller"
	"yuki_book/model/readingRoom_model"
	"yuki_book/service/readingRoom_service"
	"yuki_book/util/app"

	"github.com/gin-gonic/gin"
)

// @Tags 阅览室
// @Summary 新增阅览室记录
// @Description 新增阅览室记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/readingRoom/new [post]
type ReadingRoomNewBody struct {
	Id       string `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Position string `json:"position" validate:"required"`
}

func ReadingRoomNew(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body ReadingRoomNewBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	readingRoom, _ := readingRoom_model.GetReadingRoomInfo(body.Id)
	if readingRoom != nil {
		appG.BadResponse("已有该阅览室记录")
		return
	}
	if appG.HasError(readingRoom_service.CreateReadingRoom(
		body.Id,
		body.Name,
		body.Position)) {
		return
	}
	appG.SuccessResponse("添加阅览室记录成功")

}

// @Tags 阅览室
// @Summary 删除阅览室记录
// @Description 删除阅览室记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/readingRoom/new [post]
type ReadingRoomDeleteBody struct {
	Id string `json:"id" form:"id" validate:"required"`
}

func ReadingRoomDelete(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body ReadingRoomDeleteBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	if appG.HasError(readingRoom_service.DeleteReadingRoom(body.Id)) {
		return
	}
	appG.SuccessResponse("删除阅览室记录成功")
}

// @Tags 阅览室
// @Summary 根据阅览室id修改阅览室记录
// @Description 根据阅览室id修改阅览室记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/readingRoom/update [post]
type ReadingRoomUpdateBody struct {
	Id       string `json:"id" validate:"required"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

func ReadingRoomUpdate(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body ReadingRoomUpdateBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := readingRoom_model.GetReadingRoomInfo(body.Id)
	if err != nil {
		appG.BadResponse("未找到该阅览室记录")
		return
	}
	if appG.HasError(readingRoom_service.UpdateReadingRoomInfo(body.Id, body.Name, body.Position)) {
		return
	}
	appG.SuccessResponse("修改阅览室信息成功")
}

// @Tags 阅览室
// @Summary 根据阅览室id查找阅览室记录
// @Description 根据阅览室id查找阅览室记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/readingRoom/get [get]
type ReadingRoomGetByIdBody struct {
	Id string `json:"id" form:"id" validate:"required"`
}

func ReadingRoomGetById(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body ReadingRoomGetByIdBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	readingRoom, err := readingRoom_model.GetReadingRoomInfo(body.Id)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(readingRoom)
}

// @Tags 阅览室
// @Summary 分页模糊查询阅览室记录
// @Description 分页模糊查询阅览室记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/readingRoom/getLike [get]

type ReadingRoomGetLikeBody struct {
	FilterName string `json:"filterName" form:"filterName"`
	Page       uint   `json:"page" form:"page"`
	PageSize   uint   `json:"pageSize" form:"pageSize"`
}

func ReadingRoomGetLike(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body ReadingRoomGetLikeBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	readingRoom, err := readingRoom_model.GetLikeReadingRoomInfo(body.FilterName, body.Page, body.PageSize)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(readingRoom)
}

// @Tags 阅览室
// @Summary 分页查询全部阅览室记录
// @Description 分页查询全部阅览室记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/readingRoom/getAll [get]
func ReadingRoomGetAll(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body controller.PaginationQueryBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	readingRoomList, err := readingRoom_model.GetAllReadingRoomInfo(body.Page, body.PageSize)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(readingRoomList)
}

package v1

import (
	"yuki_book/controller"
	"yuki_book/model/bookShelf_model"
	"yuki_book/model/readingRoom_model"
	"yuki_book/service/bookShelf_service"
	"yuki_book/util/app"

	"github.com/gin-gonic/gin"
)

// @Tags 书架
// @Summary 新增书架
// @Description 新增书架
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookShelf/new [post]

type BookSelfNewBody struct {
	Id            string `json:"id" validate:"required"`
	ReadingRoomId string `json:"reading_room_id" validate:"required"`
	Classify      string `json:"classify" validate:"required"`
}

func BookShelfNew(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookSelfNewBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := readingRoom_model.GetReadingRoomInfo(body.ReadingRoomId)
	if err != nil {
		appG.BadResponse("无阅览室记录")
		return
	}

	bookSelf, _ := bookShelf_model.GetBookSelfInfo(body.Id)
	if bookSelf != nil {
		appG.BadResponse("已有该书架")
		return
	}

	if appG.HasError(bookShelf_service.CreateBookSelf(
		body.Id,
		body.ReadingRoomId,
		body.Classify)) {
		return
	}
	appG.SuccessResponse("添加书架成功")

}

// @Tags 书架
// @Summary 删除书架
// @Description 删除书架
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookShelf/delete [post]
type BookShelfDeleteBody struct {
	Id string `json:"id" form:"id" validate:"required"`
}

func BookShelfDelete(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookShelfDeleteBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	if appG.HasError(bookShelf_service.DeleteBookShelf(body.Id)) {
		return
	}
	appG.SuccessResponse("删除书架记录成功")

}

// @Tags 书架
// @Summary 根据书架id修改书架记录
// @Description 根据书架id修改书架记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookShelf/update [post]

type BookShelfUpdateBody struct {
	Id            string `json:"id" validate:"required"`
	ReadingRoomId string `json:"reading_room_id" validate:"required"`
	Classify      string `json:"classify"`
}

func BookShelfUpdate(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookShelfUpdateBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := readingRoom_model.GetReadingRoomInfo(body.ReadingRoomId)
	if err != nil {
		appG.BadResponse("未找到该阅览室记录")
		return
	}

	_, err2 := bookShelf_model.GetBookSelfInfo(body.Id)
	if err2 != nil {
		appG.BadResponse("未找到该书架记录")
		return
	}

	if appG.HasError(bookShelf_service.UpdateBookShelfInfo(body.Id, body.ReadingRoomId, body.Classify)) {
		return
	}
	appG.SuccessResponse("修改书架信息成功")
}

// @Tags 书架
// @Summary 根据书架id查询书架记录
// @Description 根据书架id查询书架记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookShelf/get [get]

type BookShelfGetByIdBody struct {
	Id string `json:"id" form:"id" validate:"required"`
}

func BookShelfGetById(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookShelfGetByIdBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	bookShelf, err := bookShelf_model.GetBookSelfInfo(body.Id)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(bookShelf)
}

// @Tags 书架
// @Summary 分页模糊查询书架记录
// @Description 分页模糊查询书架记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookShelf/getLike [get]
type BookShelfGetLikeBody struct {
	FilterName string `json:"filterName" form:"filterName"`
	Page       uint   `json:"page" form:"page"`
	PageSize   uint   `json:"pageSize" form:"pageSize"`
}

func BookShelfGetLike(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookShelfGetLikeBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	bookShelf, err := bookShelf_model.GetLikeBookShelfInfo(body.FilterName, body.Page, body.PageSize)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(bookShelf)
}


func BookShelfGetAll(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body controller.PaginationQueryBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	bookShelfList, err := bookShelf_model.GetAllbookShelfInfo(body.Page, body.PageSize)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(bookShelfList)
}

func BookShelfGetReadingRoomInfo(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookShelfGetByIdBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	info, err := bookShelf_model.GetReadingRoomInfo(body.Id)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(info)
}

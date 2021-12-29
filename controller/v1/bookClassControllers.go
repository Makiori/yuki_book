package v1

import (
	"yuki_book/controller"
	"yuki_book/model/bookClass_model"
	"yuki_book/service/bookClass_service"
	"yuki_book/util/app"

	"github.com/gin-gonic/gin"
)

// @Tags 书集
// @Summary 新增书集
// @Description 新增书集
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookClass/new [post]
type BookClassNewBody struct {
	Id               string `json:"id" validate:"required"`
	BookName         string `json:"book_name" validate:"required"`
	BookAuthor       string `json:"book_author" validate:"required"`
	BookKey          string `json:"book_key" validate:"required"`
	BookEdit         string `json:"book_edit" validate:"required"`
	BookIntroduction string `json:"Book_introduction" validate:"required"`
	PageNum          int    `json:"page_num" validate:"required"`
	PublishTime      string `json:"publish_time" validate:"required"`
}

func BookClassNew(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookClassNewBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	bookClass, _ := bookClass_model.GetBookClassInfo(body.Id)
	if bookClass != nil {
		appG.BadResponse("已有该书集，请到该书集下添加书本")
		return
	}
	if appG.HasError(bookClass_service.CreateBookClass(
		body.Id,
		body.BookName,
		body.BookAuthor,
		body.BookKey,
		body.BookEdit,
		body.BookIntroduction,
		body.PageNum,
		body.PublishTime)) {
		return
	}
	appG.SuccessResponse("添加书集成功")

}

// @Tags 书集
// @Summary 删除书集
// @Description 删除书集
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookClass/delete [post]
type BookClassDeleteBody struct {
	Id string `json:"id" form:"id" validate:"required"`
}

func BookClassDelete(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookClassDeleteBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	if appG.HasError(bookClass_service.DeleteBookClass(body.Id)) {
		return
	}
	appG.SuccessResponse("删除书集成功")
}


func BookClassGetAll(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body controller.PaginationQueryBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	bookClassList, err := bookClass_model.GetAllbookClassInfo(body.Page, body.PageSize)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(bookClassList)
}

// @Tags 书集
// @Summary 根据书集id查找书集信息
// @Description 根据书集id查找书集信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookClass/getById [get]
type BookClassGetByIdBody struct {
	Id string `json:"id" form:"id" validate:"required"`
}

func BookClassGetById(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookClassGetByIdBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	bookClass, err := bookClass_model.GetBookClassInfo(body.Id)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(bookClass)
}

// @Tags 书集
// @Summary 管理员根据书集id修改书集信息
// @Description 管理员根据书集id修改书集信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookClass/updateInfo [post]
type BookClassUpdateByIdBody struct {
	Id               string `json:"id" validate:"required"`
	BookName         string `json:"book_name"`
	BookAuthor       string `json:"book_author"`
	BookKey          string `json:"book_key"`
	BookEdit         string `json:"book_edit"`
	BookIntroduction string `json:"Book_introduction"`
	PageNum          int    `json:"page_num"`
	PublishTime      string `json:"publish_time"`
}

func BookClassUpdateById(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookClassUpdateByIdBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := bookClass_model.GetBookClassInfo(body.Id)
	if err != nil {
		appG.BadResponse("查无此书集")
		return
	}
	if appG.HasError(bookClass_service.UpdateBookClass(body.Id, body.BookName, body.BookAuthor, body.BookKey, body.BookEdit,
		body.BookIntroduction, body.PageNum, body.PublishTime)) {
		return
	}
	appG.SuccessResponse("修改书集信息成功")
}

// @Tags 书集
// @Summary 分页模糊查询书集信息
// @Description 分页模糊查询书集信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookClass/getLike [get]

type BookClassGetLikeBody struct {
	FilterName string `json:"filterName" form:"filterName"`
	Page       uint   `json:"page" form:"page"`
	PageSize   uint   `json:"pageSize" form:"pageSize"`
}

func BookClassGetLike(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookClassGetLikeBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	bookClass, err := bookClass_model.GetLikeBookClassInfo(body.FilterName, body.Page, body.PageSize)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(bookClass)
}

// // @Tags 书集
// // @Summary 查询书集中书本的存放位置
// // @Description 查询书集中书本的存放位置
// // @Produce  json
// // @Success 200 {object} app.Response
// // @Failure 500 {object} app.Response
// // @Router con/v1/bookClass/getPosition [get]
// type BookClassGetPositionBody struct {
// 	Id       string `json:"id" form:"id" validate:"required"`
// 	Page     uint   `json:"page" form:"page"`
// 	PageSize uint   `json:"pageSize" form:"pageSize"`
// }

// func BookClassGetPosition(c *gin.Context) {
// 	appG := app.Gin{Ctx: c}
// 	var body BookClassGetPositionBody
// 	if !appG.ParseQueryRequest(&body) {
// 		return
// 	}

// 	positionList, err := bookClass_model.GetBookClassPosition(body.Id, body.Page, body.PageSize)
// 	if appG.HasError(err) {
// 		return
// 	}

// 	appG.SuccessResponse(positionList)
// }

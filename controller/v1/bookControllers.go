package v1

import (
	"yuki_book/model/bookClass_model"
	"yuki_book/model/bookShelf_model"
	"yuki_book/model/book_model"
	"yuki_book/service/book_service"
	"yuki_book/util/app"

	"github.com/gin-gonic/gin"
)

// @Tags 书本
// @Summary 新增书本
// @Description 新增书本
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/book/new [post]
type BookNewBody struct {
	Id          string `json:"id" validate:"required"`
	BookClassID string `json:"book_class_id" validate:"required"`
	ShelfId     string `json:"shelf_id" validate:"required"`
	BookState   int    `json:"book_state" validate:"oneof=0 1"`
	BookDamage  int    `json:"book_damage" validate:"oneof=0 1 2 3"`
}

func BookNew(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookNewBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := bookClass_model.GetBookClassInfo(body.BookClassID)
	if err != nil {
		appG.BadResponse("请先添加书集，再向书集里添加书本")
		return
	}
	book, _ := book_model.GetBookInfo(body.Id)
	if book != nil {
		appG.BadResponse("该书本id已经存在记录")
		return
	}
	_, err2 := bookShelf_model.GetBookSelfInfo(body.ShelfId)
	if err2 != nil {
		appG.BadResponse("无书架")
		return
	}

	if appG.HasError(book_service.CreateBook(
		body.Id,
		body.BookClassID,
		body.ShelfId,
		body.BookState,
		body.BookDamage)) {
		return
	}
	appG.SuccessResponse("添加书本成功")
}

// @Tags 书本
// @Summary 删除书本
// @Description 删除书本
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/book/delete [post]
type BookDeleteBody struct {
	Id string `json:"id" form:"id" validate:"required"`
}

func BookDelete(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookDeleteBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	if appG.HasError(book_service.DeleteBook(body.Id)) {
		return
	}
	appG.SuccessResponse("删除书本成功")
}

// @Tags 书本
// @Summary 根据书本id查找书本信息
// @Description 根据书本id查找书本信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/book/getById [get]
type BookGetByIdBody struct {
	Id string `json:"id" validate:"required"`
}

func BookGetById(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookGetByIdBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	book, err := book_model.GetBookInfo(body.Id)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(book)
}

// @Tags 书本
// @Summary 根据书集id查找书本信息
// @Description 根据书集id查找书本信息
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/book/getById [get]
type BookGetByClassIdBody struct {
	BookClassID string `json:"book_class_id" form:"book_class_id" validate:"required"`
}

func BookGetByClassId(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookGetByClassIdBody
	if !appG.ParseQueryRequest(&body) {
		return
	}

	bookList, err := book_model.GetBookInfoByClassId(body.BookClassID)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(map[string]interface{}{"bookList": bookList})
}

// @Tags 书本
// @Summary 管理员修改书本状态
// @Description 管理员修改书本状态
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/book/updateBookStatu [post]
type BookUpdateStatuBody struct {
	Id        string `json:"id" validate:"required"`
	BookState int    `json:"book_state" validate:"oneof=0 1"`
}

func BookUpdateStatu(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookUpdateStatuBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := book_model.GetBookInfo(body.Id)
	if err != nil {
		appG.BadResponse("查无此书本")
		return
	}
	if appG.HasError(book_service.UpdateBookStatu(body.Id, body.BookState)) {
		return
	}
	appG.SuccessResponse("修改书本状态成功")
}

// @Tags 书本
// @Summary 管理员修改书本受损程度
// @Description 管理员修改书本受损程度
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/book/updateBookDamage [post]
type BookUpdateDamageBody struct {
	Id         string `json:"id" validate:"required"`
	BookDamage int    `json:"book_damage" validate:"oneof=0 1 2 3"`
}

func BookUpdateDamage(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookUpdateDamageBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	_, err := book_model.GetBookInfo(body.Id)
	if err != nil {
		appG.BadResponse("查无此书本")
		return
	}
	if appG.HasError(book_service.UpdateBookDamage(body.Id, body.BookDamage)) {
		return
	}
	appG.SuccessResponse("修改书本受损程度成功")
}

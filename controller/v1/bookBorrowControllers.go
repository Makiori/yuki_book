package v1

import (
	"yuki_book/model/bookBorrow_model"
	"yuki_book/model/bookClass_model"
	"yuki_book/model/book_model"
	"yuki_book/model/userType_model"
	"yuki_book/model/user_model"
	"yuki_book/service/bookBorrow_service"
	"yuki_book/util/app"

	"github.com/gin-gonic/gin"
)

// @Tags 借阅记录
// @Summary 新增借阅记录
// @Description 新增借阅记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookBorrow/new [post]
type BookBorrowNewBody struct {
	Username    string `json:"username" validate:"required"`
	BookClassId string `json:"book_class_id" validate:"required"`
	BookId      string `json:"book_id" validate:"required"`
}

func BookBorrowNew(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookBorrowNewBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	user, err := user_model.GetUserInfo(body.Username)
	if err != nil {
		appG.BadResponse("无用户")
		return
	}
	bookClass, err2 := bookClass_model.GetBookClassInfo(body.BookClassId)
	if err2 != nil {
		appG.BadResponse("无书集")
		return
	}
	book, err3 := book_model.GetBookInfo(body.BookId)
	if err3 != nil {
		appG.BadResponse("无书本")
		return
	}
	if book.BookClassID != bookClass.Id {
		appG.BadResponse("该书本不属于当前书集")
		return
	}
	if book.BookState == 1 {
		appG.BadResponse("该书本出借中，无法再次出借")
		return
	}
	userType, _ := userType_model.GetUserTypeInfoById(user.UserType)

	err4 := bookBorrow_service.CheckBookBorrowNum(body.Username, userType.MaxBorNum)
	if err4 != nil {
		appG.BadResponse("已超过可借书本的上限")
		return
	}

	// 新增借阅记录
	if appG.HasError(bookBorrow_service.CreateBookBorrow(body.Username, body.BookClassId, body.BookId)) {
		return
	}

	appG.SuccessResponse("借书成功")
}

// @Tags 借阅记录
// @Summary 根据账号查找借阅记录
// @Description 根据账号查找借阅记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookBorrow/new [post]
type BookBorrowGetBody struct {
	Username string `json:"username" form:"username" validate:"required"`
}

func BookBorrowGet(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookBorrowGetBody
	if !appG.ParseQueryRequest(&body) {
		return
	}
	bookBorrow, err := bookBorrow_model.GetBookBorrowInfo(body.Username)
	if appG.HasError(err) {
		return
	}
	appG.SuccessResponse(bookBorrow)
}

// @Tags 借阅记录
// @Summary 根据账号查找借阅记录，修改续借次数和还书日期
// @Description 根据账号查找借阅记录，修改续借次数和还书日期
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookBorrow/update [post]
type BookBorrowUpdateBody struct {
	Username    string `json:"username" validate:"required"`
	BookClassId string `json:"book_class_id" validate:"required"`
	BookId      string `json:"book_id" validate:"required"`
}

func BookBorrowUpdate(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookBorrowUpdateBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	if appG.HasError(bookBorrow_service.UpdateBookBorrow(body.Username, body.BookClassId, body.BookId)) {
		return
	}
	appG.SuccessResponse("续借成功")
}

// @Tags 借阅记录
// @Summary 根据账号查找借阅记录，修改还书记录实现还书
// @Description 根据账号查找借阅记录，修改还书记录实现还书
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router con/v1/bookBorrow/return [post]
type BookBorrowReturnUpdateBody struct {
	Username    string `json:"username" validate:"required"`
	BookClassId string `json:"book_class_id" validate:"required"`
	BookId      string `json:"book_id" validate:"required"`
}

func BookBorrowReturnUpdate(c *gin.Context) {
	appG := app.Gin{Ctx: c}
	var body BookBorrowReturnUpdateBody
	if !appG.ParseJSONRequest(&body) {
		return
	}
	if appG.HasError(bookBorrow_service.ReturnBookBorrow(body.Username, body.BookClassId, body.BookId)) {
		return
	}
	appG.SuccessResponse("还书成功")
}

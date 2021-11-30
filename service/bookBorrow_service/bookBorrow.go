package bookBorrow_service

import (
	"errors"
	"fmt"
	"time"
	"yuki_book/model/bookBorrow_model"
	"yuki_book/model/userType_model"
	"yuki_book/model/user_model"

	uuid "github.com/satori/go.uuid"
)

// 新增借阅记录
func CreateBookBorrow(username string, bookclassid string, bookid string) error {
	id := uuid.NewV4()

	// 生成借书时间
	nowTime := time.Now()

	// 查找用户类型信息
	user, err := user_model.GetUserInfo(username)
	if err != nil {
		return err
	}

	userType, err := userType_model.GetUserTypeInfoById(user.UserType)
	if err != nil {
		return err
	}

	newTime := nowTime.AddDate(0, 0, userType.MaxTime)

	bookBorrow := &bookBorrow_model.BookBorrow{
		Id:          id.String(),
		UserName:    username,
		BookClassId: bookclassid,
		BookId:      bookid,
		BorrowAt:    nowTime,
		BeReturnAt:  newTime,
		BorrowCount: 0,
	}
	return bookBorrow.Create()
}

// 查询该账号已借阅数量
func CheckBookBorrowNum(username string, maxbornum int) error {
	count, _ := bookBorrow_model.GetBookBorrowNum(username)
	if count >= maxbornum {
		return errors.New("超过上限")
	}
	return nil
}

// 续借功能，记录续借次数和更改还书时间
func UpdateBookBorrow(username string, bookClassId string, bookId string) error {

	bookBorrow, err := bookBorrow_model.GetBookBorrowInfo2(username, bookClassId, bookId)
	if err != nil {
		return err
	}

	// 查找用户类型信息
	user, err := user_model.GetUserInfo(username)
	if err != nil {
		return err
	}

	userType, err := userType_model.GetUserTypeInfoById(user.UserType)
	if err != nil {
		return err
	}

	a := bookBorrow.BeReturnAt
	b := time.Now()

	fmt.Println(bookBorrow.BorrowCount)
	fmt.Println(userType.MaxBorCount)

	if bookBorrow.BorrowCount < userType.MaxBorCount {
		if b.Before(a) {
			bookBorrow.BeReturnAt = bookBorrow.BeReturnAt.AddDate(0, 0, userType.MaxTime)
		} else {
			bookBorrow.BeReturnAt = b.AddDate(0, 0, userType.MaxTime)
		}

	} else {
		return errors.New("已到最大续借记录，无法再次续借，请还书之后再重新借阅")
	}

	bookBorrow.BorrowCount++

	return bookBorrow.Update()
}

// 还书功能，更改还书日期与还书状态
func ReturnBookBorrow(username string, bookClassId string, bookId string) error {

	bookBorrow, err := bookBorrow_model.GetBookBorrowInfo2(username, bookClassId, bookId)
	if err != nil {
		return err
	}

	bookBorrow.ReturnAt = time.Now()
	bookBorrow.ReturnStatu = 1

	//a := bookBorrow.BeReturnAt
	//b := time.Now()
	//fmt.Println(int(a.Sub(b).Hours() / 24))

	return bookBorrow.ReturnBook()
}

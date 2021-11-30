package bookBorrow_model

import (
	"errors"
	"time"
	"yuki_book/model/bookClass_model"
	"yuki_book/model/book_model"
	"yuki_book/model/database"
	"yuki_book/util/logging"
	"yuki_book/util/times"

	"github.com/jinzhu/gorm"
)

type ReturnStatu int

const (
	ReturnNo ReturnStatu = iota
	ReturnYes
)

func (r ReturnStatu) String() string {
	switch r {
	case ReturnNo:
		return "未还"
	case ReturnYes:
		return "已还"
	default:
		return "unknown"
	}
}

// 借阅记录表
type BookBorrow struct {
	Id          string         `db:"id"`
	UserName    string         `db:"user_name"`
	BookClassId string         `db:"book_class_id"`
	BookId      string         `db:"book_id"`
	BorrowAt    time.Time      `db:"borrow_at"`
	BorrowCount int            `db:"borrow_count"`
	BeReturnAt  time.Time      `db:"be_return_at"`
	ReturnAt    time.Time      `db:"return_at"`
	ReturnStatu ReturnStatu    `db:"return_statu"`
	CreatedAt   times.JsonTime `db:"created_at"`
	UpdatedAt   times.JsonTime `db:"updated_at"`
}

//通过用户账号获取借阅记录
func GetBookBorrowInfo(username string) (*[]BookBorrow, error) {
	var bookBorrow []BookBorrow
	err := database.DBCon.Where("user_name = ? and return_statu = 0", username).Find(&bookBorrow).Error
	if err != nil {
		return nil, errors.New("无借阅记录")
	}
	return &bookBorrow, nil
}

// 通过用户账号、书集id、书本id获取借阅记录
func GetBookBorrowInfo2(username string, bookclassid string, bookid string) (*BookBorrow, error) {
	var bookBorrow BookBorrow
	err := database.DBCon.Where("user_name = ? and book_class_id = ? and book_id = ? and return_statu = 0", username, bookclassid, bookid).Find(&bookBorrow).Error
	if err != nil {
		return nil, errors.New("无借阅记录")
	}
	return &bookBorrow, nil
}

//通过用户获取借阅记录数量
func GetBookBorrowNum(username string) (int, error) {
	var count int
	err := database.DBCon.Model(&BookBorrow{}).Where("user_name = ? and return_statu = 0", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 新增借阅记录
func (b *BookBorrow) Create() error {

	return database.DBCon.Transaction(func(tx *gorm.DB) error {

		// 1.往借阅记录中新增记录
		if err := tx.Create(&b).Error; err != nil {
			return err
		}

		// 2.更改图书状态
		if err := tx.Model(&book_model.Book{}).Where("id = ?", b.BookId).Update("book_state", book_model.BookState(1)).Error; err != nil {
			return err
		}

		// 3.修改在馆数量
		var bookClass bookClass_model.BookClass
		if err := tx.Model(&bookClass_model.BookClass{}).Where("id = ?", b.BookClassId).Find(&bookClass).Error; err != nil {
			return err
		}

		bookClass.BookIn = bookClass.BookIn - 1

		if err := tx.Model(&bookClass_model.BookClass{}).Where("id = ?", b.BookClassId).Update("book_in", bookClass.BookIn).Error; err != nil {
			return err
		}

		return nil
	})
}

// 续借修改借阅记录
func (b *BookBorrow) Update() error {
	sql := database.DBCon.Model(&BookBorrow{}).Where("user_name = ? and book_class_id = ? and book_id = ?", b.UserName, b.BookClassId, b.BookId).Updates(&b)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 还书修改借阅记录，同时修改书本状态和在馆数量
func (b *BookBorrow) ReturnBook() error {

	return database.DBCon.Transaction(func(tx *gorm.DB) error {

		// 1.修改借阅记录, 修改实际还书日期和还书状态
		if err := tx.Model(&BookBorrow{}).Where("user_name = ? and book_class_id = ? and book_id = ?", b.UserName, b.BookClassId, b.BookId).Updates(&b).Error; err != nil {
			return err
		}

		// 2.更改图书状态
		if err := tx.Model(&book_model.Book{}).Where("id = ?", b.BookId).Update("book_state", book_model.BookState(0)).Error; err != nil {
			return err
		}

		// 3.修改在馆数量
		var bookClass bookClass_model.BookClass
		if err := tx.Model(&bookClass_model.BookClass{}).Where("id = ?", b.BookClassId).Find(&bookClass).Error; err != nil {
			return err
		}

		bookClass.BookIn++
		if err := tx.Model(&bookClass_model.BookClass{}).Where("id = ?", b.BookClassId).Updates(bookClass).Error; err != nil {
			return err
		}

		return nil
	})
}

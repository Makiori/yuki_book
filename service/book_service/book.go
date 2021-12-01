package book_service

import (
	"yuki_book/model/book_model"
	"yuki_book/util/errors"

	"github.com/jinzhu/gorm"
)

// 新增书本
func CreateBook(id string, bookclassid string, shelfid string, bookstate int, bookdamage int) error {
	book := &book_model.Book{
		Id:          id,
		BookClassID: bookclassid,
		ShelfId:     shelfid,
		BookState:   book_model.BookState(bookstate),
		BookDamage:  book_model.BookDamage(bookdamage),
	}
	return book.Create()
}

// 删除书本
func DeleteBook(id string) error {
	b, err := book_model.GetBookInfo(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.BadError("查无此书")
		}
		return err
	}
	return b.DeleteBook()
}

// 修改书本状态
func UpdateBookStatu(id string, bookstate int) error {
	book := book_model.Book{
		Id:        id,
		BookState: book_model.BookState(bookstate),
	}
	return book.UpdateStatu()
}

// 修改书本受损程度
func UpdateBookDamage(id string, bookdamage int) error {
	book := book_model.Book{
		Id:         id,
		BookDamage: book_model.BookDamage(bookdamage),
	}
	return book.UpdateDamage()
}

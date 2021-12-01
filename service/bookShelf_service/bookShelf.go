package bookShelf_service

import (
	"yuki_book/model/bookShelf_model"
	"yuki_book/util/errors"

	"github.com/jinzhu/gorm"
)

// 新增书架记录
func CreateBookSelf(id string, readingroomid string, classify string) error {
	bookSelf := &bookShelf_model.BookShelf{
		Id:            id,
		ReadingRoomId: readingroomid,
		Classify:      classify,
	}
	return bookSelf.Create()
}

// 删除书架记录
func DeleteBookShelf(id string) error {
	bookShelf, err := bookShelf_model.GetBookSelfInfo(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.BadError("无当前阅览室记录")
		}
		return err
	}
	return bookShelf.Delete()
}

// 修改书架记录
func UpdateBookShelfInfo(id string, readingroomid string, classify string) error {
	bookShelf := bookShelf_model.BookShelf{
		ReadingRoomId: readingroomid,
		Classify:      classify,
	}
	return bookShelf.Update(id)
}

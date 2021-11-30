package bookClass_service

import (
	"yuki_book/model/bookClass_model"
	"yuki_book/util/errors"

	"github.com/jinzhu/gorm"
)

//新增书集
func CreateBookClass(id string, bookname string, bookauthor string, bookkey string, bookedit string, bookintroduction string,
	pagenum int, publishTime string) error {
	bookClass := &bookClass_model.BookClass{
		Id:               id,
		BookName:         bookname,
		BookAuthor:       bookauthor,
		BookKey:          bookkey,
		BookEdit:         bookedit,
		BookIntroduction: bookintroduction,
		PageNum:          pagenum,
		PublishTime:      publishTime,
		BookNum:          0,
		BookIn:           0,
	}
	return bookClass.Create()
}

// 删除书集
func DeleteBookClass(id string) error {
	_, err := bookClass_model.GetBookClassInfo(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.BadError("查无此书集")
		}
		return err
	}
	b := &bookClass_model.BookClass{
		Id: id,
	}
	return b.DeleteBookClass()
}

// 修改书集
func UpdateBookClass(id string, bookname string, bookauthor string, bookkey string, bookedit string,
	bookintroduction string, pagenum int, publishtime string) error {
	bookClass := bookClass_model.BookClass{
		Id:               id,
		BookName:         bookname,
		BookAuthor:       bookauthor,
		BookKey:          bookkey,
		BookEdit:         bookedit,
		BookIntroduction: bookintroduction,
		PageNum:          pagenum,
		PublishTime:      publishtime,
	}
	return bookClass.Update()
}

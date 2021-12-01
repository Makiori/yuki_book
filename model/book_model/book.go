package book_model

import (
	"yuki_book/model/bookClass_model"
	"yuki_book/model/database"
	"yuki_book/util/logging"
	"yuki_book/util/times"

	"github.com/jinzhu/gorm"
)

type BookState int

const (
	BookIn BookState = iota
	BookOut
)

func (b BookState) String() string {
	switch b {
	case BookIn:
		return "在馆"
	case BookOut:
		return "外借"
	default:
		return "unknown"
	}
}

type BookDamage int

const (
	state1 BookDamage = iota
	state2
	state3
	state4
)

func (b BookDamage) String() string {
	switch b {
	case state1:
		return "新书"
	case state2:
		return "轻微损伤"
	case state3:
		return "中等损伤"
	case state4:
		return "重度损伤"
	default:
		return "unknown"
	}
}

type Book struct {
	Id          string         `db:"id"`
	BookClassID string         `db:"book_class_id"`
	ShelfId     string         `db:"shelf_id"`
	BookState   BookState      `db:"book_state"`
	BookDamage  BookDamage     `db:"book_damage"`
	CreatedAt   times.JsonTime `db:"created_at"`
	UpdatedAt   times.JsonTime `db:"updated_at"`
}

// 通过id获取书本信息
func GetBookInfo(id string) (*Book, error) {
	var book Book
	err := database.DBCon.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// 新增书本
func (b *Book) Create() error {

	return database.DBCon.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&b).Error; err != nil {
			return err
		}

		var bookClass bookClass_model.BookClass
		if err := tx.Model(&bookClass_model.BookClass{}).Where("id = ?", b.BookClassID).Find(&bookClass).Error; err != nil {
			return err
		}

		bookClass.BookNum = bookClass.BookNum + 1
		bookClass.BookIn = bookClass.BookIn + 1

		if err := tx.Model(&bookClass_model.BookClass{}).Where("id = ?", b.BookClassID).Updates(&bookClass).Error; err != nil {
			return err
		}

		return nil
	})
}

// 删除书本
func (b *Book) DeleteBook() error {

	return database.DBCon.Transaction(func(tx *gorm.DB) error {

		if err := tx.Delete(&b).Error; err != nil {
			return err
		}

		var bookClass bookClass_model.BookClass
		if err := tx.Model(&bookClass_model.BookClass{}).Where("id = ?", b.BookClassID).Find(&bookClass).Error; err != nil {
			return err
		}

		bookClass.BookNum = bookClass.BookNum - 1
		bookClass.BookIn = bookClass.BookIn - 1

		if err := tx.Model(&bookClass_model.BookClass{}).Where("id = ?", b.BookClassID).Update("book_num", bookClass.BookNum).Error; err != nil {
			return err
		}

		if err := tx.Model(&bookClass_model.BookClass{}).Where("id = ?", b.BookClassID).Update("book_in", bookClass.BookIn).Error; err != nil {
			return err
		}

		return nil
	})
}

// 修改书本状态
func (b *Book) UpdateStatu() error {
	sql := database.DBCon.Model(&Book{}).Where("id = ?", b.Id).Update("book_state", b.BookState)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 修改书本受损程度
func (b *Book) UpdateDamage() error {
	sql := database.DBCon.Model(&Book{}).Where("id = ?", b.Id).Update("book_damage", b.BookDamage)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

func GetBookInfoByClassId(bookclassid string) (interface{}, error) {
	var bookList []Book
	err := database.DBCon.Select("*").Where("book_class_id = ?", bookclassid).Find(&bookList).Error
	if err != nil {
		return nil, err
	}
	return bookList, nil
}

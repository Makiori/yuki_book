package bookClass_model

import (
	"yuki_book/model"
	"yuki_book/model/database"
	"yuki_book/util/logging"
	"yuki_book/util/times"
)

// 书集表
type BookClass struct {
	Id               string         `db:"id"`
	BookName         string         `db:"book_name"`
	BookAuthor       string         `db:"book_author"`
	BookKey          string         `db:"book_key"`
	BookEdit         string         `db:"book_edit"`
	BookIntroduction string         `db:"book_introduction"`
	PageNum          int            `db:"page_num"`
	PublishTime      string         `db:"publish_time"`
	BookNum          int            `db:"book_num"`
	BookIn           int            `db:"book_in"`
	CreatedAt        times.JsonTime `db:"created_at"`
	UpdatedAt        times.JsonTime `db:"updated_at"`
}

// 通过id获取书集信息
func GetBookClassInfo(id string) (*BookClass, error) {
	var bookClass BookClass
	err := database.DBCon.Where("id = ?", id).First(&bookClass).Error
	if err != nil {
		return nil, err
	}
	return &bookClass, nil
}

// 新增书集
func (b *BookClass) Create() error {
	return database.DBCon.Create(&b).Error
}

// 删除书集
func (b *BookClass) DeleteBookClass() error {
	return database.DBCon.Where("id = ?", b.Id).Delete(&BookClass{}).Error
}

// 修改书集
func (b *BookClass) Update() error {
	sql := database.DBCon.Model(b).Where("id = ?", b.Id).Updates(&b)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 修改在馆数量
func UpdateBookIn(id string) error {
	var bookClass BookClass
	err := database.DBCon.Model(&BookClass{}).Where("id = ?", id).Find(&bookClass).Error
	if err != nil {
		return err
	}
	bookClass.BookIn--
	sql := database.DBCon.Model(&BookClass{}).Where("id = ?", id).Updates(bookClass)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 分页查询全部书集信息
func GetAllbookClassInfo(page uint, pagesize uint) (data *model.PaginationQ, err error) {
	q := model.PaginationQ{
		Page:     page,
		PageSize: pagesize,
		Data:     &[]BookClass{},
	}
	return q.SearchAll(database.DBCon.Model(&BookClass{}))
}

// 分页模糊查询书集信息
func GetLikeBookClassInfo(filtername string, page uint, pagesize uint) (data *model.PaginationQ, err error) {
	q := model.PaginationQ{
		PageSize: pagesize,
		Page:     page,
		Data:     &[]BookClass{},
	}
	args := "%" + filtername + "%"
	data, err = q.SearchAll(
		database.DBCon.Model(&BookClass{}).Where("book_name like ? or book_author like ? or book_key like ? or book_edit like ? or book_introduction like ? ", args, args, args, args, args),
	)
	if err != nil {
		return nil, err
	}
	return
}

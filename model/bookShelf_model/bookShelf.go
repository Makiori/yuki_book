package bookShelf_model

import (
	"yuki_book/model"
	"yuki_book/model/database"
	"yuki_book/model/readingRoom_model"
	"yuki_book/util/logging"
	"yuki_book/util/times"
)

// 书架表
type BookShelf struct {
	Id            string         `db:"id"`
	ReadingRoomId string         `db:"reading_room_id"`
	Classify      string         `db:"classify"`
	CreatedAt     times.JsonTime `db:"created_at"`
	UpdatedAt     times.JsonTime `db:"updated_at"`
}

// 通过id获取书架信息
func GetBookSelfInfo(id string) (*BookShelf, error) {
	var bookSelf BookShelf
	err := database.DBCon.Where("id = ?", id).First(&bookSelf).Error
	if err != nil {
		return nil, err
	}
	return &bookSelf, nil
}

// 新增书架记录
func (b *BookShelf) Create() error {
	return database.DBCon.Create(&b).Error
}

// 删除书架记录
func (b *BookShelf) Delete() error {
	return database.DBCon.Delete(&b).Error
}

// 更改书架记录
func (b *BookShelf) Update(id string) error {
	sql := database.DBCon.Model(b).Where("id = ?", id).Updates(&b)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 分页模糊查询书架记录
func GetLikeBookShelfInfo(filtername string, page uint, pagesize uint) (data *model.PaginationQ, err error) {
	q := model.PaginationQ{
		PageSize: pagesize,
		Page:     page,
		Data:     &[]BookShelf{},
	}
	args := "%" + filtername + "%"
	data, err = q.SearchAll(
		database.DBCon.Model(&BookShelf{}).Where("classify like ?", args),
	)
	if err != nil {
		return nil, err
	}
	return
}

// 分页查询全部阅览室信息
func GetAllbookShelfInfo(page uint, pagesize uint) (data *model.PaginationQ, err error) {
	q := model.PaginationQ{
		Page:     page,
		PageSize: pagesize,
		Data:     &[]BookShelf{},
	}
	return q.SearchAll(database.DBCon.Model(&BookShelf{}))
}

type ReadingRoomInfo struct {
	ReadingRoomId string `db:"reading_room_id"`
	Name          string `db:"name"`
	Position      string `db:"postion"`
	ShelfId       string `db:"id"`
	Classify      string `db:"classify"`
}

func GetReadingRoomInfo(id string) (interface{}, error) {
	var bookShelf BookShelf
	err := database.DBCon.Where("id = ?", id).First(&bookShelf).Error
	if err != nil {
		return nil, err
	}
	var readingRoom readingRoom_model.ReadingRoom
	err2 := database.DBCon.Where("id = ?", bookShelf.ReadingRoomId).First(&readingRoom).Error
	if err2 != nil {
		return nil, err2
	}
	readingRoomInfo := &ReadingRoomInfo{
		ReadingRoomId: bookShelf.ReadingRoomId,
		Name:          readingRoom.Name,
		Position:      readingRoom.Position,
		ShelfId:       bookShelf.Id,
		Classify:      bookShelf.Classify,
	}
	return readingRoomInfo, nil
}

package readingRoom_model

import (
	"yuki_book/model"
	"yuki_book/model/database"
	"yuki_book/util/logging"
	"yuki_book/util/times"
)

// 阅览室表
type ReadingRoom struct {
	Id        string         `db:"id"`
	Name      string         `db:"name"`
	Position  string         `db:"postion"`
	CreatedAt times.JsonTime `db:"created_at"`
	UpdatedAt times.JsonTime `db:"updated_at"`
}

// 根据id获取阅览室信息
func GetReadingRoomInfo(id string) (*ReadingRoom, error) {
	var readingRoom ReadingRoom
	err := database.DBCon.Where("id = ?", id).First(&readingRoom).Error
	if err != nil {
		return nil, err
	}
	return &readingRoom, nil
}

// 新增阅览室记录
func (r *ReadingRoom) Create() error {
	return database.DBCon.Create(&r).Error
}

// 删除阅览室记录
func (r *ReadingRoom) Delete() error {
	return database.DBCon.Delete(&r).Error
}

// 更改阅览室记录
func (r *ReadingRoom) Update(id string) error {
	sql := database.DBCon.Model(r).Where("id = ?", id).Updates(&r)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 分页模糊查询阅览室记录
func GetLikeReadingRoomInfo(filtername string, page uint, pagesize uint) (data *model.PaginationQ, err error) {
	q := model.PaginationQ{
		PageSize: pagesize,
		Page:     page,
		Data:     &[]ReadingRoom{},
	}
	args := "%" + filtername + "%"
	data, err = q.SearchAll(
		database.DBCon.Model(&ReadingRoom{}).Where("name like ? or position like ?", args, args),
	)
	if err != nil {
		return nil, err
	}
	return
}

// 分页查询全部阅览室信息
func GetAllReadingRoomInfo(page uint, pagesize uint) (data *model.PaginationQ, err error) {
	q := model.PaginationQ{
		Page:     page,
		PageSize: pagesize,
		Data:     &[]ReadingRoom{},
	}
	return q.SearchAll(database.DBCon.Model(&ReadingRoom{}))
}

package readingRoom_service

import (
	"yuki_book/model/readingRoom_model"
	"yuki_book/util/errors"

	"github.com/jinzhu/gorm"
)

// 新增阅览室记录
func CreateReadingRoom(id string, name string, position string) error {
	readingRoom := &readingRoom_model.ReadingRoom{
		Id:       id,
		Name:     name,
		Position: position,
	}
	return readingRoom.Create()
}

// 删除阅览室记录
func DeleteReadingRoom(id string) error {
	r, err := readingRoom_model.GetReadingRoomInfo(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.BadError("无当前阅览室记录")
		}
		return err
	}
	return r.Delete()
}

// 修改阅览室记录
func UpdateReadingRoomInfo(id string, name string, position string) error {
	readingRoom := readingRoom_model.ReadingRoom{
		Name:     name,
		Position: position,
	}
	return readingRoom.Update(id)
}

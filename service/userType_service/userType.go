package userType_service

import (
	"yuki_book/model/userType_model"
	"yuki_book/util/errors"

	"github.com/jinzhu/gorm"
)

func CreateUserType(typename string, maxbornum int, maxtime int, maxborcount int) error {
	userType := &userType_model.UserType{
		TypeName:    typename,
		MaxBorNum:   maxbornum,
		MaxTime:     maxtime,
		MaxBorCount: maxborcount,
	}
	return userType.Create()
}

func UpdateUserTypeInfo(id int, typename string, maxbornum int, maxtime int, maxborcount int) error {
	userType := &userType_model.UserType{
		Id:          id,
		TypeName:    typename,
		MaxBorNum:   maxbornum,
		MaxTime:     maxtime,
		MaxBorCount: maxborcount,
	}
	return userType.Update()
}

func DeleteUserType(id int) error {
	u, err := userType_model.GetUserTypeInfoById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.BadError("用户类型不存在")
		}
		return err
	}
	return u.DeleteUserType()
}

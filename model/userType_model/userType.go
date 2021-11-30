package userType_model

import (
	"yuki_book/model/database"
	"yuki_book/util/logging"
	"yuki_book/util/times"
)

// 用户类型
type UserType struct {
	Id          int            `gorm:"id"`
	TypeName    string         `gorm:"type_name"`
	MaxBorNum   int            `gorm:"max_bor_num"`
	MaxTime     int            `gorm:"max_time"`
	MaxBorCount int            `gorm:"max_bor_count"`
	CreatedAt   times.JsonTime `gorm:"created_at"`
	UpdatedAt   times.JsonTime `gorm:"updated_at"`
}

// 通过id获取相关信息
func GetUserTypeInfoById(usertype int) (*UserType, error) {
	var userType UserType
	err := database.DBCon.Where("id = ?", usertype).First(&userType).Error
	if err != nil {
		return nil, err
	}
	return &userType, nil
}

// 通过用户类型名字获取相关信息
func GetUserTypeInfoByUserType(typename string) (*UserType, error) {
	var userType UserType
	err := database.DBCon.Where("type_name = ?", typename).First(&userType).Error
	if err != nil {
		return nil, err
	}
	return &userType, nil
}

// 新增用户类型
func (u *UserType) Create() error {
	return database.DBCon.Create(&u).Error
}

// 获取全部用户类型的信息
func GetAllUserTypeInfo() (interface{}, error) {
	var userTypeList []UserType
	err := database.DBCon.Select("*").Find(&userTypeList).Error
	if err != nil {
		return nil, err
	}
	return userTypeList, nil
}

func (u *UserType) Update() error {
	sql := database.DBCon.Model(u).Where("id = ?", u.Id).Updates(&u)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

func (u *UserType) DeleteUserType() error {
	return database.DBCon.Where("id = ?", u.Id).Delete(&UserType{}).Error
}

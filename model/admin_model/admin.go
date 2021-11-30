package admin_model

import (
	"yuki_book/model/database"
	"yuki_book/util/logging"
	"yuki_book/util/times"
)

// 管理员表
type Admin struct {
	Id           string         `db:"id"`
	PhoneNumber  string         `db:phone_number"`
	Password     string         `db:"password"`
	Salt         string         `db:"salt"`
	Nickname     string         `db:"nickname"`
	EmailAddress string         `db:"email_address"`
	CreatedAt    times.JsonTime `db:"created_at"`
	UpdatedAt    times.JsonTime `db:"updated_at"`
}

//通过管理员账号获取管理员信息
func GetAdminInfo(phonenumber string) (*Admin, error) {
	var admin Admin
	err := database.DBCon.Where("phone_number = ?", phonenumber).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// 通过管理员账号获取部分管理员信息
func GetAdminInfoPart(phonenumber string) (*Admin, error) {
	var admin Admin
	err := database.DBCon.Select("id, phone_number, nickname, email_address").Where("phone_number = ?", phonenumber).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

//通过管理员账号获取全部管理员信息
func GetAllAdminInfo() (interface{}, error) {
	var adminList []Admin
	err := database.DBCon.Select("id, phone_number, nickname, email_address").Find(&adminList).Error
	if err != nil {
		return nil, err
	}
	return adminList, nil
}

// 注册管理员账号
func (a *Admin) Create() error {
	return database.DBCon.Create(&a).Error
}

// 通过管理员账号更改管理员密码
func (a *Admin) UpdateAdminPasswrod() error {
	sql := database.DBCon.Model(a).Where("phone_number = ?", a.PhoneNumber).Updates(&a)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 通过管理员账号更改管理员信息
func (a *Admin) UpdateAdminInfo(phonenumber string) error {
	sql := database.DBCon.Model(a).Where("phone_number = ?", phonenumber).Updates(&a)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 通过管理员账号注销管理员账号
func (a *Admin) DeleteAdmin() error {
	return database.DBCon.Where("phone_number = ?", a.PhoneNumber).Delete(&Admin{}).Error
}

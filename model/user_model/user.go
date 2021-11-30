package user_model

import (
	"yuki_book/model"
	"yuki_book/model/database"
	"yuki_book/util/logging"
	"yuki_book/util/times"
)

// 用户表
type User struct {
	Id           string         `db:"id"`
	Username     string         `db:"username"`
	Password     string         `db:"password"`
	Salt         string         `db:"salt"`
	Nickname     string         `db:"Nickname"`
	PhoneNumber  string         `db:"phone_number"`
	Class        string         `db:"class"`
	EmailAddress string         `db:"Email_address"`
	UserType     int            `db:"user_type"`
	CreatedAt    times.JsonTime `db:"created_at"`
	UpdatedAt    times.JsonTime `db:"updated_at"`
}

//通过用户账号获取用户信息
func GetUserInfo(username string) (*User, error) {
	var user User
	err := database.DBCon.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 通过用户账号获取部分用户信息
func GetUserInfoPart(username string) (*User, error) {
	var user User
	err := database.DBCon.Select("id, username, nickname, phone_number, class, email_address, user_type").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 获取全部账号信息
func GetAllUserInfo(page uint, pagesize uint) (data *model.PaginationQ, err error) {
	q := model.PaginationQ{
		Page:     page,
		PageSize: pagesize,
		Data:     &[]User{},
	}
	return q.SearchAll(database.DBCon.Model(&User{}))
}

// 注册用户账号
func (u *User) Create() error {
	return database.DBCon.Create(&u).Error
}

// 通过用户账号更改用户密码
func (u *User) UpdateUserPassword() error {
	sql := database.DBCon.Model(u).Where("username = ?", u.Username).Updates(&u)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 通过用户账号更改用户信息
func (u *User) UpdateUserInfo(userName string) error {
	sql := database.DBCon.Model(u).Where("username = ?", userName).Updates(&u)
	rowsAffected := sql.RowsAffected
	logging.Infof("更新影响的记录数%d", rowsAffected)
	logging.Infoln(sql.Error)
	return sql.Error
}

// 通过用户账号注销用户账号
func (u *User) DeleteUser() error {
	return database.DBCon.Where("username = ?", u.Username).Delete(&User{}).Error
}

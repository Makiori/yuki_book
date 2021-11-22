package user_service

import (
	"yuki_book/model/user_model"
	"yuki_book/util/errors"
	"yuki_book/util/sign"

	"github.com/jinzhu/gorm"
)

// 生成用户token
func GenerateToken(username, password string) (interface{}, error) {
	user, err := user_model.GetUserInfo(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.BadError("用户账号不存在")
		}
		return nil, err
	}
	if user.Password != sign.EncodeMD5(password+user.Salt) {
		return nil, errors.BadError("密码错误")
	}
	token, err := sign.GenerateToken(string(user.Username), username, sign.UserClaimsType)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"token": token, "Username": user.Username}, nil
}

// 用户注册
func CreateUser(username string, password string) error {
	salt := "ABCDEF"
	userType := "学生"
	user := &user_model.User{
		Username: username,
		Password: sign.EncodeMD5(password + salt),
		Salt:     salt,
		UserType: userType,
	}
	return user.Create()
}

// 通过用户账号更改用户密码
func UpdateUserPassword(username string, password string, newPassword string) error {
	user, err := user_model.GetUserInfo(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.BadError("用户账号不存在")
		}
		return err
	}
	// 校验密码
	if user.Password != sign.EncodeMD5(password+user.Salt) {
		return errors.BadError("密码错误")
	}
	u := &user_model.User{
		Username: username,
		Password: sign.EncodeMD5(newPassword + user.Salt),
	}
	return u.UpdateUserPassword()
}

// 通过用户账号修改用户信息
func UpdateUserInfo(userName string, name *string, userPhonenumber *string, userAddress *string, userClass *string, userEmail *string) error {
	user := user_model.User{
		Name:            name,
		UserPhonenumber: userPhonenumber,
		UserClass:       userClass,
		UserAddress:     userAddress,
		UserEmail:       userEmail,
	}
	if err := user.UpdateUserInfo(userName); err != nil {
		return errors.BadError("修改用户信息失败")
	}
	return nil
}

// 管理员通过用户账号修改用户类型
func UpdateUserType(username string, userType string) error {
	user := user_model.User{
		UserType: userType,
	}
	if err := user.UpdateUserInfo(username); err != nil {
		return errors.BadError("修改用户类型失败")
	}
	return nil
}

// 注销用户账号
func DeleteUser(userName string, password string) error {
	user, err := user_model.GetUserInfo(userName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.BadError("用户账号不存在")
		}
		return err
	}
	if user.Password != sign.EncodeMD5(password+user.Salt) {
		return errors.BadError("密码错误")
	}
	u := &user_model.User{
		Username: userName,
		Password: sign.EncodeMD5(password + user.Salt),
	}
	return u.DeleteUser()
}

package admin_service

import (
	"yuki_book/model/admin_model"
	"yuki_book/util/errors"
	"yuki_book/util/sign"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// 生成管理员token
func GenerateToken(phonenumber, password string) (interface{}, error) {
	admin, err := admin_model.GetAdminInfo(phonenumber)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.BadError("管理员账号不存在")
		}
		return nil, err
	}
	if admin.Password != sign.EncodeMD5(password+admin.Salt) {
		return nil, errors.BadError("密码错误")
	}
	token, err := sign.GenerateToken(string(admin.PhoneNumber), phonenumber, sign.AdminClaimsType)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"token": token, "PhoneNumber": admin.PhoneNumber}, nil
}

// 管理员注册
func CreateAdmin(phoneNumber string, password string) (err error) {
	id := uuid.NewV4()
	salt := "ABCDEF"
	admin := &admin_model.Admin{
		Id:          id.String(),
		PhoneNumber: phoneNumber,
		Password:    sign.EncodeMD5(password + salt),
		Salt:        salt,
	}
	return admin.Create()
}

// 通过管理员账号更改管理员密码
func UpdateAdminPassword(phoneNumber string, password string, newPassword string) error {
	admin, err := admin_model.GetAdminInfo(phoneNumber)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.BadError("管理员账号不存在")
		}
		return err
	}
	// 校验密码
	if admin.Password != sign.EncodeMD5(password+admin.Salt) {
		return errors.BadError("密码错误")
	}
	a := &admin_model.Admin{
		PhoneNumber: phoneNumber,
		Password:    sign.EncodeMD5(newPassword + admin.Salt),
	}
	return a.UpdateAdminPasswrod()
}

// 通过管理员账号修改管理员信息
func UpdateAdminInfo(userName string, nickname string, emailaddress string) error {
	admin := admin_model.Admin{
		Nickname:     nickname,
		EmailAddress: emailaddress,
	}
	if err := admin.UpdateAdminInfo(userName); err != nil {
		return errors.BadError("修改管理员信息失败")
	}
	return nil
}

// 注销管理员账号
func DeleteAdmin(phonenumber string, password string) error {
	admin, err := admin_model.GetAdminInfo(phonenumber)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.BadError("管理员账号不存在")
		}
		return err
	}
	if admin.Password != sign.EncodeMD5(password+admin.Salt) {
		return errors.BadError("密码错误")
	}
	a := &admin_model.Admin{
		PhoneNumber: phonenumber,
		Password:    sign.EncodeMD5(password + admin.Salt),
	}
	return a.DeleteAdmin()
}

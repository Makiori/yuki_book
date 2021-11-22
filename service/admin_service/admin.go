package admin_service

import (
	"yuki_book/model/admin_model"
	"yuki_book/util/errors"
	"yuki_book/util/sign"

	"github.com/jinzhu/gorm"
)

// 生成管理员token
func GenerateToken(username, password string) (interface{}, error) {
	admin, err := admin_model.GetAdminInfo(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.BadError("管理员账号不存在")
		}
		return nil, err
	}
	if admin.Password != sign.EncodeMD5(password+admin.Salt) {
		return nil, errors.BadError("密码错误")
	}
	token, err := sign.GenerateToken(string(admin.Username), username, sign.AdminClaimsType)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"token": token, "Username": admin.Username}, nil
}

// 管理员注册
func CreateAdmin(username string, password string) (err error) {
	salt := "ABCDEF"
	admin := &admin_model.Admin{
		Username: username,
		Password: sign.EncodeMD5(password + salt),
		Salt:     salt,
	}
	return admin.Create()
}

// 通过管理员账号更改管理员密码
func UpdateAdminPassword(username string, password string, newPassword string) error {
	admin, err := admin_model.GetAdminInfo(username)
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
		Username: username,
		Password: sign.EncodeMD5(newPassword + admin.Salt),
	}
	return a.UpdateAdminPasswrod()
}

// 通过管理员账号修改管理员信息
func UpdateAdminInfo(userName string, adminName *string, adminPhoneNumber *string, adminAddress *string) error {
	admin := admin_model.Admin{
		AdminName:        adminName,
		AdminPhonenumber: adminPhoneNumber,
		AdminAddress:     adminAddress,
	}
	if err := admin.UpdateAdminInfo(userName); err != nil {
		return errors.BadError("修改管理员信息失败")
	}
	return nil
}

// 注销管理员账号
func DeleteAdmin(userName string, password string) error {
	admin, err := admin_model.GetAdminInfo(userName)
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
		Username: userName,
		Password: sign.EncodeMD5(password + admin.Salt),
	}
	return a.DeleteAdmin()
}

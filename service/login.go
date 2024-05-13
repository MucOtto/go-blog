package service

import (
	"errors"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/utils"
)

func GetLoginInfo(username string, password string) (*models.LoginResponse, error) {
	password = utils.Md5Crypt(password, "mszlu")
	user := dao.GetUser(username, password)
	if user == nil {
		return nil, errors.New("账号或密码不正确")
	}

	uid := user.Uid
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token生成失败")
	}

	lr := &models.LoginResponse{
		Token: token,
		UserInfo: models.UserInfo{
			Uid:      user.Uid,
			UserName: user.UserName,
			Avatar:   user.Avatar,
		},
	}
	return lr, nil
}

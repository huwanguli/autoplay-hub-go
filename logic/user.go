package logic

import (
	"autoplay-hub/dao/mysql"
	"autoplay-hub/models"
	"autoplay-hub/pkg/jwt"
	"autoplay-hub/pkg/snowflake"
)

func Register(p *models.ParamRegister) (err error) {
	// 判断是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	userID := snowflake.GenID()
	// 生成用户实例
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
		UserID:   userID,
		IsAdmin:  p.IsAdmin,
	}
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(user.UserID, user.Username, user.IsAdmin)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return
}

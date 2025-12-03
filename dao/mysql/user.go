package mysql

import (
	"autoplay-hub/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "autoplay"

// 对密码进行md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from users where username=?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

func InsertUser(user *models.User) (err error) {
	sqlStr := `insert into users(user_id,username,password,is_admin) values(?,?,?,?)`
	user.Password = encryptPassword(user.Password)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.IsAdmin)
	return
}

func Login(p *models.User) (err error) {
	sqlStr := `select user_id, username, password, is_admin from users where username=?`
	oPassword := p.Password
	err = db.Get(p, sqlStr, p.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrorUserNotExist
	} else if err != nil {
		return err
	}
	if encryptPassword(oPassword) != p.Password {
		return ErrorInvalidPassword
	}
	return
}

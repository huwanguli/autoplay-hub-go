package logic

import (
	"autoplay-hub/dao/mysql"
	"autoplay-hub/models"

	"go.uber.org/zap"
)

func CreateScript(p *models.ParamScript) error {
	if err := mysql.CheckScriptExist(p.ScriptName); err != nil {
		return err
	}
	return mysql.InsertScript(p)
}

// GetAllScriptList 按照page，size，以及排序规则获取脚本列表
func GetAllScriptList(page int64, size int64, userID int64) (scripts []*models.ApiScript, err error) {
	scripts = make([]*models.ApiScript, 0)
	var script []*models.Script
	if userID == 0 {
		script, err = mysql.GetAllScripts(page, size)
		if err != nil {
			return nil, err
		}

	} else {
		script, err = mysql.GetAllScriptsByUserID(page, size, userID)
		if err != nil {
			return nil, err
		}
	}
	for _, data := range script {
		username, err := mysql.CheckUserNameByID(data.OwnerID)
		if err != nil {
			zap.L().Error("mysql.CheckUserNameByID failed", zap.Error(err), zap.Int64("user_id", data.OwnerID))
			continue
		}
		scripts = append(scripts, &models.ApiScript{
			Id:          data.ID,
			Name:        data.Name,
			Description: data.Description,
			CreateAt:    data.CreateAt,
			UpdateAt:    data.UpdateAt,
			Owner:       username,
		})
	}
	return

}

// GetScriptDetail 获取脚本详情
func GetScriptDetail(id int64, userID int64) (script *models.ApiScriptDetail, err error) {
	script = new(models.ApiScriptDetail)
	data := new(models.Script)
	data, err = mysql.GetScriptDetailByScriptID(id)
	if err != nil {
		zap.L().Error("mysql.GetScriptDetailByScriptID failed", zap.Int64("script_id", id), zap.Error(err))
		return nil, err
	}
	var uID int64
	uID, err = mysql.CheckUserIDByScriptID(data.ID)
	if err != nil {
		zap.L().Error("mysql.CheckUerIDByScriptID failed", zap.Int64("user_id", userID), zap.Int64("uid", uID), zap.Error(err))
		return nil, err
	}
	if uID != userID {
		err = ErrorInvalidUserID
		return nil, err
	}
	script = &models.ApiScriptDetail{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Content:     data.Content,
		UpdateAt:    data.UpdateAt,
		CreateAt:    data.CreateAt,
	}
	return
}

// UpdateScript 编辑脚本功能的实现
func UpdateScript(p *models.ParamUpdateScript, id int64, ownerID int64) (err error) {
	// 脚本所有人的校验
	var Owner int64
	if Owner, err = mysql.GetUserIdByScriptID(id); err != nil {
		return err
	}
	if Owner != ownerID {
		err = ErrorInvalidUserID
		return err
	}
	Sp := new(models.Script)
	Sp, err = mysql.GetScriptDetailByScriptID(id)
	if err != nil {
		return
	}
	PMap := make(map[string]string)
	if p.ScriptName != nil {
		PMap["Name"] = *p.ScriptName
	} else {
		PMap["Name"] = Sp.Name
	}
	if p.Description != nil {
		PMap["Description"] = *p.Description
	} else {
		PMap["Description"] = Sp.Description
	}
	if p.Content != nil {
		PMap["Content"] = *p.Content
	} else {
		PMap["Content"] = Sp.Content
	}

	// 修改数据库，更新脚本内容
	return mysql.UpdateScript(id, PMap)
}

// DeleteScript 删除脚本
func DeleteScript(id int64, userID int64) (err error) {
	OwnerID, err := mysql.GetUserIdByScriptID(id)
	if err != nil {
		return err
	}
	if OwnerID != userID {
		err = ErrorInvalidUserID
		return err
	}
	return mysql.DeleteScript(id)
}

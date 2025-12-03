package logic

import (
	"autoplay-hub/dao/mysql"
	"autoplay-hub/models"
)

func CreateScript(p *models.ParamScript) error {
	if err := mysql.CheckScriptExist(p.ScriptName); err != nil {
		return err
	}
	return mysql.InsertScript(p)
}

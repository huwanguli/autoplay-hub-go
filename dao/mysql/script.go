package mysql

import "autoplay-hub/models"

func CheckScriptExist(scriptName string) (err error) {
	sqlStr := `select count(id) from scripts where name=?`
	var count int
	if err = db.Get(&count, sqlStr, scriptName); err != nil {
		return err
	}
	if count > 0 {
		return ErrorScriptExist
	}
	return nil
}

func InsertScript(p *models.ParamScript) (err error) {
	sqlStr := `insert into scripts(name,description,owner_id,content) values(?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ScriptName, p.Description, p.OwnerID, p.Content)
	return
}

package mysql

import (
	"autoplay-hub/models"
	"database/sql"
	"errors"
)

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

// GetAllScriptsByUserID 根据ID获取脚本列表
func GetAllScriptsByUserID(page int64, size int64, userID int64) (list []*models.Script, err error) {
	sqlStr := `select id,name,description,owner_id,content,created_at,updated_at
			   from scripts 
			   where owner_id=? 
			   order by created_at 
    		   desc limit ?, ?`
	list = make([]*models.Script, 0)
	err = db.Select(&list, sqlStr, userID, (page-1)*size, size)
	return
}

// GetAllScripts 管理员获取全部脚本列表
func GetAllScripts(page, size int64) (list []*models.Script, err error) {
	sqlStr := `select id,name,description,owner_id,content,created_at,updated_at
			   from scripts 
			   order by created_at 
    		   desc limit ?, ?`
	list = make([]*models.Script, 0)
	err = db.Select(&list, sqlStr, (page-1)*size, size)
	return
}

func GetScriptDetailByScriptID(scriptID int64) (script *models.Script, err error) {
	sqlStr := `select id,name,description,owner_id,content,created_at,updated_at
               from scripts
               where id=?`
	script = new(models.Script)
	err = db.Get(script, sqlStr, scriptID)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrorScriptNotExist
	}
	return
}

// CheckUserIDByScriptID  查询脚本对应的userID
func CheckUserIDByScriptID(scriptID int64) (userID int64, err error) {
	sqlStr := `select owner_id from scripts where id=?`
	err = db.Get(&userID, sqlStr, scriptID)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrorScriptNotExist
	}
	return
}

func UpdateScript(id int64, PMap map[string]string) (err error) {
	sqlStr := `update scripts set name=?,description=?,content=? where id=?`
	result, err := db.Exec(sqlStr, PMap["Name"], PMap["Description"], PMap["Content"], id)
	if err != nil {
		return err
	}
	count, _ := result.RowsAffected()
	if count != 1 {
		return ErrorUpdateFailed
	}
	return nil
}

// GetUserIdByScriptID 根据脚本ID获取所有人ID
func GetUserIdByScriptID(id int64) (userID int64, err error) {
	sqlStr := `select owner_id from scripts where id=?`
	err = db.Get(&userID, sqlStr, id)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrorScriptNotExist
		return 0, err
	}
	return
}

// DeleteScript 删除脚本
func DeleteScript(id int64) (err error) {
	sqlStr := `delete from scripts where id=?`
	result, err := db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	count, _ := result.RowsAffected()
	if count != 1 {
		return ErrorScriptNotExist
	}
	return

}

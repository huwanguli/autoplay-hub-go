package mysql

import (
	"autoplay-hub/models"
	"database/sql"
	"errors"
)

func InsertTask(p *models.Task) (err error) {
	sqlStr := `insert into tasks (script_id, task_name, status, log_content, user_id) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ScriptID, p.TaskName, p.Status, p.Log, p.UserID)
	return
}

// GetAllTasks 获取全部任务
func GetAllTasks(page, size int64) (list []*models.Task, err error) {
	sqlStr := `select id,script_id,task_name,status,log_content,user_id
			   from tasks
			   order by created_at 
    		   desc limit ?, ?`
	list = make([]*models.Task, 0)
	err = db.Select(&list, sqlStr, (page-1)*size, size)
	return
}

// GetAllTasksByUserID  根据ID获取任务列表
func GetAllTasksByUserID(page int64, size int64, userID int64) (list []*models.Task, err error) {
	sqlStr := `select id,script_id,task_name,status,log_content,user_id
			   from tasks
			   where user_id=? 
			   order by created_at 
    		   desc limit ?, ?`
	list = make([]*models.Task, 0)
	err = db.Select(&list, sqlStr, userID, (page-1)*size, size)
	return
}

// GetTaskDetailByID 获取任务详情
func GetTaskDetailByID(id int64) (data *models.ApiTaskDetail, err error) {
	sqlStr := `select id,script_id,task_name,status,log_content, user_id , created_at from tasks where id=?`
	data = new(models.ApiTaskDetail)
	err = db.Get(data, sqlStr, id)
	return
}

// CheckUserIDByTaskID 根据任务ID查用户ID
func CheckUserIDByTaskID(id int64) (userID int64, err error) {
	sqlStr := `select user_id from tasks where id=?`
	err = db.Get(&userID, sqlStr, id)
	return
}

func GetMaxTaskID() (id int64, err error) {
	sqlStr := `select max(id) from tasks`
	err = db.Get(&id, sqlStr)
	return
}

func UpdateTask(p *models.ParamStopTask) (err error) {
	sqlStr := `update tasks set status=?, executed_at=? where id=?`
	_, err = db.Exec(sqlStr, p.Status, p.ExecutedAt, p.TaskID)
	return
}

// CheckTaskExistsByID 判断任务是否存在
func CheckTaskExistsByID(id int64) (err error) {
	sqlStr := `select count(id) from tasks where id=?`
	var count int
	err = db.Get(&count, sqlStr, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrorTaskNotExist
		}
	}
	if count == 0 {
		err = ErrorTaskNotExist
	}
	return
}

// UpdateTaskName 更新任务名称
func UpdateTaskName(p *models.ParamUpdateTask) (err error) {
	sqlStr := `update tasks set task_name=? where id=?`
	_, err = db.Exec(sqlStr, p.Name, p.TaskID)
	return
}

// DeleteTask 删除任务
func DeleteTask(id int64) (err error) {
	sqlStr := `delete from tasks where id=?`
	_, err = db.Exec(sqlStr, id)
	return
}

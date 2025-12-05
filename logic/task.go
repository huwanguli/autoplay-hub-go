package logic

import (
	"autoplay-hub/dao/mysql"
	"autoplay-hub/dao/redis"
	"autoplay-hub/models"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func ScriptRunFirst(p *models.ParamScriptRun, userID int64) (err error) {
	// 1.根据参数填充Task所需参数
	task := &models.Task{
		UserID:   userID,
		ScriptID: p.ScriptID,
		TaskName: fmt.Sprintf("Run script %d in %s",
			p.ScriptID,
			time.Now().Format("2006-01-02 15:04")),
		Status: 1,
	}
	// 2.将初始任务信息进行入库
	if err := mysql.InsertTask(task); err != nil {
		return err
	}
	return
}

func ScriptRun(p *models.ParamScriptRun) (err error) {
	// 1.根据脚本ID获取脚本内容
	var data string
	data, err = mysql.GetScriptByScriptID(p.ScriptID)
	if err != nil {
		return err
	}
	// 2.将消息RPush到redis中供python消费者读取
	taskID, err := mysql.GetMaxTaskID()
	if err != nil {
		return err
	}
	MsgP := &models.Message{
		DeviceURL: p.DeviceUrl,
		TaskID:    taskID,
		Content:   data,
	}
	// 此处变为字节码，在python端中需要解码(utf-8)
	Msg, err := json.Marshal(MsgP)
	fmt.Println(Msg)
	if err != nil {
		zap.L().Error("JSONMarshal Error", zap.Any("Msg", Msg), zap.Error(err))
		return err
	}
	if err := redis.PushMsg(Msg); err != nil {
		zap.L().Error("Redis Push Error", zap.Any("Msg", Msg), zap.Error(err))
		return err
	}
	return
}

// GetAllTaskList 按照page，size，以及排序规则获取脚本列表
func GetAllTaskList(page int64, size int64, userID int64) (tasks []*models.ApiTask, err error) {
	tasks = make([]*models.ApiTask, 0)
	var task []*models.Task
	if userID == 0 {
		task, err = mysql.GetAllTasks(page, size)
		if err != nil {
			return nil, err
		}

	} else {
		task, err = mysql.GetAllTasksByUserID(page, size, userID)
		if err != nil {
			return nil, err
		}
	}
	for _, data := range task {
		scriptName, err := mysql.CheckScriptNameByID(data.ScriptID)
		if err != nil {
			zap.L().Error("mysql.CheckUserNameByID failed", zap.Error(err), zap.Int64("user_id", data.UserID))
			continue
		}
		tasks = append(tasks, &models.ApiTask{
			TaskID:     data.ID,
			TaskName:   data.TaskName,
			ScriptName: scriptName,
			Status:     data.Status,
		})
	}
	return

}

// GetTaskDetail 获取脚本详情
func GetTaskDetail(id int64, userID int64) (task *models.ApiTaskDetail, err error) {
	task = new(models.ApiTaskDetail)
	task, err = mysql.GetTaskDetailByID(id)
	if err != nil {
		zap.L().Error("mysql.GetTaskDetailByScriptID failed", zap.Int64("task_id", id), zap.Error(err))
		return nil, err
	}
	var uID int64
	uID, err = mysql.CheckUserIDByTaskID(task.ID)
	if err != nil {
		zap.L().Error("mysql.CheckUerIDByTaskID failed", zap.Int64("user_id", userID), zap.Int64("uid", uID), zap.Error(err))
		return nil, err
	}
	if uID != userID {
		err = ErrorInvalidUserID
		return nil, err
	}
	return
}

// TaskStop  停止任务
func TaskStop(p *models.ParamStopTask) (err error) {
	// 向消费者发送停止信号
	if err = redis.TaskStop(p); err != nil {
		zap.L().Error("redis.TaskStop failed", zap.Error(err))
		return err
	}
	// 更新mysql中的任务信息
	if err = mysql.UpdateTask(p); err != nil {
		zap.L().Error("mysql.UpdateTask failed", zap.Error(err))
		return err
	}
	return
}

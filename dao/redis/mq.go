package redis

import (
	"autoplay-hub/models"

	"go.uber.org/zap"
)

// 暂时使用list来做消息队列
// 因为对于消息的真实性、安全性、以及消息的数量要求都不高
// 即使消息丢失也不会有大影响，用户只要重新运行即可

func PushMsg(msg []byte) (err error) {
	_, err = client.RPush("tasks", msg).Result()
	if err != nil {
		zap.L().Error("Redis Push Error", zap.Any("Msg", msg), zap.Error(err))
		return err
	}
	return
}

func TaskStop(p *models.ParamStopTask) (err error) {
	_, err = client.RPush("tasks:stop", p.TaskID).Result()
	if err != nil {
		zap.L().Error("Redis Push Error", zap.Any("Msg", p), zap.Error(err))
		return err
	}
	return
}

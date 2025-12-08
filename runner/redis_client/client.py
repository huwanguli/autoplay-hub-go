import redis
import json
import logging
logger = logging.getLogger("runner")
class RedisClient:
    def __init__(self, host, port, db):
        self.client = redis.Redis(
            host = host,
            port = port,
            db = db
        )

    def Ping(self):
        self.client.ping()

    # 获取任务
    def GetTask(self):
        data_bytes = self.client.blpop(["tasks"], 0)
        try:
            data_str = data_bytes[1].decode('utf-8')
            # 解析JSON返回字典类型
            data = json.loads(data_str)
            return data
        except Exception as e:
            logger.error("Error decoding or parsing:", e)

    def GetStopMsg(self):
        stop_msg = self.client.blpop(["tasks:stop"],0)
        try:
            data_str = stop_msg[1].decode('utf-8')
            # 解析JSON返回字典类型
            data = json.loads(data_str)
            return data
        except Exception as e:
            logger.error("Error decoding or parsing:", e)



def init_redis(cfg):
    rc = RedisClient(cfg.host, cfg.port, cfg.db)
    return rc


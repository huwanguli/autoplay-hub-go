import asyncio
import yaml
from setting.setting import AppConfig
from log.logger import init_logger
from redis_client.client import init_redis
from models.task import Task
from runner.runner import RunTask

# 1.初始化配置类
# 读取项目根目录的yaml文件
with open("../config.yaml", 'r') as f:
    cfgYaml = yaml.safe_load(f)
# 配置类
Appcfg = AppConfig(cfgYaml)

# 2.日志初始化
logger = init_logger(Appcfg.logger)
# 3.redis初始化
client = init_redis(Appcfg.redis)
try:
    client.Ping()
    logger.info("runner redis ping success")
except Exception as e:
    logger.error(e)




stop_event = asyncio.Event()

async def CheckStopMsg(taskid):
    while True :
        print("正在监听停止信号")
        msg = client.GetStopMsg()
        if msg == taskid:
            stop_event.set()

async def Main(ptask,taskid):
    print("已开启task1")
    task1 = asyncio.create_task(RunTask(ptask))
    print("已开启task2")
    task2 = asyncio.create_task(CheckStopMsg(taskid))
    await asyncio.gather(task1, task2)

# 启动任务
while True:
    print("已开启主循环")
    task = Task(client.GetTask())
    task_id = task.task_id
    asyncio.run(Main(task,task_id))

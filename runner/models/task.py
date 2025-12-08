import json

class Task:
    def __init__(self, task_info):
        self.device_url = task_info["device_id"]
        self.script_id = task_info["script_id"]
        self.task_id = task_info["task_id"]
        self.content = json.loads(task_info["content"])
        self.log = ""

    #任务开始时给LOG写入开始运行
    def StartLog(self):
        self.log += str(self.script_id) + " 开始运行。" + "\n"

    #任务失败时给LOG写入任务失败
    def FailedLog(self):
        self.log += str(self.script_id) + " 任务失败。" + "\n"

    #任务成功时写入任务完成
    def SuccessLog(self):
        self.log += str(self.script_id) + " 任务完成。" + "\n"

    def __str__(self):
        return self.device_url + " " + str(self.script_id) + " " + str(self.task_id) + " " +json.dumps(self.content) +" " + self.log

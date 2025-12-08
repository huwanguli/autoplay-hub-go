import asyncio
import logging
from airtest.core.api import *

logger = logging.getLogger("runner")

# 由于bug的存在，目前暂时只支持基于图像识别的点击和输入操作
# 由于bug的存在暂时对runner的模块不进行更新，在MUMU进行BUG修复之后在进行更新
async def RunTask(task):
    print("开始执行任务")
    task.StartLog()
    try:
        # airtest有bug 暂时直接指定
        connect_device("Android:///127.0.0.1:16384")
        await asyncio.sleep(100)
    except Exception as e:
        logger.error("auto_setup failed:", e)
        return
    for step in task.content["steps"]:
        print("loop start")
        if step["action"] == "touch":
            touchPic(step["params"])
        elif step["action"] == "text":
            text(step["params"])
        elif step["action"] == "sleep":
            sleep(int(step["params"]))
        # 每个节点结束后固定休眠1秒等待操作完成
        await asyncio.sleep(1)
    task.SuccessLog()

# 现在airtest有bug 截图固定为竖屏，现在为该bug修正如下
def touchPic(picture_dir):
    template = Template(picture_dir)
    pos = exists(template)
    if pos:
        touch((1980 - pos[1], pos[0]))
        return True
    else:
        return False
FRONTDIRSTR = "../"

class RedisConfig:
    def __init__(self, cfg):
        self.host = cfg["host"]
        self.port = cfg["port"]
        self.password = cfg["password"]
        self.db = cfg["db"]
        self.pool_size = cfg["pool_size"]

class LoggerConfig:
    def __init__(self, cfg):
        self.level = cfg["level"]
        self.filename = FRONTDIRSTR + cfg["filename"] #写入项目根目录的日志中
        self.max_size = cfg["max_size"]
        self.min_age = cfg["max_age"]
        self.max_backup = cfg["max_backup"]

class AppConfig:
   def __init__(self, cfg):
       self.name = cfg["name"]
       self.mode = cfg["mode"]
       self.version = cfg["version"]
       self.start_time = cfg["start_time"]
       self.redis = RedisConfig(cfg["redis"])
       self.logger = LoggerConfig(cfg["log"])


if __name__ == "__main__":
    pass

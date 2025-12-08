import logging


def set_logger(cfg):
    logging.basicConfig(
        level=logging.DEBUG,
        format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
        datefmt="%Y-%m-%d %H:%M:%S",
        filename = cfg.filename
    )


def init_logger(cfg):
    set_logger(cfg)
    logger = logging.getLogger("runner")
    logger.info("runner log init success")
    return logger





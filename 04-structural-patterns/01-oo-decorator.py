from datetime import datetime

# Base Logger Interface
class Logger:
    def log(self, message: str):
        raise NotImplementedError()

# Concrete Component
class ConsoleLogger(Logger):
    def log(self, message: str):
        print(message)

# Base Decorator
class LoggerDecorator(Logger):
    def __init__(self, logger: Logger):
        self._logger = logger

    def log(self, message: str):
        self._logger.log(message)

# Concrete Decorators
class TimestampLogger(LoggerDecorator):
    def log(self, message: str):
        timestamped = f"{datetime.now()} - {message}"
        super().log(timestamped)

class LevelLogger(LoggerDecorator):
    def __init__(self, logger: Logger, level: str):
        super().__init__(logger)
        self._level = level

    def log(self, message: str):
        leveled = f"[{self._level.upper()}] {message}"
        super().log(leveled)

# Usage
logger = ConsoleLogger()
logger = TimestampLogger(logger)
logger = LevelLogger(logger, "debug")

logger.log("Something happened")

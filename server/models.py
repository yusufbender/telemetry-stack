from pydantic import BaseModel
from datetime import datetime

class Metrics(BaseModel):
    timestamp: datetime
    cpu_percent: float
    memory_used: int
    memory_total: int
    disk_used: int
    disk_total: int

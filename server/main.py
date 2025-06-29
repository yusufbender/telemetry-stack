from fastapi import FastAPI
from models import Metrics
from fastapi.responses import JSONResponse
from typing import List
import os
from influxdb_client import InfluxDBClient, Point
from influxdb_client.client.write_api import SYNCHRONOUS

app = FastAPI()
metrics_storage = []

# Influx bağlantısı güvenli hale alındı
try:
    influx_client = InfluxDBClient(
        url=os.environ["INFLUXDB_URL"],
        token=os.environ["INFLUXDB_TOKEN"],
        org=os.environ["INFLUXDB_ORG"]
    )
    write_api = influx_client.write_api(write_options=SYNCHRONOUS)
    bucket = os.environ["INFLUXDB_BUCKET"]
except Exception as e:
    print("❌ Influx bağlantı hatası:", e)
    influx_client = None
    write_api = None
    bucket = None

@app.post("/metrics", response_class=JSONResponse)
def receive_metrics(data: Metrics):
    metrics_storage.append(data)
    print("✅ Yeni metrik alındı:", data)

    if write_api:
        try:
            point = (
                Point("system_metrics")
                .tag("source", "agent")
                .field("cpu_percent", data.cpu_percent)
                .field("memory_used", data.memory_used)
                .field("memory_total", data.memory_total)
                .field("disk_used", data.disk_used)
                .field("disk_total", data.disk_total)
                .time(data.timestamp)
            )
            write_api.write(bucket=bucket, record=point)
        except Exception as e:
            print("❌ Influx'a yazılamadı:", e)

    return {"status": "ok"}

@app.get("/metrics", response_model=List[Metrics])
def get_metrics():
    return metrics_storage[-10:]

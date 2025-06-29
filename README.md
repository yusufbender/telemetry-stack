# 📊 Telemetry Stack

Sunuculardan sistem metriklerini (CPU, RAM, Disk) toplayıp, FastAPI üzerinden alıp InfluxDB'ye kaydeden ve Grafana ile görselleştiren tam entegre bir Docker projesi.

## 🔧 Bileşenler

- 🐍 FastAPI → REST API sunucusu
- 🐹 Go Agent → Sistem verisi toplayıcı
- 🧠 InfluxDB → Zaman serisi veritabanı
- 📈 Grafana → Dashboard & görselleştirme
- 🐳 Docker Compose → Her şey konteynerde!

## 🚀 Kurulum

```bash
# Ortam dosyasını oluştur
cp .env.example .env

# Projeyi ayağa kaldır
docker compose up -d --build
```
## 📡 Agent Logları
docker logs -f telemetry-stack-agent-1


## 📍 Grafana Arayüzü
http://localhost:3000

Default kullanıcı: admin / admin
```
telemetry-stack/
├── agent/               # Go ile yazılmış sistem agent
├── server/              # FastAPI backend
├── docker-compose.yml   # Tüm stack'i tanımlar
└── .env.example         # Ortam değişkenleri (örnek)
```

## 🛡️ Güvenlik
Gerçek token .env içinde tutulur, .gitignore sayesinde GitHub'a gönderilmez.

## ✨ Katkıda Bulun
Pull request gönder, issue aç ya da yıldız bırak ⭐

![image](https://github.com/user-attachments/assets/ee97baba-9a2b-4a45-89dc-bf252568325a)


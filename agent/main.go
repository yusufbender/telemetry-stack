package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// Metrics struct
type Metrics struct {
	Timestamp   time.Time `json:"timestamp"`
	CPUPercent  float64   `json:"cpu_percent"`
	MemoryUsed  uint64    `json:"memory_used"`
	MemoryTotal uint64    `json:"memory_total"`
	DiskUsed    uint64    `json:"disk_used"`
	DiskTotal   uint64    `json:"disk_total"`
}

func collectMetrics() (*Metrics, error) {
	cpuPercents, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		return nil, err
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	// DEBUG: Metrikleri yazdır
	fmt.Printf("CPU: %.2f%%, RAM: %d/%d, Disk: %d/%d\n",
		cpuPercents[0], vm.Used, vm.Total, diskStat.Used, diskStat.Total)

	return &Metrics{
		Timestamp:   time.Now(),
		CPUPercent:  cpuPercents[0],
		MemoryUsed:  vm.Used,
		MemoryTotal: vm.Total,
		DiskUsed:    diskStat.Used,
		DiskTotal:   diskStat.Total,
	}, nil
}

func sendToServer(metrics *Metrics) error {
	jsonData, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	apiURL := os.Getenv("API_URL") // Environment variable’dan al
	if apiURL == "" {
		apiURL = "http://api:8000/metrics" // fallback
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Gönderildi:", resp.Status)
	return nil
}

func main() {
	for {
		metrics, err := collectMetrics()
		if err != nil {
			fmt.Println("Metrik toplama hatası:", err)
			continue
		}

		err = sendToServer(metrics)
		if err != nil {
			fmt.Println("Gönderme hatası:", err)
		}

		time.Sleep(10 * time.Second)
	}
}

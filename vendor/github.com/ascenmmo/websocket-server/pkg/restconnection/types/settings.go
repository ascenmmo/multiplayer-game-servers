package types

import (
	"fmt"
	"github.com/ascenmmo/websocket-server/env"
	"runtime"
)

type Settings struct {
	ServerType          string `json:"serverType"`
	ServerPort          string `json:"serverPort"`
	ConnectionPort      string `json:"connectionPort"`
	ServerAddress       string `json:"serverAddress"`
	MaxConnections      int    `json:"maxConnections"`
	MaxRequestPerSecond int    `json:"maxRequestPerSecond"`
}

func NewSettings() (settings Settings) {
	settings.ServerType = "websocket"
	settings.ServerPort = env.TCPPort
	settings.ConnectionPort = env.WebsocketPort
	settings.ServerAddress = env.ServerAddress
	settings.MaxConnections = CountConnectionsMAX()
	settings.MaxRequestPerSecond = env.MaxRequestPerSecond
	return settings
}

func CountConnectionsMAX() int {
	numCPUs := runtime.NumCPU()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	connections := calculateConnections(numCPUs, memStats.Sys)
	fmt.Printf("Количество CPU: %d\n", numCPUs)
	fmt.Printf("Объем оперативной памяти: %d MB\n", memStats.Sys/(1024*1024))
	fmt.Printf("Рекомендуемое количество соединений по UDP: %d\n", connections)

	return connections
}

func calculateConnections(cpuCount int, totalRAM uint64) int {
	connectionsPerCPU := 1000
	//connectionsPerGB := 5000

	//totalMemoryGB := totalRAM / (1024 * 1024 * 1024)

	connectionsByCPU := cpuCount * connectionsPerCPU
	//connectionsByRAM := int(totalMemoryGB) * connectionsPerGB

	return connectionsByCPU
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

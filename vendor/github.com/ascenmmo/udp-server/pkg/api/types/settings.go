package types

import (
	"github.com/ascenmmo/udp-server/env"
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
	settings.ServerType = "upd"
	settings.ServerPort = env.TCPPort
	settings.ConnectionPort = env.UDPPort
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

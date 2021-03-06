package interfaces

import (
	"gomine/tasks"
	"gomine/resources"
)

type IServer interface {
	IsRunning() bool
	Start()
	Shutdown()
	GetTickRate() int
	SetTickRate(int)
	ResetTickRate()
	GetScheduler() *tasks.Scheduler
	GetServerPath() string
	GetLogger() ILogger
	GetConfiguration() *resources.GoMineConfig
	GetCommandHolder() ICommandHolder
	GetLoadedLevels() map[int]ILevel
	IsLevelLoaded(string) bool
	IsLevelGenerated(string) bool
	LoadLevel(string) bool
	HasPermission(string) bool
	SendMessage(message string)
	GetName() string
	GetAddress() string
	GetPort() uint16
	GetMaximumPlayers() uint
	GetMotd() string
	Tick(int)
}

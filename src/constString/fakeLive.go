package constString

const (
	FakeLiveNormal = 0

	FakeLiveStart     = 1
	FakeLiveRestart   = 2
	FakeLiveStop      = 3
	FakeLiveKill      = 4
	FakeLiveSendStart = 5
	FakeLiveSendStop  = 6

	ApiLiveStart   = 7
	ApiLiveStop    = 8
	ApiLiveRestart = 9
)

const (
	FakeLiveCreateCheckTask = "createCheckTask"
	FakeLiveDeleteTask      = "deleteTask"
	FakeLiveStartTask       = "startTask"
	FakeLiveStopTask        = "stopTask"
	FakeLiveRestartTask     = "restartTask"
	FakeLiveSendStartTask   = "sendStartTask"
	FakeLiveSendStopTask    = "sendStopTask"
)

const LuaCreateTask = `
local taskKey = ARGV[1]
local result = redis.call('SETNX', taskKey, 1)  -- 自动实现原子性检查+设置
return result == 1 and 0 or 1
`

const LuaDeleteTask = `
local taskKey = ARGV[1]
local result = redis.call('DEL', taskKey)  -- 删除任务
return result == 1 and 0 or 1
`

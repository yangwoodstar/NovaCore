package tools

import "fmt"

func GetTCTaskIDKey(prefix, appID, roomID, taskID string) string {
	return fmt.Sprintf("%s_task:{%s}:%s:%s:tc_task_id", prefix, appID, roomID, taskID)
}

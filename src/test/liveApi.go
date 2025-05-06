package test

import (
	"encoding/json"
	"fmt"
	"github.com/yangwoodstar/NovaCore/src/constString"
	"github.com/yangwoodstar/NovaCore/src/core/liveByteInstance"
)

func CreateLiveApiTest() {
	createPullToPushTask := &liveByteInstance.PullToPushTaskConfig{
		Title:            "pushTest",
		StartTime:        1745919794,
		EndTime:          1745920814,
		CallbackUrl:      "",
		Type:             constString.RtmpVideoStream,
		CycleMode:        0,
		PushUrl:          "",
		PullUrl:          "",
		PreDownload:      1,
		ContinueStrategy: 1,
		StartOffset:      1800,
		EndOffset:        0,
	}

	liveConfig := &liveByteInstance.LiveConfig{
		AK: "",
		SK: "",
	}

	liveInstance := liveByteInstance.NewInstance(liveConfig)

	// 创建拉流转推组
	// 1. 创建拉流转推组
	// 2. 创建拉流转推任务
	/*	responseGroup, err := liveInstance.CreatePullToPushGroup("default", "default", "test", "test", "Custom")
		if err != nil {
			fmt.Println("Error creating group:", err)
			return
		}

		fmt.Println("Group created successfully:", responseGroup)*/

	responseCreate, err := liveInstance.CreatePullToPushTask(createPullToPushTask)
	if err != nil {
		fmt.Println("Error creating task:", err)
		return
	}

	responseData, err := json.Marshal(responseCreate)
	if err != nil {
		fmt.Println("Error marshaling response:", err)
		return
	}

	fmt.Println("Response data:", string(responseData))

	fmt.Println("Task created successfully:", responseCreate)
}

func ListLiveApiTest() {
	liveConfig := &liveByteInstance.LiveConfig{
		AK: "",
		SK: "",
	}

	liveInstance := liveByteInstance.NewInstance(liveConfig)

	response, err := liveInstance.ListPullToPushTask(1, 20, "")
	if err != nil {
		fmt.Println("Error listing tasks:", err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling response:", err)
		return
	}

	fmt.Println("Response data:", string(responseData))

	fmt.Println("Tasks listed successfully:", response)
}

func DeleteLiveApiTest() {
	liveConfig := &liveByteInstance.LiveConfig{
		AK: "",
		SK: "",
	}

	liveInstance := liveByteInstance.NewInstance(liveConfig)

	response, err := liveInstance.DeletePullToPushTask("")
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling response:", err)
		return
	}

	fmt.Println("Response data:", string(responseData))

	fmt.Println("Task deleted successfully:", response)
}

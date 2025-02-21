package main

import (
	"fmt"
	"github.com/yangwoodstar/NovaCore/src/core/instanceAllocator"
	"log"
)

func main() {
	//InitConfig()
	appIDMap := map[string]instanceAllocator.AppIDMapConfig{
		"app1": {
			AppID:  "app1_id",
			AppKey: "app1_key",
		},
		"app2": {
			AppID:  "app2_id",
			AppKey: "app2_key",
		},
	}

	// 初始化 InstanceManager
	ak := "your_access_key"
	sk := "your_secret_key"
	region := "your_region"
	manager := instanceAllocator.GetInstanceManager(appIDMap, ak, sk, region)

	// 测试获取实例
	key := "app1"
	instance, err := manager.GetInstance(key)
	if err != nil {
		log.Fatalf("Failed to get instance for key %s: %v", key, err)
	}
	fmt.Printf("Instance for key %s created successfully: %+v\n", key, instance.Config)

	// 测试获取另一个实例
	key = "app2"
	instance, err = manager.GetInstance(key)
	if err != nil {
		log.Fatalf("Failed to get instance for key %s: %v", key, err)
	}
	fmt.Printf("Instance for key %s created successfully: %+v\n", key, instance.Config)

	// 测试获取不存在的实例
	key = "app3"
	instance, err = manager.GetInstance(key)
	if err != nil {
		fmt.Printf("Expected error for key %s: %v\n", key, err)
	} else {
		log.Fatalf("Unexpected success for key %s: %+v", key, instance.Config)
	}

	instance, err = manager.GetAppIDInstance("app1_id")
	if err != nil {
		fmt.Printf("Expected error for key %s: %v\n", key, err)
	} else {
		log.Fatalf("Unexpected success for key 1111111 %s: %+v", key, instance.Config)
	}

	// 列出所有实例
	keys := manager.ListInstances()
	fmt.Printf("All instances: %v\n", keys)

	// 移除一个实例
	manager.RemoveInstance("app1")
	keys = manager.ListInstances()
	fmt.Printf("Instances after removal: %v\n", keys)
}

package instanceAllocator

import (
	"errors"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	rtc_v20231101 "github.com/volcengine/volc-sdk-golang/service/rtc/v20231101"
	"sync"
)

type ByteDanceInstance struct {
	Rtc    *rtc_v20231101.Rtc
	Config ByteDanceConfig
}

type ByteDanceConfig struct {
	AK     string
	SK     string
	AppID  string
	AppKey string
	Region string
}

type AppIDMapConfig struct {
	AppID         string `json:"appID"`
	AppKey        string `json:"appKey"`
	Secret        string `json:"secret"`
	PushDomain    string `json:"pushDomain"`
	PullDomain    string `json:"pullDomain"`
	PushSecret    string `json:"pushSecret"`
	PullSecret    string `json:"pullSecret"`
	PushDomainBak string `json:"pushDomainBak"`
	PullDomainBak string `json:"pullDomainBak"`
	PushSecretBak string `json:"pushSecretBak"`
	PullSecretBak string `json:"pullSecretBak"`
	Account       string `json:"account"`
	Bucket        string `json:"bucket"`
	AK            string `json:"ak"`
	SK            string `json:"sk"`
}

// InstanceManager 管理多个字节跳动实例
type InstanceManager struct {
	instances       sync.Map                  // key -> *ByteDanceInstance
	appIDMap        map[string]AppIDMapConfig `mapstructure:"appIDMap"`
	reverseAppIDMap map[string]string         // appID -> key
	region          string
	mu              sync.RWMutex
}

var (
	defaultManager *InstanceManager
	once           sync.Once
)

// GetInstanceManager 获取实例管理器的单例
func GetInstanceManager(appIDMap map[string]AppIDMapConfig, region string) *InstanceManager {
	once.Do(func() {
		defaultManager = &InstanceManager{
			instances:       sync.Map{},
			appIDMap:        appIDMap,
			region:          region,
			reverseAppIDMap: make(map[string]string),
		}
		for key, value := range appIDMap {
			defaultManager.reverseAppIDMap[value.AppID] = key
		}

	})
	return defaultManager
}

func (m *InstanceManager) GetAppIDInstance(appID string) (*ByteDanceInstance, error) {
	key, ok := m.reverseAppIDMap[appID]
	if ok {
		return m.GetInstance(key)
	}
	return nil, errors.New("not found")
}

// GetInstance 根据 key 获取或创建实例
func (m *InstanceManager) GetInstance(key string) (*ByteDanceInstance, error) {
	// 先尝试获取现有实例
	if instance, ok := m.instances.Load(key); ok {
		return instance.(*ByteDanceInstance), nil
	}

	// 获取该 key 对应的配置
	appConfig, err := m.getConfigForKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get config for key %s: %w", key, err)
	}
	// 创建新实例
	m.mu.Lock()
	defer m.mu.Unlock()

	// 双重检查，避免并发创建
	if instance, ok := m.instances.Load(key); ok {
		return instance.(*ByteDanceInstance), nil
	}

	// 创建新实例
	instance := &ByteDanceInstance{
		Rtc:    rtc_v20231101.NewInstance(),
		Config: appConfig,
	}

	// 设置凭证
	instance.Rtc.SetCredential(base.Credentials{
		AccessKeyID:     appConfig.AK,
		SecretAccessKey: appConfig.SK,
		Region:          appConfig.Region,
	})

	// 存储实例
	m.instances.Store(key, instance)

	return instance, nil
}

// getConfigForKey 根据 key 获取配置
func (m *InstanceManager) getConfigForKey(appID string) (ByteDanceConfig, error) {
	// 从配置中获取对应的凭证
	// 这里需要根据你的实际配置管理方式来实现
	appConfig, ok := m.appIDMap[appID]
	if ok == false {
		return ByteDanceConfig{}, errors.New("not exist")
	}

	return ByteDanceConfig{
		AK:     appConfig.AK,
		SK:     appConfig.SK,
		AppID:  appConfig.AppID,
		AppKey: appConfig.AppKey,
		Region: m.region, // 可以根据需要从配置中获取
	}, nil
}

// RemoveInstance 移除指定的实例
func (m *InstanceManager) RemoveInstance(key string) {
	m.instances.Delete(key)
}

// ListInstances 列出所有实例的 key
func (m *InstanceManager) ListInstances() []string {
	var keys []string
	m.instances.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}

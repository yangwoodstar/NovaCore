package liveNacos

import (
	"fmt"
	"sync"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type ConfigListener func(dataId, group, content string)

type ConfigCenter struct {
	client    config_client.IConfigClient
	cache     sync.Map
	listeners sync.Map // map[string][]ConfigListener (key: dataId+group)
	mu        sync.Mutex
}

func NewConfigCenter(cfg *NacosConfig) (*ConfigCenter, error) {
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cfg.ClientConfig,
			ServerConfigs: cfg.ServerConfigs,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create config client: %v", err)
	}

	return &ConfigCenter{
		client: client,
	}, nil
}

func (c *ConfigCenter) GetConfig(dataId, group string) (string, error) {
	if val, ok := c.cache.Load(dataId); ok {
		return val.(string), nil
	}

	content, err := c.client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get config: %v", err)
	}

	c.cache.Store(dataId, content)
	return content, nil
}

func (c *ConfigCenter) AddListener(dataId, group string, listener ConfigListener) error {
	key := dataId + "+" + group

	c.mu.Lock()
	defer c.mu.Unlock()

	// 添加监听器到列表
	var listeners []ConfigListener
	if ls, ok := c.listeners.Load(key); ok {
		listeners = append(ls.([]ConfigListener), listener)
	} else {
		listeners = []ConfigListener{listener}
	}
	c.listeners.Store(key, listeners)

	// 首次监听需要注册到Nacos
	if len(listeners) == 1 {
		err := c.client.ListenConfig(vo.ConfigParam{
			DataId: dataId,
			Group:  group,
			OnChange: func(namespace, group, dataId, content string) {
				c.handleConfigChange(dataId, group, content)
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ConfigCenter) handleConfigChange(dataId, group, content string) {
	// 更新缓存
	c.cache.Store(dataId, content)

	// 触发监听器
	key := dataId + "+" + group
	if ls, ok := c.listeners.Load(key); ok {
		for _, listener := range ls.([]ConfigListener) {
			go listener(dataId, group, content)
		}
	}
}

package liveNacos

import (
	"fmt"
	"sync"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type ConfigCenter struct {
	client config_client.IConfigClient
	cache  sync.Map
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

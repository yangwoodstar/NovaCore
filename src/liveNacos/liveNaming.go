package liveNacos

import (
	"fmt"
	"sync"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type ServiceRegistry struct {
	client naming_client.INamingClient
	cache  sync.Map
}

func NewServiceRegistry(cfg *NacosConfig) (*ServiceRegistry, error) {
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cfg.ClientConfig,
			ServerConfigs: cfg.ServerConfigs,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create naming client: %v", err)
	}

	return &ServiceRegistry{
		client: client,
	}, nil
}

func (s *ServiceRegistry) RegisterInstance(serviceName string, instance vo.RegisterInstanceParam) error {
	instance.ServiceName = serviceName
	success, err := s.client.RegisterInstance(instance)
	if !success || err != nil {
		return fmt.Errorf("failed to register instance: %v", err)
	}
	return nil
}

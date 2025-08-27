package liveNacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

type NacosConfig struct {
	ServerConfigs []constant.ServerConfig
	ClientConfig  *constant.ClientConfig
}

package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func (r *NacosConnConf) NewNamingClient() (naming_client.INamingClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(r.NacosAddr, r.NacosPort),
	}
	cc := constant.ClientConfig{
		NamespaceId:         r.NacosNamespaceId,
		TimeoutMs:           60000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	return client, err
}
func (r *NacosConnConf) NewConfigClient() (config_client.IConfigClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(r.NacosAddr, r.NacosPort),
	}
	cc := constant.ClientConfig{
		NamespaceId:         r.NacosNamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
	var (
		client config_client.IConfigClient
		err    error
	)
	client, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	return client, err
}

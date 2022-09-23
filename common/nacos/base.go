package nacos

import (
	"fmt"
	"goproject/common/util"
	"strings"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

type NacosConnConf struct {
	NacosAddr        string
	NacosNamespaceId string
	NacosGroup       string
	NacosDataIDs     []string
	NacosPort        uint64
}

func (r *NacosConnConf) GetConfigSource(env, configFilePath string) (s []config.Source, err error) {
	if env == "local" {
		fmt.Println("ENV:Local")
		s = append(s, file.NewSource(configFilePath))
		return
	}
	cl, err := r.NewConfigClient()
	if err != nil {
		return
	}
	for _, v := range r.NacosDataIDs {
		s = append(s, NewConfigSource(cl, WithGroup(r.NacosGroup), WithDataID(v)))
	}
	return
}

func NewConfigConnConf() (ncc *NacosConnConf, err error) {
	env := util.GetEnv("ENV", "local")
	if env == "local" {
		err = fmt.Errorf("env error")
		return
	}
	ncc = &NacosConnConf{}
	ncc.NacosAddr = util.GetEnv("NACOS_ADDR", "10.189.208.201")
	t := util.GetEnv("NACOS_PORT", "8848")
	ncc.NacosPort, err = util.StringToUint64(t)
	if err != nil {
		return
	}
	ncc.NacosNamespaceId = util.GetEnv("NACOS_NAMESPACE_ID", "sre-test")
	ncc.NacosGroup = util.GetEnv("NACOS_GROUP", "SRE")
	dataIDsstr := util.GetEnv("NACOS_DATA_IDS", "demo.yaml;registry.yaml;auth.yaml")
	ncc.NacosDataIDs = strings.Split(dataIDsstr, ";")
	return
}

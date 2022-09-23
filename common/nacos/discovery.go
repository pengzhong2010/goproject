package nacos

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
)

func (r *NacosConnConf) Discovery() (reg *nacos.Registry, err error) {
	cl, err := r.NewNamingClient()
	if err != nil {
		return
	}
	reg = nacos.New(cl, nacos.WithGroup(r.NacosGroup))
	return
}

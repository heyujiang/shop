package service

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/heyujiang/shop/config"
)

func GetService(conf config.RegisterCenter, serviceId string) *consulApi.AgentService {
	consulConf := &consulApi.Config{
		Address: conf.Consul.Address,
	}
	client, err := consulApi.NewClient(consulConf)
	if err != nil {
		fmt.Println(err.Error())
	}

	service, _, err := client.Agent().Service(serviceId, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	return service
}

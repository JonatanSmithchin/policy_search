package consul

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

const (
	consulAgentAddress = "121.37.119.47:8500"
)

func FindServer(serviceID string) (string, error) {
	config := consulapi.DefaultConfig()
	config.Address = consulAgentAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	service, _, err := client.Agent().Service(serviceID, nil)
	if err != nil {
		return "", err
	}
	return service.Address + ":" + strconv.Itoa(service.Port), nil
}

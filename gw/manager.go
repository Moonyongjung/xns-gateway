package gw

import "github.com/Moonyongjung/xns-gateway/service"

type servicesManager struct {
	services map[string]service.Service
}

func NewServicesManager(contractServices ...service.Service) *servicesManager {
	s := make(map[string]service.Service)
	for _, service := range contractServices {
		s[service.Name()] = service
	}

	return &servicesManager{
		services: s,
	}
}

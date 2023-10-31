package registrar

import "encoding/json"

type registrarSetControllerExecuteMsg struct {
	SetController struct {
		ControllerAddress string `json:"controller_address"`
		TopDomain         string `json:"top_domain"`
	} `json:"set_controller"`
}

func NewRegistrarSetControllerExecuteMsg(controllerAddress, topDomain string) (string, error) {
	var msg registrarSetControllerExecuteMsg

	msg.SetController.ControllerAddress = controllerAddress
	msg.SetController.TopDomain = topDomain

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type registrarRemoveControllerExecuteMsg struct {
	RemoveController struct {
		ControllerAddress string `json:"controller_address"`
	} `json:"remove_controller"`
}

func NewRegistrarRemoveControllerExecuteMsg(controllerAddress string) (string, error) {
	var msg registrarRemoveControllerExecuteMsg

	msg.RemoveController.ControllerAddress = controllerAddress

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type registrarSetResolverExecuteMsg struct {
	SetResolver struct {
		ResolverAddress string `json:"resolver_address"`
	} `json:"set_resolver"`
}

func NewRegistrarSetResolverExecuteMsg(resolverAddress string) (string, error) {
	var msg registrarSetResolverExecuteMsg

	msg.SetResolver.ResolverAddress = resolverAddress

	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

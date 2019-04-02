package utils

import (
	"os"
	"strconv"

	"github.com/lgsvl/data-marketplace-chaincode-rest/resources"
)

func LoadConfig() (resources.ServerConfig, error) {

	config := resources.ServerConfig{}

	port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 32)
	if err != nil {
		return config, err
	}
	config.Port = int(port)
	return config, nil
}

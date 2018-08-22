package config

import (
	"os"
	"encoding/json"
)

type Configuration struct {
	LBAddress string `json:"lb_address"`
	OSSAddress string `json:"oss_address"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("./config.json ")
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Configuration{}

	err := decoder.Decode(config); if err != nil {
		panic(err)
	}
}

func GetLBAddress() (lbAddress string) {
	return configuration.LBAddress
}

func GetOSSAddress() (ossAddress string) {
	return configuration.OSSAddress
}
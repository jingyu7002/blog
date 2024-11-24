package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type System struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}
type Config struct {
	System System `yaml:"system"`
}

var confPath = "settings.yaml"

func ReadConf() {
	byteData, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(byteData, &config)
	if err != nil {
		panic(fmt.Sprintf("yaml Unmarshal err: %v", err))
	}
	fmt.Printf("%#v\n", config)
}

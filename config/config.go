package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	DB struct {
		URL string `json:"url"`
	} `json:"db"`
}

func MustLoad(filename string, flag string) {
	cfgFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer cfgFile.Close()

	fullCfg := map[string]interface{}{}
	d := yaml.NewDecoder(cfgFile)
	err = d.Decode(fullCfg)
	if err != nil {
		panic(err)
	}

	cfgJSON, err := json.Marshal(fullCfg[flag])
	if err != nil {
		panic(err)
	}

	cfg := &Config{}
	if err = json.Unmarshal(cfgJSON, cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}

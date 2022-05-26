package config

import (
	"os"
	"log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host	      string	`yaml:"host", 	envconfig: "SERVER_HOST"`
		Port	      string	`yaml:"port", 	envconfig: "SERVER_PORT"`
		ReadTimeout   int64	`yaml:"readTimeout",  envconfig: "READ_TIMEOUT"`
		WriteTimeout  int64	`yaml:"writeTimeout",  envconfig: "WRITE_TIMEOUT"`
	
	}	`yaml: "server"`

}


func NewConfig(configFile string) *Config {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal("Can't open config file. Are you sure it's there ?")
	}
	defer file.Close()
	
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		log.Fatal("Error while decoding config file. Exiting the application..")
	}

	return cfg
}

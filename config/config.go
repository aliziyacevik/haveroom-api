package config

import (
	"os"
	"gopkg.in/yaml.v2
)


type Config struct {
	Server struct {
		Host:		`yaml:	"host", 	envconfig: "SERVER_HOST"`
		Port:		`yaml:  "port", 	envconfig: "SERVER_PORT"`
		ReadTimeout:	`yaml:	"readTimeout",  envconfig: "READ_TIMEOUT"`
		WriteTimeout:	`yaml:	"writeTimeout"  envconfig: "WRITE_TIMEOUT"`
	
	}	`yaml: "server"`

}


func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

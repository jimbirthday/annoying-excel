package app

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	Logs     Logs     `yaml:"Logs"`
	Database Database `yaml:"Database"`
	Server   Server   `yaml:"Server"`

	IsDebug bool `yaml:"IsDebug"`
}

type Server struct {
	Port int `yaml:"Port"`
}

type Database struct {
	URL      string `yaml:"Path"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Database string `yaml:"Database"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	ShowSQL  bool   `yaml:"Showsql"`
}

type Logs struct {
	Path string `yaml:"Path"`
}

func initYaml() error {
	err := yamlFile("conf.yaml", &Cnf)
	if err != nil {
		err = yamlFile("configs.yaml", &Cnf)
		if err != nil {
			return err
		}
	}
	return nil
}

func yamlFile(pth string, data interface{}) error {
	fmt.Println("yamlFile:", pth)
	bts, err := ioutil.ReadFile(pth)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bts, data)
}

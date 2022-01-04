package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Configuration struct {
	Database struct{
		Host     string  `yaml:"host"`
		User     string  `yaml:"user"`
		Password string  `yaml:"password"`
		Dbname   string  `yaml:"dbname"`
		Port     int     `yaml:"port"`
		Sslmode  string  `yaml:"sslmode"`
		TimeZone string  `yaml:"timezone"`
	}
	Logger struct{
		BaseDir  string  `yaml:"basedir"`
		HttpLog  string  `yaml:"httplog"`
	}
}

var (
	Config Configuration
)

const (
	confDir  = "conf"
	confYaml = "conf.yaml"
)

func init()  {
	currentDir, _ := os.Getwd()
	configFile, err := ioutil.ReadFile(filepath.Join(currentDir, confDir, confYaml))
	if err != nil {
		fmt.Printf("Failed to read yaml config file: %v\n\n", err)
		return
	}
	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		fmt.Printf("Failed to unmarshal yaml config file: %v\n", err)
		return
	}

	fmt.Printf("%+v\n", Config)
}
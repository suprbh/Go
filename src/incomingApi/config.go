package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// ServerConfig contains the HTTP server configuration.
type ServerConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type config struct {
	Server   ServerConfig `yaml:"server"`
	rootPath string
}

// loadConfig returns a struct representation of the YAML file.
func (c *config) loadConfig(configFile string) (err error) {

	// Set the working directory path so can read things relative to it.
	// c.rootPath, err = filepath.Abs(filepath.Dir(os.Args[0])) // this gives the .exe path
	c.rootPath, err = os.Getwd()
	if err != nil {
		return
	}

	// Defaults
	c.Server.Address = "127.0.0.1"
	c.Server.Port = 8080

	contents, err := ioutil.ReadFile(c.rootPath + "/" + configFile)
	if err == nil {
		err = yaml.Unmarshal(contents, &c)
	}

	return
}

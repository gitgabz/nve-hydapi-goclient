package config

import (
	"errors"
	"flag"
	"log"
	"os"
	"reflect"
)

type Config struct {
	ApiKey string //https://hydapi.nve.no/Users
}

var c Config

func init() {
	c.flagsParse()
	c.environmentParse()
	err := c.validateConfig()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func (c *Config) flagsParse() {
	var apikey string

	flag.StringVar(&apikey, "API_KEY", "", "NVE Api key")
	flag.Parse()

	c.ApiKey = apikey
}

func (c *Config) environmentParse() {
	if c.ApiKey == "" {
		c.ApiKey = os.Getenv("API_KEY")
	}

}

func (c *Config) validateConfig() (err error) {
	apiKey := reflect.ValueOf(c.ApiKey)
	if apiKey.IsZero() {
		err = errors.New("API_KEY is required to access the NVE HydAPI, you may find more information here: https://hydapi.nve.no/UserDocumentation/")
	}
	return
}

func (c *Config) returnApiKey() (a string) {
	a = c.ApiKey
	return
}

func ReturnApiKey() (a string) {
	a = c.returnApiKey()
	return
}

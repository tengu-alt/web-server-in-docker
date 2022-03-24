package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"web-server-in-docker/pkg/models"
)

type Config = models.Config

func GetConfig() string {
	yfile, err := ioutil.ReadFile("../cmd/config.yaml")

	if err != nil {

		log.Fatal(err)
	}
	conf := &Config{}

	err2 := yaml.Unmarshal(yfile, &conf)

	if err2 != nil {

		log.Fatal(err2)
	}
	result := fmt.Sprintf("postgres://%s:%s@sqlserver/%s?sslmode=disable", conf.User, conf.Password, conf.DB)
	return result
}

func GetKey() string {
	yfile, err := ioutil.ReadFile("../cmd/config.yaml")

	if err != nil {

		log.Fatal(err)
	}
	conf := &Config{}
	err2 := yaml.Unmarshal(yfile, &conf)
	if err2 != nil {

		log.Fatal(err2)
	}
	return conf.Key
}

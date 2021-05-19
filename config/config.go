package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Ntd   Ntd    `yaml:"ntd"`
	Test1 int    `yaml:"test1"`
	Test2 string `yaml:"test2"`
}
type Ntd struct {
	Input  Input `yaml:"inputs"`
	Test33 int   `yaml:"test33"`
}
type Input struct {
	NicName string `yaml:"nicName"`
}

var cf *Conf

func ConfigInit() *Conf {
	cf = &Conf{}

	yamlFile, err := ioutil.ReadFile("conf.yml")
	if err != nil {
		log.Println("YamlFile Read Failed", err)
	}
	//fmt.Println(Cf)
	err = yaml.Unmarshal(yamlFile, cf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return cf
}

func Default() *Conf {
	return cf
}

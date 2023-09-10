package core

import (
	"../config"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//InitConf 读取yaml文件的配置
func Initconf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf  err: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init sucess.")
	fmt.Println(c)
}

package config

import (
    "encoding/json"
    "os"
    "fmt"
)

var Conf Config

type Config struct {
	Mysql struct{
		User string `json:"user"`
		Password string `json:"password"`
		Host string `json:"host"`
		Port int `json:"port"`
		Db string `json:"db"`
	} `json:"mysql"`

	Redis struct{
		Host string `json:"host"`
		Port int `json:"port"`
	} `json:"redis"`
}

func Load(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil{
		panic(err)
	}
	decoder := json.NewDecoder(file)
	Conf = Config{}
	err = decoder.Decode(&Conf)
	if err != nil {
	  fmt.Println("error:", err)
	}
	fmt.Println(Conf.Mysql) // output: [UserA, UserB]
}

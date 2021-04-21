package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type configData struct {
	Keywords string `json:"keywords"`
	Filepath string `json:"filepath"`
}

var ConfigData configData

func GetConfig() {
	filepath, _ := os.Getwd()
	filepath = filepath + "/config.json"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	config, _ := io.ReadAll(file)
	bytes := []byte(config)
	if err := json.Unmarshal(bytes, &ConfigData); err != nil {
		fmt.Println("Invalid config: ", err.Error())
		os.Exit(-1)
	}
}

func Getfilelist() []string {
	GetConfig()
	filepath, _ := os.Getwd()
	filepath = filepath + "/config.json"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	config, _ := io.ReadAll(file)
	bytes := []byte(config)
	if err := json.Unmarshal(bytes, &ConfigData); err != nil {
		fmt.Println("Invalid config: ", err.Error())
		os.Exit(-1)
	}

	files, err := os.ReadDir(ConfigData.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	var filelist []string
	for _, filename := range files {
		ok := strings.HasSuffix(filename.Name(), ".xlsx")
		if ok {
			filelist = append(filelist, ConfigData.Filepath+"/"+filename.Name())
		}
	}
	return filelist
}

package main

import (
	"encoding/json"
	"errors"
	"log"
)

type configPlugin struct {}

func (c configPlugin) ProcessConfig(configMap map[string]json.RawMessage) (string, error) {
	configJson, ok := configMap["srb2kart"]
	if !ok {
		log.Printf("Could not find hostmod config")
		return "", errors.New("Could not find config")
	}
	
	var srb2kart = srb2kart{}
	err := json.Unmarshal(configJson, &srb2kart)
	if err != nil {
		log.Print(err)
		return "", err
	}
	log.Printf("Processing config")
	return mapSrb2kart(&srb2kart), nil
}

var ConfigPlugin configPlugin

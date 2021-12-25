package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

func newSolution() *solution {
	return &solution{}
}

func (sol *solution) loadConfig() error {
	jsonBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.New("failed to load config file: " + err.Error())
	}

	err = json.Unmarshal(jsonBytes, &sol.Config)
	if err != nil {
		return errors.New("failed to decode config: " + err.Error())
	}
	return nil
}

func main() {
	app := newSolution()
	err := checkErrors(
		app.loadConfig,
	)
	if err != nil {
		log.Fatalln(err)
	}
}

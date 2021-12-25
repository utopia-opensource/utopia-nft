package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	utopiago "github.com/Sagleft/utopialib-go"
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

func (sol *solution) connect() error {
	sol.Client = utopiago.UtopiaClient{
		Protocol: sol.Client.Protocol,
		Token:    sol.Client.Token,
		Host:     sol.Client.Host,
		Port:     sol.Client.Port,
	}
	isConnected := sol.Client.CheckClientConnection()
	if !isConnected {
		return errors.New("failed to connect to Utopia client")
	}
	return nil
}

func main() {
	app := newSolution()
	err := checkErrors(
		app.loadConfig,
		app.connect,
	)
	if err != nil {
		log.Fatalln(err)
	}
}

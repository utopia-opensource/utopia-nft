package main

import (
	"errors"

	utopiago "github.com/Sagleft/utopialib-go"
)

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

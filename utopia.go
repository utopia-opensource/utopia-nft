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

// return image in base64
func (sol *solution) getRandomStickerImage() (string, error) {
	stickers, err := sol.Client.GetStickerNamesByCollection(sol.Config.Task.StickerCollection)
	if err != nil {
		return "", nil
	}

	stickerName := stickers[getRandomInt(len(stickers))]

	return sol.Client.GetStickerImage(
		sol.Config.Task.StickerCollection,
		stickerName,
	)
}

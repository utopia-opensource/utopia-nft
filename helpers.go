package main

import (
	"math/rand"
	"time"
)

type errorFunc func() error

func checkErrors(errChecks ...errorFunc) error {
	for _, errFunc := range errChecks {
		err := errFunc()
		if err != nil {
			return err
		}
	}
	return nil
}

// [0,max)
func getRandomInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

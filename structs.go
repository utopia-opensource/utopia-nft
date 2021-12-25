package main

import utopiago "github.com/Sagleft/utopialib-go"

type solution struct {
	Config config
	Client utopiago.UtopiaClient
}

type config struct {
	Utopia clientConnectionConfig `json:"utopia"`
	Task   taskConfig             `json:"task"`
}

type clientConnectionConfig struct {
	Protocol string `json:"protocol"`
	Token    string `json:"token"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type taskConfig struct{}

package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	CheckDelay int       `json:"check_delay"`
	Offset     int       `json:"offset"`
	Displays   []Display `json:"displays"`
}

type Display struct {
	Id          string  `json:"id"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	RefreshRate int     `json:"refresh_rate"`
	PositionX   int     `json:"position_x"`
	PositionY   int     `json:"position_y"`
	Scale       float32 `json:"scale"`
	GameDisplay bool    `json:"game_display"`
}

func loadConfig(path string) Config {
	var config Config
	configFile, err := os.Open(path)
	defer configFile.Close()
	check(err)
	parser := json.NewDecoder(configFile)
	parser.Decode(&config)
	return config
}

func findGameDisplayIndex(displays []Display) int {
	for index, display := range displays {
		if display.GameDisplay {
			return index
		}
	}

	panic("there's no display with game_display: true")
	return -1
}

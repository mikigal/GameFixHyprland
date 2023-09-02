package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var previousRunning = false

func main() {
	config := loadConfig("./config.json")
	displays := config.Displays
	gameDisplayIndex := findGameDisplayIndex(displays)

	logMessage("Loaded configuration!")

	for {
		time.Sleep(time.Duration(config.CheckDelay) * time.Second)

		running := isLeagueRunning()
		if running == previousRunning {
			continue
		}

		previousRunning = running
		if running {
			logMessage("Started League, applying displays fix...")
		} else {
			logMessage("Stopped League, applying default displays layout...")
		}

		for index, display := range displays {
			if index < gameDisplayIndex {
				continue
			}

			calculatedPositionX := display.PositionX
			if running {
				calculatedPositionX = display.PositionX + (config.Offset * index)
			}

			parameters := display.Id + ", " +
				strconv.Itoa(display.Width) + "x" + strconv.Itoa(display.Height) + "@" + strconv.Itoa(display.RefreshRate) + ", " +
				strconv.Itoa(calculatedPositionX) + "x" + strconv.Itoa(display.PositionY) + ", " +
				fmt.Sprintf("%f", display.Scale)

			logMessage("Executing: " + "hyprctl keyword monitor " + parameters)
			output, err := exec.Command("hyprctl", "keyword monitor "+parameters).Output()
			check(err)
			logMessage("hyprctl output: " + strings.TrimSpace(string(output)))
		}
	}
}

func isLeagueRunning() bool {
	processes := exec.Command("ps", "ax")
	outputRaw, err := processes.Output()
	check(err)
	output := string(outputRaw)
	return strings.Contains(output, "LeagueClient.exe")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func logMessage(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("02.01.2006 15:04:05"), message)
}

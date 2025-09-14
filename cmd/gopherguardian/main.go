package main

import (
	"fmt"
	"gopehrguardian/pkg/config"
	"gopehrguardian/pkg/flags"
	"gopehrguardian/pkg/monitor"
	"log"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	parsedFlags, err := flags.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	loadedConfig, err := config.LoadConfig(parsedFlags.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("config: %+v\n", loadedConfig)

	var wg sync.WaitGroup
	wg.Add(1)

	for _, target := range loadedConfig.Targets {
		timedChecker := monitor.TimedChecker{
			Checker: monitor.GetChecker(&target),
		}

		go monitor.Monitor(&timedChecker, &target)
	}

	wg.Wait()
}

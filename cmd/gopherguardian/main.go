package main

import (
	"fmt"
	"gopehrguardian/pkg/config"
	"gopehrguardian/pkg/flags"
	"gopehrguardian/pkg/monitor"
	"log"
	"sync"
)

func main() {
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

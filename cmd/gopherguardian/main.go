package gopherguardian

import (
	"fmt"
	"gopehrguardian/pkg/config"
	"gopehrguardian/pkg/flags"
	"log"
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
}

package alert

import (
	"gopehrguardian/pkg/config"
	"log"
)

func AlertTarget(t *config.Target) {
	log.Printf("down: %s - %s\n", t.Name, t.Address)
}

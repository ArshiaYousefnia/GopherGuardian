package alert

import (
	"fmt"
	"gopehrguardian/pkg/config"
	"log"
	"time"
)

func AlertTarget(t *config.Target) {
	log.Printf("down: %s - %s\n", t.Name, t.Address)
	s := fmt.Sprintf("down: %s - %s - %s\n", t.Name, t.Address, time.Now())

	if t.Alert.Telegram != "" {
		telegramAlert(t, s)
	}
	if t.Alert.Email != "" {
		emailAlert(t, s)
	}
}

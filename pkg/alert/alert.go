package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopehrguardian/pkg/config"
	"log"
	"net/http"
	"os"
	"time"
)

func AlertTarget(t *config.Target) {
	log.Printf("down: %s - %s\n", t.Name, t.Address)
	s := fmt.Sprintf("down: %s - %s - %s\n", t.Name, t.Address, time.Now())

	telegramAlert(t, s)
}

func telegramAlert(target *config.Target, text string) {
	bot := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := target.Alert.Telegram

	requestUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", bot)

	values := map[string]string{
		"chat_id": chatID,
		"text":    text,
	}

	jsonParams, err := json.Marshal(values)
	if err != nil {
		log.Println("Error marshalling params:", err)
		return
	}

	req, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonParams))

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Printf("Alerting: %s - %s error sending request: %s\n", target.Name, target.Address, err)
	}
}

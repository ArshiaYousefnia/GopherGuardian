package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopehrguardian/pkg/config"
	"log"
	"net/http"
	"os"
)

const url = "https://api.telegram.org/bot%s/sendMessage"

func telegramAlert(target *config.Target, text string) {
	bot := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := target.Alert.Telegram

	requestUrl := fmt.Sprintf(url, bot)

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

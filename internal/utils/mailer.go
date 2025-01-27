package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func SendToTGBot(message string) error {
	// Replace with your bot token
	chatIDSt := os.Getenv("CHAT_ID")
	chatID, _ := strconv.ParseUint(chatIDSt, 10, 64)
	token := os.Getenv("TG_API_KEY")

	// Telegram API endpoint
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	// Request payload
	payload := map[string]interface{}{
		"chat_id": chatID,
		"text":    message,
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Send the request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	// Check if the message was sent successfully
	if !result["ok"].(bool) {
		return fmt.Errorf("failed to send message: %v", result["description"])
	}

	return nil
}

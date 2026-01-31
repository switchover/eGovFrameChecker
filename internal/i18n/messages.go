package i18n

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
)

//go:embed messages/*.json
var messageFS embed.FS

type ErrorMessage struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type Messages struct {
	Errors map[string]ErrorMessage `json:"errors"`
}

var messageCache = make(map[string]*Messages)

func loadMessages(locale string) (*Messages, error) {
	if cached, ok := messageCache[locale]; ok {
		return cached, nil
	}

	filename := fmt.Sprintf("messages/messages_%s.json", locale)
	data, err := messageFS.ReadFile(filename)
	if err != nil {
		// fallback to default locale
		data, err = messageFS.ReadFile("messages/messages.json")
		if err != nil {
			return nil, err
		}
	}

	var msgs Messages
	if err := json.Unmarshal(data, &msgs); err != nil {
		return nil, err
	}

	messageCache[locale] = &msgs

	return &msgs, nil
}

func GetErrorMessage(code, locale string) (*ErrorMessage, error) {
	messages, err := loadMessages(locale)
	if err != nil {
		return nil, fmt.Errorf("failed to load error messages: %w", err)
	}

	if message, ok := messages.Errors[code]; ok {
		return &message, nil
	}
	return nil, errors.New("error message not found")
}

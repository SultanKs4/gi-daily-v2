package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SultanKs4/gi-daily/setting"
)

func SendMessageTelegram(text string, photo string) (map[string]interface{}, error) {
	if setting.Environment.TelegramBotToken == "" || setting.Environment.TelegramOwnerId == "" {
		return nil, errors.New("telegram error: Bot Token or Owner Id not set")
	}

	uri := fmt.Sprintf("https://api.telegram.org/bot%v/", setting.Environment.TelegramBotToken)
	data := map[string]string{"chat_id": setting.Environment.TelegramOwnerId}
	if photo == "" {
		uri += "sendMessage"
		data["text"] = text
	} else {
		uri += "sendPhoto"
		data["photo"] = photo
		data["caption"] = text
	}
	json, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("create json data error: %v", err)
	}

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json))
	if err != nil {
		return nil, fmt.Errorf("telegram request error: %v", err)
	}
	defer resp.Body.Close()

	jsonResp, err := Bodytojson(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("body to json response error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("telegram respond error: %v - %v (%v)", resp.StatusCode, resp.Status, jsonResp["description"])
	}

	return jsonResp, nil
}

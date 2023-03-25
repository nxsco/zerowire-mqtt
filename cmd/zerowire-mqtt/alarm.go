package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type alarm struct {
	Bankstates string `json:"bankstates"`
}

func GetAlarm(zw *config) (*alarm, error) {
	client := http.Client{}
	statusURL := fmt.Sprintf("http://%s/user/status.json", zw.alarmHost)
	resp, err := client.PostForm(statusURL, url.Values{
		"lgname": {zw.alarmUser},
		"lgpin":  {zw.alarmPass},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var a *alarm
	err = json.Unmarshal(body, &a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s alarm) GetStatus() (string, error) {
	if len(s.Bankstates) < 10 {
		return "", fmt.Errorf("invalid zerowire alarm status")
	}
	if string(s.Bankstates[9]) == "1" {
		return "pending", nil
	}
	if string(s.Bankstates[5]) == "1" {
		return "armed_home", nil
	}
	if string(s.Bankstates[7]) == "1" {
		return "armed_away", nil
	}

	return "disarmed", nil
}

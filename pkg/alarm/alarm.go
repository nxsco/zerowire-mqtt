package alarm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/nxsco/zerowire-mqtt/pkg/config"
)

type Alarm struct {
	client     http.Client
	statusURL  string
	alarmUser  string
	alarmPass  string
	Bankstates string `json:"bankstates"`
}

func New(zw *config.Config) *Alarm {
	return &Alarm{
		client:    http.Client{},
		statusURL: fmt.Sprintf("http://%s/user/status.json", zw.AlarmHost),
		alarmUser: zw.AlarmUser,
		alarmPass: zw.AlarmPass,
	}
}

func (s Alarm) GetStatus() (string, error) {
	resp, err := s.client.PostForm(s.statusURL, url.Values{
		"lgname": {s.alarmUser},
		"lgpin":  {s.alarmPass},
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return "", err
	}

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

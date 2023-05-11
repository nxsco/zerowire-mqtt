package config

import (
	"fmt"
	"os"
	"strconv"
)

const defaultAlarmUser = "User 1"
const defaultAlarmPollInterval = 30 // seconds
const defaultMQTTStateTopic = "home/alarm"

type Config struct {
	AlarmHost         string
	AlarmUser         string
	AlarmPass         string
	AlarmPollInterval int64
	MqttHost          string
	MqttUser          string
	MqttPass          string
	MqttStateTopic    string
}

func New() (*Config, error) {
	conf := &Config{
		AlarmHost:      os.Getenv("ZW_ALARM_HOST"),
		AlarmUser:      os.Getenv("ZW_ALARM_USER"),
		AlarmPass:      os.Getenv("ZW_ALARM_PASS"),
		MqttHost:       os.Getenv("ZW_MQTT_HOST"),
		MqttUser:       os.Getenv("ZW_MQTT_USER"),
		MqttPass:       os.Getenv("ZW_MQTT_PASS"),
		MqttStateTopic: os.Getenv("ZW_MQTT_STATE_TOPIC"),
	}

	// Check for the required vars
	missingConfig := []string{}
	if conf.AlarmHost == "" {
		missingConfig = append(missingConfig, "ZW_ALARM_HOST")
	}
	if conf.AlarmPass == "" {
		missingConfig = append(missingConfig, "ZW_ALARM_PASS")
	}
	if conf.MqttHost == "" {
		missingConfig = append(missingConfig, "ZW_MQTT_HOST")
	}

	if len(missingConfig) > 0 {
		return nil, fmt.Errorf("missing required config %s", missingConfig)
	}

	// Default some other values
	if conf.AlarmUser == "" {
		conf.AlarmUser = defaultAlarmUser
	}
	if conf.MqttStateTopic == "" {
		conf.MqttStateTopic = defaultMQTTStateTopic
	}

	// Setup the polling interval
	intervalVar := os.Getenv("ZW_ALARM_POLL_INTERVAL")
	if intervalVar != "" {
		value, err := strconv.ParseInt(intervalVar, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("unable to parse ZW_ALARM_POLL_INTERVAL [%s]", intervalVar)
		}
		conf.AlarmPollInterval = value
	} else {
		conf.AlarmPollInterval = defaultAlarmPollInterval
	}

	return conf, nil
}

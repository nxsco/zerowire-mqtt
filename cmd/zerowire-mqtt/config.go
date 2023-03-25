package main

import (
	"fmt"
	"os"
	"strconv"
)

const defaultAlarmUser = "User 1"
const defaultAlarmPollInterval = 30 // seconds
const defaultMQTTStateTopic = "home/alarm"

type config struct {
	alarmHost         string
	alarmUser         string
	alarmPass         string
	alarmPollInterval int64
	mqttHost          string
	mqttUser          string
	mqttPass          string
	mqttStateTopic    string
}

func getConfig() (*config, error) {
	conf := &config{
		alarmHost:      os.Getenv("ZW_ALARM_HOST"),
		alarmUser:      os.Getenv("ZW_ALARM_USER"),
		alarmPass:      os.Getenv("ZW_ALARM_PASS"),
		mqttHost:       os.Getenv("ZW_MQTT_HOST"),
		mqttUser:       os.Getenv("ZW_MQTT_USER"),
		mqttPass:       os.Getenv("ZW_MQTT_PASS"),
		mqttStateTopic: os.Getenv("ZW_MQTT_STATE_TOPIC"),
	}

	// Check for the required vars
	missingConfig := []string{}
	if conf.alarmHost == "" {
		missingConfig = append(missingConfig, "ZW_ALARM_HOST")
	}
	if conf.alarmPass == "" {
		missingConfig = append(missingConfig, "ZW_ALARM_PASS")
	}
	if conf.mqttHost == "" {
		missingConfig = append(missingConfig, "ZW_MQTT_HOST")
	}

	if len(missingConfig) > 0 {
		return nil, fmt.Errorf("missing required config %s", missingConfig)
	}

	// Default some other values
	if conf.alarmUser == "" {
		conf.alarmUser = defaultAlarmUser
	}
	if conf.mqttStateTopic == "" {
		conf.mqttStateTopic = defaultMQTTStateTopic
	}

	// Setup the polling interval
	intervalVar := os.Getenv("ZW_ALARM_POLL_INTERVAL")
	if intervalVar != "" {
		value, err := strconv.ParseInt(intervalVar, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("unable to parse ZW_ALARM_POLL_INTERVAL [%s]", intervalVar)
		}
		conf.alarmPollInterval = value
	} else {
		conf.alarmPollInterval = defaultAlarmPollInterval
	}

	return conf, nil
}

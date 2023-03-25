package main

import (
	"os"
	"testing"
)

func TestGetConfigValid(t *testing.T) {
	t.Setenv("ZW_ALARM_HOST", "10.0.0.3")
	t.Setenv("ZW_ALARM_USER", "user1")
	t.Setenv("ZW_ALARM_PASS", "pass1")
	t.Setenv("ZW_ALARM_POLL_INTERVAL", "9999")
	t.Setenv("ZW_MQTT_HOST", "10.0.0.2")
	t.Setenv("ZW_MQTT_USER", "user2")
	t.Setenv("ZW_MQTT_PASS", "pass2")
	t.Setenv("ZW_MQTT_STATE_TOPIC", "some/random/topic")

	conf, err := getConfig()
	if err != nil {
		t.Errorf("wanted valid cofig got %s", err)
	}
	if conf.alarmHost != os.Getenv("ZW_ALARM_HOST") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_ALARM_HOST"), conf.alarmHost)
	}
	if conf.alarmUser != os.Getenv("ZW_ALARM_USER") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_ALARM_USER"), conf.alarmUser)
	}
	if conf.alarmPass != os.Getenv("ZW_ALARM_PASS") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_ALARM_PASS"), conf.alarmPass)
	}
	if conf.alarmPollInterval != 9999 {
		t.Errorf("wanted %d got %d", 9999, conf.alarmPollInterval)
	}
	if conf.mqttHost != os.Getenv("ZW_MQTT_HOST") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_MQTT_HOST"), conf.mqttHost)
	}
	if conf.mqttUser != os.Getenv("ZW_MQTT_USER") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_MQTT_USER"), conf.mqttUser)
	}
	if conf.mqttPass != os.Getenv("ZW_MQTT_PASS") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_MQTT_PASS"), conf.mqttPass)
	}
	if conf.mqttStateTopic != os.Getenv("ZW_MQTT_STATE_TOPIC") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_MQTT_STATE_TOPIC"), conf.mqttStateTopic)
	}
}

func TestGetConfigMissingVars(t *testing.T) {
	t.Setenv("ZW_ALARM_HOST", "")
	t.Setenv("ZW_ALARM_USER", "")
	t.Setenv("ZW_ALARM_PASS", "")
	t.Setenv("ZW_MQTT_HOST", "")
	t.Setenv("ZW_MQTT_USER", "")
	t.Setenv("ZW_MQTT_PASS", "")

	conf, err := getConfig()
	if conf != nil {
		t.Error("wanted err got config")
	}
	if err == nil {
		t.Error("wanted err got config")
	}
}

func TestGetConfigInvalidInterval(t *testing.T) {
	t.Setenv("ZW_ALARM_HOST", "10.0.0.3")
	t.Setenv("ZW_ALARM_PASS", "pass1")
	t.Setenv("ZW_MQTT_HOST", "10.0.0.2")
	t.Setenv("ZW_ALARM_POLL_INTERVAL", "blah")

	conf, err := getConfig()
	if conf != nil {
		t.Error("wanted err got config")
	}
	if err == nil {
		t.Error("wanted err got config")
	}
}

func TestGetConfigDefaults(t *testing.T) {
	t.Setenv("ZW_ALARM_HOST", "10.0.0.3")
	t.Setenv("ZW_ALARM_USER", "")
	t.Setenv("ZW_ALARM_PASS", "pass1")
	t.Setenv("ZW_MQTT_HOST", "10.0.0.2")

	conf, err := getConfig()
	if err != nil {
		t.Errorf("wanted valid cofig got %s", err)
	}
	if conf.alarmUser != defaultAlarmUser {
		t.Errorf("wanted %s got %s", defaultAlarmUser, conf.alarmUser)
	}
	if conf.mqttStateTopic != defaultMQTTStateTopic {
		t.Errorf("wanted %s got %s", defaultMQTTStateTopic, conf.mqttStateTopic)
	}
	if conf.alarmPollInterval != defaultAlarmPollInterval {
		t.Errorf("wanted %d got %d", defaultAlarmPollInterval, conf.alarmPollInterval)
	}
}

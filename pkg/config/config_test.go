package config

import (
	"os"
	"testing"
)

func TestNewValid(t *testing.T) {
	t.Setenv("ZW_ALARM_HOST", "10.0.0.3")
	t.Setenv("ZW_ALARM_USER", "user1")
	t.Setenv("ZW_ALARM_PASS", "pass1")
	t.Setenv("ZW_ALARM_POLL_INTERVAL", "9999")
	t.Setenv("ZW_MQTT_HOST", "10.0.0.2")
	t.Setenv("ZW_MQTT_USER", "user2")
	t.Setenv("ZW_MQTT_PASS", "pass2")
	t.Setenv("ZW_MQTT_STATE_TOPIC", "some/random/topic")

	conf, err := New()
	if err != nil {
		t.Errorf("wanted valid cofig got %s", err)
	}
	if conf.AlarmHost != os.Getenv("ZW_ALARM_HOST") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_ALARM_HOST"), conf.AlarmHost)
	}
	if conf.AlarmUser != os.Getenv("ZW_ALARM_USER") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_ALARM_USER"), conf.AlarmUser)
	}
	if conf.AlarmPass != os.Getenv("ZW_ALARM_PASS") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_ALARM_PASS"), conf.AlarmPass)
	}
	if conf.AlarmPollInterval != 9999 {
		t.Errorf("wanted %d got %d", 9999, conf.AlarmPollInterval)
	}
	if conf.MqttHost != os.Getenv("ZW_MQTT_HOST") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_MQTT_HOST"), conf.MqttHost)
	}
	if conf.MqttUser != os.Getenv("ZW_MQTT_USER") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_MQTT_USER"), conf.MqttUser)
	}
	if conf.MqttPass != os.Getenv("ZW_MQTT_PASS") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_MQTT_PASS"), conf.MqttPass)
	}
	if conf.MqttStateTopic != os.Getenv("ZW_MQTT_STATE_TOPIC") {
		t.Errorf("wanted %s got %s", os.Getenv("ZW_MQTT_STATE_TOPIC"), conf.MqttStateTopic)
	}
}

func TestNewMissingVars(t *testing.T) {
	t.Setenv("ZW_ALARM_HOST", "")
	t.Setenv("ZW_ALARM_USER", "")
	t.Setenv("ZW_ALARM_PASS", "")
	t.Setenv("ZW_MQTT_HOST", "")
	t.Setenv("ZW_MQTT_USER", "")
	t.Setenv("ZW_MQTT_PASS", "")

	conf, err := New()
	if conf != nil {
		t.Error("wanted err got config")
	}
	if err == nil {
		t.Error("wanted err got config")
	}
}

func TestNewInvalidInterval(t *testing.T) {
	t.Setenv("ZW_ALARM_HOST", "10.0.0.3")
	t.Setenv("ZW_ALARM_PASS", "pass1")
	t.Setenv("ZW_MQTT_HOST", "10.0.0.2")
	t.Setenv("ZW_ALARM_POLL_INTERVAL", "blah")

	conf, err := New()
	if conf != nil {
		t.Error("wanted err got config")
	}
	if err == nil {
		t.Error("wanted err got config")
	}
}

func TestNewDefaults(t *testing.T) {
	t.Setenv("ZW_ALARM_HOST", "10.0.0.3")
	t.Setenv("ZW_ALARM_USER", "")
	t.Setenv("ZW_ALARM_PASS", "pass1")
	t.Setenv("ZW_MQTT_HOST", "10.0.0.2")

	conf, err := New()
	if err != nil {
		t.Errorf("wanted valid cofig got %s", err)
	}
	if conf.AlarmUser != defaultAlarmUser {
		t.Errorf("wanted %s got %s", defaultAlarmUser, conf.AlarmUser)
	}
	if conf.MqttStateTopic != defaultMQTTStateTopic {
		t.Errorf("wanted %s got %s", defaultMQTTStateTopic, conf.MqttStateTopic)
	}
	if conf.AlarmPollInterval != defaultAlarmPollInterval {
		t.Errorf("wanted %d got %d", defaultAlarmPollInterval, conf.AlarmPollInterval)
	}
}

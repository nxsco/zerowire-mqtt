package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/nxsco/zerowire-mqtt/pkg/config"
)

func mqttClient(zw *config.Config) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:1883", zw.MqttHost))
	opts.SetClientID("go_mqtt_client")
	if zw.MqttUser != "" && zw.MqttPass != "" {
		opts.SetUsername(zw.MqttUser)
		opts.SetPassword(zw.MqttPass)
	}
	return mqtt.NewClient(opts)
}

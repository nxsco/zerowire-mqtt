package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func mqttClient(zw *config) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:1883", zw.mqttHost))
	opts.SetClientID("go_mqtt_client")
	if zw.mqttUser != "" && zw.mqttPass != "" {
		opts.SetUsername(zw.mqttUser)
		opts.SetPassword(zw.mqttPass)
	}
	return mqtt.NewClient(opts)
}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nxsco/zerowire-mqtt/pkg/alarm"
	"github.com/nxsco/zerowire-mqtt/pkg/config"
)

func main() {
	zwConfig, err := config.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := mqttClient(zwConfig)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	zalarm := alarm.New(zwConfig)
	channel := make(chan string)
	go monitorAlarm(zalarm, channel, zwConfig.AlarmPollInterval)

	for {
		token := client.Publish(zwConfig.MqttStateTopic, 0, false, <-channel)
		token.Wait()
	}
}

func monitorAlarm(a *alarm.Alarm, c chan string, i int64) {
	for {
		status, err := a.GetStatus()
		if err == nil {
			c <- status
		} else {
			fmt.Println(err)
		}

		time.Sleep(time.Duration(i) * time.Second)
	}
}

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	zwConfig, err := getConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := mqttClient(zwConfig)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	for {
		a, _ := GetAlarm(zwConfig)
		status, _ := a.GetStatus()

		token := client.Publish(zwConfig.mqttStateTopic, 0, false, status)
		token.Wait()

		time.Sleep(time.Duration(zwConfig.alarmPollInterval) * time.Second)
	}
}

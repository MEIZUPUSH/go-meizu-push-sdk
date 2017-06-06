package goPushSdk

import (
	"fmt"
	"testing"
)

func printMessage(message string) {
	fmt.Println("current message " + message)
}

func TestGetRegisterSwitch(t *testing.T) {
	message := GetRegisterSwitch(APP_ID, PUSH_ID, APP_KEY)
	printMessage(message.message)
	if message.code != 200 {
		t.Error("Status Code not 200")
	}
}

func TestChangeRegisterSwitch(t *testing.T) {
	message := ChangeRegisterSwitch(APP_ID, PUSH_ID, NOTIFICATION_SWITCH, false, APP_KEY)
	printMessage(message.message)
	if message.code != 200 {
		t.Error("Status Code not 200")
	}

	message1 := ChangeRegisterSwitch(APP_ID, PUSH_ID, THROUGH_MESSAGE_SWITCH, false, APP_KEY)
	printMessage(message1.message)
	if message1.code != 200 {
		t.Error("Status Code not 200")
	}
}

func TestChangeAllSwitch(t *testing.T) {
	message := ChangeAllSwitch(APP_ID, PUSH_ID, false, APP_KEY)
	printMessage(message.message)
	if message.code != 200 {
		t.Error("Status Code not 200")
	}
}

func TestSubscribeAlias(t *testing.T) {
	message := SubscribeAlias(APP_ID, PUSH_ID, "go-push", APP_KEY)
	printMessage(message.message)
	if message.code != 200 {
		t.Error("Status Code not 200")
	}
}

package goPushSdk

import (
	"testing"
	"fmt"
)

func TestGetRegisterSwitch(t *testing.T) {
	message := GetRegisterSwitch(APP_ID,PUSH_ID,APP_KEY)
	fmt.Println("current message "+message.message)
	if(message.code != 200){
		t.Error("Status Code not 200")
	}
}
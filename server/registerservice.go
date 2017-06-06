package goPushSdk

import (
	"github.com/ddliu/go-httpclient"
	"fmt"
)

const (
	getRegisterSwitch = PUSH_API_SERVER + "/garcia/api/server/message/getRegisterSwitch"
)

const (
	APP_ID = "100999"
	APP_KEY = "531732bc45324098978bf41c6954c09e"
	PUSH_ID = "862891030007404100999"
)

//获取订阅开关状态
func GetRegisterSwitch(appId string, pushId string, appKey string) PushResponse {
	registerMap := map[string]string{
		"appId":  appId,
		"pushId": pushId,
	}

	sign := GenerateSign(registerMap, appKey)
	registerMap["sign"] = sign

	for k,v := range registerMap {
		fmt.Println(k, ":", v)
	}

	res, err := httpclient.Get(getRegisterSwitch, registerMap)

	return ResolvePushResponse(res,err)

}



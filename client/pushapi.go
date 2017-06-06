package goPushSdk

import (
	"github.com/ddliu/go-httpclient"
)



//订阅接口
func Register(appId string, appKey string, deviceId string) PushResponse {
	registerRequestMap := map[string]string{
		"appId":    appId,
		"deviceId": deviceId,
	}

	res, err := httpclient.Post(SERVER+"message/registerPush", map[string]string{
		"appId":    appId,
		"deviceId": deviceId,
		"sign":     GenerateSign(registerRequestMap, appKey),
	})
	return ResolvePushResponse(res,err)
}



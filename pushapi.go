package goPushSdk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/ddliu/go-httpclient"
)

const (
	SERVER      = "https://api-push.meizu.com/garcia/api/client/"
	STATUS_CODE = 200
)

type PushResponse struct {
	code    int
	message string
}

//订阅接口
func Register(appId string, appKey string, deviceId string) PushResponse {
	registerRequestMap := map[string]string{
		"appId":    appId,
		"deviceId": deviceId,
	}

	res, err := httpclient.Post(SERVER+"message/registerPush", map[string]string{
		"appId":    appId,
		"deviceId": deviceId,
		"sign":     generateSign(registerRequestMap, appKey),
	})
	var response PushResponse
	if err != nil {
		response = PushResponse{
			code:    0,
			message: err.Error(),
		}
	} else {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		body := buf.String()

		response = PushResponse{
			code:    STATUS_CODE,
			message: body,
		}
	}
	return response
}

// md5 sign
func generateSign(params map[string]string, appKey string) string {
	var signStr string
	if params != nil {
		for k, v := range params {
			signStr += k + "=" + v
		}
		signStr += appKey
	}
	hasher := md5.New()
	hasher.Write([]byte(signStr))
	return hex.EncodeToString(hasher.Sum(nil))
}

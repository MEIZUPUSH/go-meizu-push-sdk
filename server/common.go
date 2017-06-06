package goPushSdk

import (
	"fmt"
	"encoding/hex"
	"github.com/ddliu/go-httpclient"
	"bytes"
	"crypto/md5"
)


const (
	SERVER      = "https://api-push.meizu.com/garcia/api/client/"
	STATUS_CODE = 200
	//服务端SDK调用API的应用的私钥Secret Key为 appSecret
	PUSH_API_SERVER = "https://api-push.meizu.com"
)

type PushResponse struct {
	code    int
	message string
}


// md5 sign
func GenerateSign(params map[string]string, appKey string) string {
	var signStr string
	if params != nil {
		for k, v := range params {
			signStr += k + "=" + v
		}
		signStr += appKey
		fmt.Println("signStr ",signStr)
	}
	hasher := md5.New()
	hasher.Write([]byte(signStr))
	return hex.EncodeToString(hasher.Sum(nil))
}


//resolve push response
func ResolvePushResponse(res *httpclient.Response, err error) PushResponse {
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

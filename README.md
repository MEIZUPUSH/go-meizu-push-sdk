# go-meizu-push-sdk
[![Build Status][travis-image]][travis] [![Coverage Status][coveralls-image]][coveralls]  [![License][license-image]][license]

## Installation
```
   go get github.com/comsince/go-meizu-push-sdk
```

## QuickStart

* 获取订阅开关状态

```go
func TestGetRegisterSwitch(t *testing.T) {
	message := GetRegisterSwitch(APP_ID,PUSH_ID,APP_KEY)
	fmt.Println("current message "+message.message)
	if(message.code != 200){
		t.Error("Status Code not 200")
	}
}

```

## ServerRegisterService

* API 列表

```
    func ChangeAllSwitch(appId string, pushId string, subSwitch bool, appKey string) PushResponse
    func ChangeRegisterSwitch(appId string, pushId string, msgType string, subSwitch bool, appKey string) PushResponse
    func GetRegisterSwitch(appId string, pushId string, appKey string) PushResponse
    func GetSubAlias(appId string, pushId string, appKey string) PushResponse
    func GetSubTags(appId string, pushId string, appKey string) PushResponse
    func Register(appId string, appKey string, deviceId string) PushResponse
    func ResolvePushResponse(res *httpclient.Response, err error) PushResponse
    func SubscribeAlias(appId string, pushId string, alias string, appKey string) PushResponse
    func SubscribeTags(appId string, pushId string, tags string, appKey string) PushResponse
    func UnSubAllTags(appId string, pushId string, appKey string) PushResponse
    func UnSubscribeAlias(appId string, pushId string, appKey string) PushResponse
    func UnSubscribeTags(appId string, pushId string, tags string, appKey string) PushResponse

```


## PushMessageService



**NOTE:**  详情请参考[meizu-push-godoc](https://godoc.org/github.com/comsince/go-meizu-push-sdk)


[travis]: https://travis-ci.org/comsince/go-meizu-push-sdk
[travis-image]: https://travis-ci.org/comsince/go-meizu-push-sdk.svg?branch=master

[license-image]: http://img.shields.io/badge/license-Apache--2-blue.svg?style=flat
[license]: http://www.apache.org/licenses/LICENSE-2.0

[coveralls-image]: https://coveralls.io/repos/github/comsince/go-meizu-push-sdk/badge.svg?branch=master
[coveralls]: https://coveralls.io/github/comsince/go-meizu-push-sdk?branch=master

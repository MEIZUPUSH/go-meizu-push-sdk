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




[travis]: https://travis-ci.org/comsince/go-meizu-push-sdk
[travis-image]: https://travis-ci.org/comsince/go-meizu-push-sdk.svg?branch=master

[license-image]: http://img.shields.io/badge/license-Apache--2-blue.svg?style=flat
[license]: http://www.apache.org/licenses/LICENSE-2.0

[coveralls-image]: https://coveralls.io/repos/github/comsince/go-meizu-push-sdk/badge.svg?branch=master
[coveralls]: https://coveralls.io/github/comsince/go-meizu-push-sdk?branch=master

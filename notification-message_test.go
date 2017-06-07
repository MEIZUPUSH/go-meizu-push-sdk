package goPushSdk

import (
	"fmt"
	"testing"
)

func TestBuildNotificationMessage(t *testing.T) {
	n := BuildNotificationMessage().
		noticeBarType(2).
		noticeTitle("标题").
		noticeContent("测试内容").toJson()

	fmt.Println(n)
}

package mzpush

import (
	"testing"
)

var client = NewClient("110029", "743cf8e3ebdf403db0881f675ab277a4")

var pushIds string = "Z496e0605725e6e45030106487d0d0242796700007450"

var msg1 *ThroughMessage = NewThroughMessage("透传消息:标题", "内容")
var msg2 *NotificationMessage = NewNotificationMessage("通知栏消息:标题", "内容")

func TestMzPush_ThroughSend(t *testing.T) {
	result, err := client.ThroughSend(msg1, pushIds)
	if err != nil {
		t.Errorf("TestMzPush_ThroughSend failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
}

func TestMzPush_NotificationSend(t *testing.T) {
	result, err := client.NotificationSend(msg2, pushIds)
	if err != nil {
		t.Errorf("TestMzPush_NotificationSend failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
}

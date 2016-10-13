package mzpush

import (
	"strconv"
	"testing"
)

var client = NewClient("App ID", "App Secret")

var pushIds string = "Z496e0605725e6e45030106487d0d0242796700007450"

var msg1 *ThroughMessage = NewThroughMessage("透传消息:标题", "内容")
var msg2 *NotificationMessage = NewNotificationMessage("通知栏消息:标题", "内容").
	SetClickTypeInfoClickType(2).
	SetClickTypeInfoUrl("http://www.sina.com.cn")
var msg3 *ThroughMessage = NewThroughMessage("通过taskId透传消息:标题", "内容")
var msg4 *NotificationMessage = NewNotificationMessage("通过taskId通知栏消息:标题", "内容")
var msg5 *ThroughMessage = NewThroughMessage("App消息:标题", "内容")
var msg6 *NotificationMessage = NewNotificationMessage("App通知栏消息:标题", "内容")

var ThroughtaskId, NotificationtaskId, ApptaskId float64

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

func TestMzPush_GetThroughTask(t *testing.T) {
	result, err := client.GetThroughTask(msg3)
	if err != nil {
		t.Errorf("TestMzPush_GetThroughTask failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
	ThroughtaskId = result.Value["taskId"].(float64)
}

func TestMzPush_GetNotificationTask(t *testing.T) {
	result, err := client.GetNotificationTask(msg4)
	if err != nil {
		t.Errorf("TestMzPush_GetNotificationTask failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
	NotificationtaskId = result.Value["taskId"].(float64)
}

func TestMzPush_TaskThroughSend(t *testing.T) {
	result, err := client.TaskSend(1, strconv.FormatInt(int64(ThroughtaskId), 10), pushIds)
	if err != nil {
		t.Errorf("TestMzPush_TaskThroughSend failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
}

func TestMzPush_TaskNotificationSend(t *testing.T) {
	result, err := client.TaskSend(0, strconv.FormatInt(int64(NotificationtaskId), 10), pushIds)
	if err != nil {
		t.Errorf("TestMzPush_TaskNotificationSend failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
	t.Logf("result.Value=%#v\n", result.Value)
}

func TestMzPush_AppThroughSend(t *testing.T) {
	result, err := client.AppSend(msg5, 1)
	if err != nil {
		t.Errorf("TestMzPush_AppThroughSend failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
	ApptaskId = result.Value["taskId"].(float64)
}

func TestMzPush_AppNotificationSend(t *testing.T) {
	result, err := client.AppSend(msg6, 0)
	if err != nil {
		t.Errorf("TestMzPush_AppNotificationSend failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
}

func TestMzPush_CancelTask(t *testing.T) {
	result, err := client.CancelTask(1, strconv.FormatInt(int64(ApptaskId), 10))
	if err != nil {
		t.Errorf("TestMzPush_CancelTask failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
}

package mzpush

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	//"fmt"
	"net/url"
	"sort"
	"strconv"
)

type MzPush struct {
	appId     string
	appSecret string
}

func NewClient(appId, appSecret string) *MzPush {
	return &MzPush{
		appId:     appId,
		appSecret: appSecret,
	}
}

// pushId推送接口（透传消息）
func (m *MzPush) ThroughSend(msg *ThroughMessage, pushIds string) (*SendResult, error) {
	params := url.Values{}
	params.Add("appId", m.appId)
	params.Add("pushIds", pushIds)
	msgStr, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	params.Add("messageJson", string(msgStr))
	sign := m.signStr(params)
	params.Add("sign", sign)
	bytes, err := doPost(Host+PushThroughSendURL, params)
	if err != nil {
		return nil, err
	}
	var result SendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// pushId推送接口（通知栏消息）
func (m *MzPush) NotificationSend(msg *NotificationMessage, pushIds string) (*SendResult, error) {
	params := url.Values{}
	params.Add("appId", m.appId)
	params.Add("pushIds", pushIds)
	msgStr, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	params.Add("messageJson", string(msgStr))
	sign := m.signStr(params)
	params.Add("sign", sign)
	bytes, err := doPost(Host+PushNotificationSendURL, params)
	if err != nil {
		return nil, err
	}
	var result SendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取推送taskId 透传
func (m *MzPush) GetThroughTask(msg *ThroughMessage) (*SendResult, error) {
	params := url.Values{}
	params.Add("appId", m.appId)
	params.Add("pushType", "1")
	msgStr, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	params.Add("messageJson", string(msgStr))
	sign := m.signStr(params)
	params.Add("sign", sign)
	bytes, err := doPost(Host+GetTaskURL, params)
	if err != nil {
		return nil, err
	}
	var result SendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取推送taskId 通知栏
func (m *MzPush) GetNotificationTask(msg *NotificationMessage) (*SendResult, error) {
	params := url.Values{}
	params.Add("appId", m.appId)
	params.Add("pushType", "0")
	msgStr, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	params.Add("messageJson", string(msgStr))
	sign := m.signStr(params)
	params.Add("sign", sign)
	bytes, err := doPost(Host+GetTaskURL, params)
	if err != nil {
		return nil, err
	}
	var result SendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// taskId推送接口 透传or通知栏
func (m *MzPush) TaskSend(pushType int, taskId, pushIds string) (*SendResult, error) {
	var TaskSendURL string
	params := url.Values{}
	params.Add("taskId", taskId)
	params.Add("appId", m.appId)
	params.Add("pushIds", pushIds)
	sign := m.signStr(params)
	params.Add("sign", sign)
	if pushType == 1 {
		TaskSendURL = TaskThroughSendURL
	} else {
		TaskSendURL = TaskNotificationSendURL
	}
	bytes, err := doPost(Host+TaskSendURL, params)
	if err != nil {
		return nil, err
	}
	var result SendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 全部用户推送
func (m *MzPush) AppSend(msg interface{}, pushType int) (*SendResult, error) {
	params := url.Values{}
	params.Add("appId", m.appId)
	params.Add("pushType", strconv.Itoa(pushType))
	msgStr, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	params.Add("messageJson", string(msgStr))
	sign := m.signStr(params)
	params.Add("sign", sign)
	bytes, err := doPost(Host+AppSendURL, params)
	if err != nil {
		return nil, err
	}
	var result SendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 取消任务推送
func (m *MzPush) CancelTask(pushType int, taskId string) (*SendResult, error) {
	params := url.Values{}
	params.Add("appId", m.appId)
	params.Add("pushType", strconv.Itoa(pushType))
	params.Add("taskId", taskId)
	sign := m.signStr(params)
	params.Add("sign", sign)
	bytes, err := doPost(Host+TaskCancelURL, params)
	if err != nil {
		return nil, err
	}
	var result SendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 签名
func (m *MzPush) signStr(params url.Values) string {
	var keys []string
	var parameterStr string

	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, value := range keys {
		parameterStr += value + "=" + params[value][0]
	}
	parameterStr += m.appSecret

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(parameterStr))
	signStr := hex.EncodeToString(md5Ctx.Sum(nil))

	return signStr
}

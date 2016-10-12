package mzpush

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	//"fmt"
	"net/url"
	"sort"
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

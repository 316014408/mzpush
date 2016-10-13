package mzpush

type ThroughMessage struct {
	Title        string       `json:"title"`   // 推送标题, 【string 非必填，字数显示1~32个】
	Content      string       `json:"content"` // 推送内容,  【string 必填，字数限制2000以内】
	PushTimeInfo PushTimeInfo `json:"pushTimeInfo"`
}

type NotificationMessage struct {
	NoticeBarInfo    NoticeBarInfo    `json:"noticeBarInfo"`
	NoticeExpandInfo NoticeExpandInfo `json:"noticeExpandInfo"`
	ClickTypeInfo    ClickTypeInfo    `json:"clickTypeInfo"`
	PushTimeInfo     PushTimeInfo     `json:"pushTimeInfo"`
	AdvanceInfo      AdvanceInfo      `json:"advanceInfo"`
}

type NoticeBarInfo struct {
	NoticeBarType int32  `json:"noticeBarType,omitempty"` // 通知栏样式(0, "标准")【int 非必填，值为0】
	Title         string `json:"title"`                   // 推送标题, 【string 必填，字数限制1~32】
	Content       string `json:"content"`                 // 推送内容, 【string 必填，字数限制1~100】
}

type NoticeExpandInfo struct {
	NoticeExpandType    int32  `json:"noticeExpandType,omitempty"` // 展开方式 (0, "标准"),(1, "文本")【int 非必填，值为0、1】
	NoticeExpandContent string `json:"noticeExpandContent"`        // 展开内容, 【string noticeExpandType为文本时，必填】
}

type ClickTypeInfo struct {
	ClickType  int32  `json:"clickType,omitempty"`  // 展开方式 (0, "标准"),(1, "文本")【int 非必填，值为0、1】
	Url        string `json:"url"`                  // URI页面地址, 【string clickType为打开URI页面时，必填, 长度限制1000】
	Parameters string `json:"parameters,omitempty"` // 参数 【JSON格式】【非必填】
	Activity   string `json:"activity"`             // 应用页面地址 【string clickType为打开应用页面时，必填, 长度限制1000】
}

type PushTimeInfo struct {
	OffLine      int32  `json:"offLine,omitempty"`      // 是否进离线消息(0 否 1 是[validTime]) 【int 非必填，默认值为1】
	ValidTime    int32  `json:"validTime"`              // 有效时长 (1 72 小时内的正整数) 【int offLine值为1时，必填，默认24】
	PushTimeType int32  `json:"pushTimeType,omitempty"` // 定时推送 (0, "即时"),(1, "定时")【必填，默认0】
	StartTime    string `json:"startTime,omitempty"`    // 任务定时开始时间(yyyy-MM-dd HH:mm:ss) 【非必填pushTimeType为1必填】
}

type AdvanceInfo struct {
	Suspend          int32            `json:"suspend,omitempty"`        // 是否通知栏悬浮窗显示 (1 显示  0 不显示) 【int 非必填，默认1】
	ClearNoticeBar   int32            `json:"clearNoticeBar,omitempty"` // 是否可清除通知栏 (1 可以  0 不可以) 【int 非必填，默认1】
	NotificationType NotificationType `json:"notificationType"`
}

type NotificationType struct {
	Vibrate int32 `json:"vibrate,omitempty"` // 震动 (0关闭  1 开启), 【int 非必填，默认1】
	Lights  int32 `json:"lights,omitempty"`  // 闪光 (0关闭  1 开启), 【int 非必填，默认1】
	Sound   int32 `json:"sound,omitempty"`   // 声音 (0关闭  1 开启), 【int 非必填，默认1】
}

func NewThroughMessage(notificationTitle, notificationContent string) *ThroughMessage {
	return &ThroughMessage{
		Title:   notificationTitle,
		Content: notificationContent,
		PushTimeInfo: PushTimeInfo{
			ValidTime: 24,
		},
	}
}

func NewNotificationMessage(notificationTitle, notificationContent string) *NotificationMessage {
	return &NotificationMessage{
		NoticeBarInfo: NoticeBarInfo{
			Title:   notificationTitle,
			Content: notificationContent,
		},
		PushTimeInfo: PushTimeInfo{
			ValidTime: 24,
		},
	}
}

func (t *ThroughMessage) SetThroughPushTimeInfoOffLine(offLine int32) *ThroughMessage {
	t.PushTimeInfo.OffLine = offLine
	return t
}

func (t *ThroughMessage) SetThroughPushTimeInfoValidTime(validTime int32) *ThroughMessage {
	t.PushTimeInfo.ValidTime = validTime
	return t
}

func (t *ThroughMessage) SetThroughPushTimeInfoPushTimeType(pushTimeType int32) *ThroughMessage {
	t.PushTimeInfo.PushTimeType = pushTimeType
	return t
}

func (t *ThroughMessage) SetThroughPushTimeInfoStartTime(startTime string) *ThroughMessage {
	t.PushTimeInfo.StartTime = startTime
	return t
}

func (n *NotificationMessage) SetNoticeBarInfoNoticeBarType(noticeBarType int32) *NotificationMessage {
	n.NoticeBarInfo.NoticeBarType = noticeBarType
	return n
}

func (n *NotificationMessage) SetNoticeBarInfoTitle(title string) *NotificationMessage {
	n.NoticeBarInfo.Title = title
	return n
}

func (n *NotificationMessage) SetNoticeBarInfoContent(content string) *NotificationMessage {
	n.NoticeBarInfo.Content = content
	return n
}

func (n *NotificationMessage) SetNoticeExpandInfoNoticeExpandType(noticeExpandType int32) *NotificationMessage {
	n.NoticeExpandInfo.NoticeExpandType = noticeExpandType
	return n
}

func (n *NotificationMessage) SetNoticeExpandInfoNoticeExpandContent(noticeExpandContent string) *NotificationMessage {
	n.NoticeExpandInfo.NoticeExpandContent = noticeExpandContent
	return n
}

func (n *NotificationMessage) SetClickTypeInfoClickType(clickType int32) *NotificationMessage {
	n.ClickTypeInfo.ClickType = clickType
	return n
}

func (n *NotificationMessage) SetClickTypeInfoUrl(url string) *NotificationMessage {
	n.ClickTypeInfo.Url = url
	return n
}

func (n *NotificationMessage) SetClickTypeInfoParameters(parameters string) *NotificationMessage {
	n.ClickTypeInfo.Parameters = parameters
	return n
}

func (n *NotificationMessage) SetClickTypeInfoActivity(activity string) *NotificationMessage {
	n.ClickTypeInfo.Activity = activity
	return n
}

func (n *NotificationMessage) SetPushTimeInfoOffLine(offLine int32) *NotificationMessage {
	n.PushTimeInfo.OffLine = offLine
	return n
}

func (n *NotificationMessage) SetPushTimeInfoValidTime(validTime int32) *NotificationMessage {
	n.PushTimeInfo.ValidTime = validTime
	return n
}

func (n *NotificationMessage) SetPushTimeInfoPushTimeType(pushTimeType int32) *NotificationMessage {
	n.PushTimeInfo.PushTimeType = pushTimeType
	return n
}

func (n *NotificationMessage) SetPushTimeInfoPushStartTime(startTime string) *NotificationMessage {
	n.PushTimeInfo.StartTime = startTime
	return n
}

func (n *NotificationMessage) SetAdvanceInfoSuspend(suspend int32) *NotificationMessage {
	n.AdvanceInfo.Suspend = suspend
	return n
}

func (n *NotificationMessage) SetAdvanceInfoClearNoticeBar(clearNoticeBar int32) *NotificationMessage {
	n.AdvanceInfo.ClearNoticeBar = clearNoticeBar
	return n
}

func (n *NotificationMessage) SetAdvanceInfoNotificationTypeVibrate(vibrate int32) *NotificationMessage {
	n.AdvanceInfo.NotificationType.Vibrate = vibrate
	return n
}

func (n *NotificationMessage) SetAdvanceInfoNotificationTypeLights(lights int32) *NotificationMessage {
	n.AdvanceInfo.NotificationType.Lights = lights
	return n
}

func (n *NotificationMessage) SetAdvanceInfoNotificationTypeSound(sound int32) *NotificationMessage {
	n.AdvanceInfo.NotificationType.Sound = sound
	return n
}

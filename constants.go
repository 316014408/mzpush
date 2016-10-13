package mzpush

// 文档地址 http://open-wiki.flyme.cn/index.php?title=%E9%AD%85%E6%97%8F%E6%8E%A8%E9%80%81%E5%B9%B3%E5%8F%B0Server%E7%AB%AFAPI%E6%96%87%E6%A1%A3
const (
	Host = "https://api-push.meizu.com"
)

const (
	PushThroughSendURL      = "/garcia/api/server/push/unvarnished/pushByPushId "     // pushId推送接口（透传消息）
	PushNotificationSendURL = "/garcia/api/server/push/varnished/pushByPushId"        // pushId推送接口（通知栏消息）
	GetTaskURL              = "/garcia/api/server/push/pushTask/getTaskId"            // 获取推送taskId
	TaskThroughSendURL      = "/garcia/api/server/push/task/unvarnished/pushByPushId" // taskId推送接口（透传消息）
	TaskNotificationSendURL = "/garcia/api/server/push/task/varnished/pushByPushId"   // taskId推送接口（通知栏消息）
	AppSendURL              = "/garcia/api/server/push/pushTask/pushToApp"            // appId推送接口
	TaskCancelURL           = "/garcia/api/server/push/pushTask/cancel"               // 推送取消接口（只针对taskId推送接口和appId推送接口?）
)

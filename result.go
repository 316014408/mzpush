package mzpush

type SendResult struct {
	Code     string      `json:"code"`     // 必选,返回码
	Message  string      `json:"message"`  // 可选，返回消息，网页端接口出现错误时使用此消息展示给用户，手机端可忽略此消息，甚至服务端不传输此消息
	Value    interface{} `json:"value"`    // 必选，返回结果
	Redirect string      `json:"redirect"` // 可选, returnCode=300 重定向时，使用此URL重新请求
}

// Package i18n Code generated by go generate; DO NOT EDIT. Please adjust your TOML file to regenerate.
package i18n

import (
	i18ncore "github.com/eliassama/go-i18n/core"
	"golang.org/x/text/language"
	"strings"
)

var Msg = struct {
	AcceptedMsg                      *i18ncore.Bundle
	AlreadyReportedMsg               *i18ncore.Bundle
	BadGatewayMsg                    *i18ncore.Bundle
	BadRequestMsg                    *i18ncore.Bundle
	ConflictMsg                      *i18ncore.Bundle
	ContinueMsg                      *i18ncore.Bundle
	CreatedMsg                       *i18ncore.Bundle
	EarlyHintsMsg                    *i18ncore.Bundle
	ExpectationFailedMsg             *i18ncore.Bundle
	FailedDependencyMsg              *i18ncore.Bundle
	ForbiddenMsg                     *i18ncore.Bundle
	FoundMsg                         *i18ncore.Bundle
	GatewayTimeoutMsg                *i18ncore.Bundle
	GoneMsg                          *i18ncore.Bundle
	HTTPVersionNotSupportedMsg       *i18ncore.Bundle
	IMUsedMsg                        *i18ncore.Bundle
	InsufficientStorageMsg           *i18ncore.Bundle
	InternalServerErrorMsg           *i18ncore.Bundle
	LengthRequiredMsg                *i18ncore.Bundle
	LockedMsg                        *i18ncore.Bundle
	LoopDetectedMsg                  *i18ncore.Bundle
	MethodNotAllowedMsg              *i18ncore.Bundle
	MisdirectedRequestMsg            *i18ncore.Bundle
	MovedPermanentlyMsg              *i18ncore.Bundle
	MultiStatusMsg                   *i18ncore.Bundle
	MultipleChoicesMsg               *i18ncore.Bundle
	NetworkAuthenticationRequiredMsg *i18ncore.Bundle
	NoContentMsg                     *i18ncore.Bundle
	NonAuthoritativeInfoMsg          *i18ncore.Bundle
	NotAcceptableMsg                 *i18ncore.Bundle
	NotExtendedMsg                   *i18ncore.Bundle
	NotFoundMsg                      *i18ncore.Bundle
	NotImplementedMsg                *i18ncore.Bundle
	NotModifiedMsg                   *i18ncore.Bundle
	OKMsg                            *i18ncore.Bundle
	PartialContentMsg                *i18ncore.Bundle
	PaymentRequiredMsg               *i18ncore.Bundle
	PermanentRedirectMsg             *i18ncore.Bundle
	PreconditionFailedMsg            *i18ncore.Bundle
	PreconditionRequiredMsg          *i18ncore.Bundle
	ProcessingMsg                    *i18ncore.Bundle
	ProxyAuthRequiredMsg             *i18ncore.Bundle
	RequestEntityTooLargeMsg         *i18ncore.Bundle
	RequestHeaderFieldsTooLargeMsg   *i18ncore.Bundle
	RequestTimeoutMsg                *i18ncore.Bundle
	RequestURITooLongMsg             *i18ncore.Bundle
	RequestedRangeNotSatisfiableMsg  *i18ncore.Bundle
	ResetContentMsg                  *i18ncore.Bundle
	SeeOtherMsg                      *i18ncore.Bundle
	ServiceUnavailableMsg            *i18ncore.Bundle
	SwitchingProtocolsMsg            *i18ncore.Bundle
	TeapotMsg                        *i18ncore.Bundle
	TemporaryRedirectMsg             *i18ncore.Bundle
	TooEarlyMsg                      *i18ncore.Bundle
	TooManyRequestsMsg               *i18ncore.Bundle
	UnauthorizedMsg                  *i18ncore.Bundle
	UnavailableForLegalReasonsMsg    *i18ncore.Bundle
	UnprocessableEntityMsg           *i18ncore.Bundle
	UnsupportedMediaTypeMsg          *i18ncore.Bundle
	UpgradeRequiredMsg               *i18ncore.Bundle
	UseProxyMsg                      *i18ncore.Bundle
	VariantAlsoNegotiatesMsg         *i18ncore.Bundle
}{
	AcceptedMsg: &i18ncore.Bundle{
		Status:   202,
		Code:     100202,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The task has been created and is being queued for execution",
			"zh-cn": "任务已创建，正在排队执行中",
		},
	},
	AlreadyReportedMsg: &i18ncore.Bundle{
		Status:   208,
		Code:     100208,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The result has already been reported",
			"zh-cn": "已经报告的结果",
		},
	},
	BadGatewayMsg: &i18ncore.Bundle{
		Status:   502,
		Code:     100502,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Bad gateway",
			"zh-cn": "网关错误",
		},
	},
	BadRequestMsg: &i18ncore.Bundle{
		Status:   400,
		Code:     100400,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Bad request",
			"zh-cn": "错误的请求",
		},
	},
	ConflictMsg: &i18ncore.Bundle{
		Status:   409,
		Code:     100409,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Request conflict",
			"zh-cn": "请求冲突",
		},
	},
	ContinueMsg: &i18ncore.Bundle{
		Status:   100,
		Code:     100100,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The requester should continue with the request",
			"zh-cn": "请求者应继续进行请求",
		},
	},
	CreatedMsg: &i18ncore.Bundle{
		Status:   201,
		Code:     100201,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The operation was completed",
			"zh-cn": "操作已完成",
		},
	},
	EarlyHintsMsg: &i18ncore.Bundle{
		Status:   103,
		Code:     100103,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Early hints",
			"zh-cn": "提前提示",
		},
	},
	ExpectationFailedMsg: &i18ncore.Bundle{
		Status:   417,
		Code:     100417,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Expectation failed",
			"zh-cn": "期望值失败",
		},
	},
	FailedDependencyMsg: &i18ncore.Bundle{
		Status:   424,
		Code:     100424,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Failed dependency",
			"zh-cn": "依赖关系失败",
		},
	},
	ForbiddenMsg: &i18ncore.Bundle{
		Status:   403,
		Code:     100403,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Forbidden",
			"zh-cn": "禁止访问",
		},
	},
	FoundMsg: &i18ncore.Bundle{
		Status:   302,
		Code:     100302,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The resource was found",
			"zh-cn": "找到资源",
		},
	},
	GatewayTimeoutMsg: &i18ncore.Bundle{
		Status:   504,
		Code:     100504,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Gateway timeout",
			"zh-cn": "网关超时",
		},
	},
	GoneMsg: &i18ncore.Bundle{
		Status:   410,
		Code:     100410,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Resource gone",
			"zh-cn": "资源已删除",
		},
	},
	HTTPVersionNotSupportedMsg: &i18ncore.Bundle{
		Status:   505,
		Code:     100505,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "HTTP version not supported",
			"zh-cn": "不支持的HTTP版本",
		},
	},
	IMUsedMsg: &i18ncore.Bundle{
		Status:   226,
		Code:     100226,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "IM used",
			"zh-cn": "IM已使用",
		},
	},
	InsufficientStorageMsg: &i18ncore.Bundle{
		Status:   507,
		Code:     100507,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Insufficient storage",
			"zh-cn": "存储不足",
		},
	},
	InternalServerErrorMsg: &i18ncore.Bundle{
		Status:   500,
		Code:     100500,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Internal server error",
			"zh-cn": "服务器内部错误",
		},
	},
	LengthRequiredMsg: &i18ncore.Bundle{
		Status:   411,
		Code:     100411,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Length required",
			"zh-cn": "需要内容长度",
		},
	},
	LockedMsg: &i18ncore.Bundle{
		Status:   423,
		Code:     100423,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Resource locked",
			"zh-cn": "资源被锁定",
		},
	},
	LoopDetectedMsg: &i18ncore.Bundle{
		Status:   508,
		Code:     100508,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Loop detected",
			"zh-cn": "检测到循环",
		},
	},
	MethodNotAllowedMsg: &i18ncore.Bundle{
		Status:   405,
		Code:     100405,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Method not allowed",
			"zh-cn": "方法不允许",
		},
	},
	MisdirectedRequestMsg: &i18ncore.Bundle{
		Status:   421,
		Code:     100421,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Misdirected request",
			"zh-cn": "错误定向的请求",
		},
	},
	MovedPermanentlyMsg: &i18ncore.Bundle{
		Status:   301,
		Code:     100301,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The resource has been permanently moved",
			"zh-cn": "资源已永久移动",
		},
	},
	MultiStatusMsg: &i18ncore.Bundle{
		Status:   207,
		Code:     100207,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Multi-status response",
			"zh-cn": "多状态响应",
		},
	},
	MultipleChoicesMsg: &i18ncore.Bundle{
		Status:   300,
		Code:     100300,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Multiple choices",
			"zh-cn": "多种选择",
		},
	},
	NetworkAuthenticationRequiredMsg: &i18ncore.Bundle{
		Status:   511,
		Code:     100511,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Network authentication required",
			"zh-cn": "需要网络身份验证",
		},
	},
	NoContentMsg: &i18ncore.Bundle{
		Status:   204,
		Code:     100204,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The operation was successful but there is no content to return",
			"zh-cn": "操作已成功但无返回内容",
		},
	},
	NonAuthoritativeInfoMsg: &i18ncore.Bundle{
		Status:   203,
		Code:     100203,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The request was successful, but the information comes from non-authoritative sources",
			"zh-cn": "请求已成功，但信息来自非权威信息",
		},
	},
	NotAcceptableMsg: &i18ncore.Bundle{
		Status:   406,
		Code:     100406,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Request not acceptable",
			"zh-cn": "请求不可接受",
		},
	},
	NotExtendedMsg: &i18ncore.Bundle{
		Status:   510,
		Code:     100510,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Not extended",
			"zh-cn": "政策不延展",
		},
	},
	NotFoundMsg: &i18ncore.Bundle{
		Status:   404,
		Code:     100404,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Resource not found",
			"zh-cn": "未找到资源",
		},
	},
	NotImplementedMsg: &i18ncore.Bundle{
		Status:   501,
		Code:     100501,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Not implemented",
			"zh-cn": "未实现",
		},
	},
	NotModifiedMsg: &i18ncore.Bundle{
		Status:   304,
		Code:     100304,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The resource has not been modified",
			"zh-cn": "资源未修改",
		},
	},
	OKMsg: &i18ncore.Bundle{
		Status:   200,
		Code:     100200,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Successful operation",
			"zh-cn": "操作成功",
		},
	},
	PartialContentMsg: &i18ncore.Bundle{
		Status:   206,
		Code:     100206,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Partial content has been returned",
			"zh-cn": "部分内容已返回",
		},
	},
	PaymentRequiredMsg: &i18ncore.Bundle{
		Status:   402,
		Code:     100402,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Payment required",
			"zh-cn": "需要付款",
		},
	},
	PermanentRedirectMsg: &i18ncore.Bundle{
		Status:   308,
		Code:     100308,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Permanent redirect",
			"zh-cn": "永久重定向",
		},
	},
	PreconditionFailedMsg: &i18ncore.Bundle{
		Status:   412,
		Code:     100412,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Precondition failed",
			"zh-cn": "前提条件失败",
		},
	},
	PreconditionRequiredMsg: &i18ncore.Bundle{
		Status:   428,
		Code:     100428,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Precondition required",
			"zh-cn": "需要前提条件",
		},
	},
	ProcessingMsg: &i18ncore.Bundle{
		Status:   102,
		Code:     100102,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The request has been accepted, but not yet processed",
			"zh-cn": "请求已被接受，但还未处理",
		},
	},
	ProxyAuthRequiredMsg: &i18ncore.Bundle{
		Status:   407,
		Code:     100407,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Proxy authentication required",
			"zh-cn": "需要代理身份验证",
		},
	},
	RequestEntityTooLargeMsg: &i18ncore.Bundle{
		Status:   413,
		Code:     100413,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Request entity too large",
			"zh-cn": "请求实体过大",
		},
	},
	RequestHeaderFieldsTooLargeMsg: &i18ncore.Bundle{
		Status:   431,
		Code:     100431,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Request header fields too large",
			"zh-cn": "请求头字段过大",
		},
	},
	RequestTimeoutMsg: &i18ncore.Bundle{
		Status:   408,
		Code:     100408,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Request timeout",
			"zh-cn": "请求超时",
		},
	},
	RequestURITooLongMsg: &i18ncore.Bundle{
		Status:   414,
		Code:     100414,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Request URI too long",
			"zh-cn": "请求URI过长",
		},
	},
	RequestedRangeNotSatisfiableMsg: &i18ncore.Bundle{
		Status:   416,
		Code:     100416,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Requested range not satisfiable",
			"zh-cn": "请求范围不满足",
		},
	},
	ResetContentMsg: &i18ncore.Bundle{
		Status:   205,
		Code:     100205,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The operation was successful, please reset the view",
			"zh-cn": "操作已成功，请重置视图",
		},
	},
	SeeOtherMsg: &i18ncore.Bundle{
		Status:   303,
		Code:     100303,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Please see other address",
			"zh-cn": "请参见其他地址",
		},
	},
	ServiceUnavailableMsg: &i18ncore.Bundle{
		Status:   503,
		Code:     100503,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Service unavailable",
			"zh-cn": "服务不可用",
		},
	},
	SwitchingProtocolsMsg: &i18ncore.Bundle{
		Status:   101,
		Code:     100101,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "The server is switching protocols",
			"zh-cn": "服务器切换协议",
		},
	},
	TeapotMsg: &i18ncore.Bundle{
		Status:   418,
		Code:     100418,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "I'm a teapot",
			"zh-cn": "我是一个茶壶",
		},
	},
	TemporaryRedirectMsg: &i18ncore.Bundle{
		Status:   307,
		Code:     100307,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Temporary redirect",
			"zh-cn": "临时重定向",
		},
	},
	TooEarlyMsg: &i18ncore.Bundle{
		Status:   425,
		Code:     100425,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Request too early",
			"zh-cn": "请求过早",
		},
	},
	TooManyRequestsMsg: &i18ncore.Bundle{
		Status:   429,
		Code:     100429,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Too many requests",
			"zh-cn": "请求过多",
		},
	},
	UnauthorizedMsg: &i18ncore.Bundle{
		Status:   401,
		Code:     100401,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Unauthorized access",
			"zh-cn": "未授权访问",
		},
	},
	UnavailableForLegalReasonsMsg: &i18ncore.Bundle{
		Status:   451,
		Code:     100451,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Unavailable for legal reasons",
			"zh-cn": "因法律原因不可用",
		},
	},
	UnprocessableEntityMsg: &i18ncore.Bundle{
		Status:   422,
		Code:     100422,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Unprocessable entity",
			"zh-cn": "无法处理的实体",
		},
	},
	UnsupportedMediaTypeMsg: &i18ncore.Bundle{
		Status:   415,
		Code:     100415,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Unsupported media type",
			"zh-cn": "不支持的媒体类型",
		},
	},
	UpgradeRequiredMsg: &i18ncore.Bundle{
		Status:   426,
		Code:     100426,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Upgrade required",
			"zh-cn": "需要升级",
		},
	},
	UseProxyMsg: &i18ncore.Bundle{
		Status:   305,
		Code:     100305,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Please access the resource through a proxy",
			"zh-cn": "请通过代理访问资源",
		},
	},
	VariantAlsoNegotiatesMsg: &i18ncore.Bundle{
		Status:   506,
		Code:     100506,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Variant also negotiates",
			"zh-cn": "变体也在协商",
		},
	},
}

func SendAcceptedMsg(langCode string) (int, int, string) {
	bundle := Msg.AcceptedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendAlreadyReportedMsg(langCode string) (int, int, string) {
	bundle := Msg.AlreadyReportedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendBadGatewayMsg(langCode string) (int, int, string) {
	bundle := Msg.BadGatewayMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendBadRequestMsg(langCode string) (int, int, string) {
	bundle := Msg.BadRequestMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendConflictMsg(langCode string) (int, int, string) {
	bundle := Msg.ConflictMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendContinueMsg(langCode string) (int, int, string) {
	bundle := Msg.ContinueMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendCreatedMsg(langCode string) (int, int, string) {
	bundle := Msg.CreatedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendEarlyHintsMsg(langCode string) (int, int, string) {
	bundle := Msg.EarlyHintsMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendExpectationFailedMsg(langCode string) (int, int, string) {
	bundle := Msg.ExpectationFailedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendFailedDependencyMsg(langCode string) (int, int, string) {
	bundle := Msg.FailedDependencyMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendForbiddenMsg(langCode string) (int, int, string) {
	bundle := Msg.ForbiddenMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendFoundMsg(langCode string) (int, int, string) {
	bundle := Msg.FoundMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendGatewayTimeoutMsg(langCode string) (int, int, string) {
	bundle := Msg.GatewayTimeoutMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendGoneMsg(langCode string) (int, int, string) {
	bundle := Msg.GoneMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendHTTPVersionNotSupportedMsg(langCode string) (int, int, string) {
	bundle := Msg.HTTPVersionNotSupportedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendIMUsedMsg(langCode string) (int, int, string) {
	bundle := Msg.IMUsedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendInsufficientStorageMsg(langCode string) (int, int, string) {
	bundle := Msg.InsufficientStorageMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendInternalServerErrorMsg(langCode string) (int, int, string) {
	bundle := Msg.InternalServerErrorMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendLengthRequiredMsg(langCode string) (int, int, string) {
	bundle := Msg.LengthRequiredMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendLockedMsg(langCode string) (int, int, string) {
	bundle := Msg.LockedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendLoopDetectedMsg(langCode string) (int, int, string) {
	bundle := Msg.LoopDetectedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendMethodNotAllowedMsg(langCode string) (int, int, string) {
	bundle := Msg.MethodNotAllowedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendMisdirectedRequestMsg(langCode string) (int, int, string) {
	bundle := Msg.MisdirectedRequestMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendMovedPermanentlyMsg(langCode string) (int, int, string) {
	bundle := Msg.MovedPermanentlyMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendMultiStatusMsg(langCode string) (int, int, string) {
	bundle := Msg.MultiStatusMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendMultipleChoicesMsg(langCode string) (int, int, string) {
	bundle := Msg.MultipleChoicesMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendNetworkAuthenticationRequiredMsg(langCode string) (int, int, string) {
	bundle := Msg.NetworkAuthenticationRequiredMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendNoContentMsg(langCode string) (int, int, string) {
	bundle := Msg.NoContentMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendNonAuthoritativeInfoMsg(langCode string) (int, int, string) {
	bundle := Msg.NonAuthoritativeInfoMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendNotAcceptableMsg(langCode string) (int, int, string) {
	bundle := Msg.NotAcceptableMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendNotExtendedMsg(langCode string) (int, int, string) {
	bundle := Msg.NotExtendedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendNotFoundMsg(langCode string) (int, int, string) {
	bundle := Msg.NotFoundMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendNotImplementedMsg(langCode string) (int, int, string) {
	bundle := Msg.NotImplementedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendNotModifiedMsg(langCode string) (int, int, string) {
	bundle := Msg.NotModifiedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendOKMsg(langCode string) (int, int, string) {
	bundle := Msg.OKMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendPartialContentMsg(langCode string) (int, int, string) {
	bundle := Msg.PartialContentMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendPaymentRequiredMsg(langCode string) (int, int, string) {
	bundle := Msg.PaymentRequiredMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendPermanentRedirectMsg(langCode string) (int, int, string) {
	bundle := Msg.PermanentRedirectMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendPreconditionFailedMsg(langCode string) (int, int, string) {
	bundle := Msg.PreconditionFailedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendPreconditionRequiredMsg(langCode string) (int, int, string) {
	bundle := Msg.PreconditionRequiredMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendProcessingMsg(langCode string) (int, int, string) {
	bundle := Msg.ProcessingMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendProxyAuthRequiredMsg(langCode string) (int, int, string) {
	bundle := Msg.ProxyAuthRequiredMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendRequestEntityTooLargeMsg(langCode string) (int, int, string) {
	bundle := Msg.RequestEntityTooLargeMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendRequestHeaderFieldsTooLargeMsg(langCode string) (int, int, string) {
	bundle := Msg.RequestHeaderFieldsTooLargeMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendRequestTimeoutMsg(langCode string) (int, int, string) {
	bundle := Msg.RequestTimeoutMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendRequestURITooLongMsg(langCode string) (int, int, string) {
	bundle := Msg.RequestURITooLongMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendRequestedRangeNotSatisfiableMsg(langCode string) (int, int, string) {
	bundle := Msg.RequestedRangeNotSatisfiableMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendResetContentMsg(langCode string) (int, int, string) {
	bundle := Msg.ResetContentMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendSeeOtherMsg(langCode string) (int, int, string) {
	bundle := Msg.SeeOtherMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendServiceUnavailableMsg(langCode string) (int, int, string) {
	bundle := Msg.ServiceUnavailableMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendSwitchingProtocolsMsg(langCode string) (int, int, string) {
	bundle := Msg.SwitchingProtocolsMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendTeapotMsg(langCode string) (int, int, string) {
	bundle := Msg.TeapotMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendTemporaryRedirectMsg(langCode string) (int, int, string) {
	bundle := Msg.TemporaryRedirectMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendTooEarlyMsg(langCode string) (int, int, string) {
	bundle := Msg.TooEarlyMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendTooManyRequestsMsg(langCode string) (int, int, string) {
	bundle := Msg.TooManyRequestsMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendUnauthorizedMsg(langCode string) (int, int, string) {
	bundle := Msg.UnauthorizedMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendUnavailableForLegalReasonsMsg(langCode string) (int, int, string) {
	bundle := Msg.UnavailableForLegalReasonsMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendUnprocessableEntityMsg(langCode string) (int, int, string) {
	bundle := Msg.UnprocessableEntityMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendUnsupportedMediaTypeMsg(langCode string) (int, int, string) {
	bundle := Msg.UnsupportedMediaTypeMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendUpgradeRequiredMsg(langCode string) (int, int, string) {
	bundle := Msg.UpgradeRequiredMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendUseProxyMsg(langCode string) (int, int, string) {
	bundle := Msg.UseProxyMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}
func SendVariantAlsoNegotiatesMsg(langCode string) (int, int, string) {
	bundle := Msg.VariantAlsoNegotiatesMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}

func convertLangCode(langCode string) string {
	if langTag, err := language.Parse(langCode); err != nil {
		return ""
	} else {
		return strings.ToLower(langTag.String())
	}
}
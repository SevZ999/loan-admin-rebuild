package vo

// Code 业务状态码
const (
	// 请求参数无效
	INVALID_REQUEST_PARAMETERS = "INVALID_REQUEST_PARAMETERS"
	// 操作成功
	SUCCESS = "SUCCESS"
	// 操作失败
	FAILED = "FAILED"
	// 创建操作状态
	CREATION_FAILED  = "CREATION_FAILED"
	CREATION_SUCCESS = "CREATION_SUCCESS"

	// 更新操作状态
	UPDATE_FAILED  = "UPDATE_FAILED"
	UPDATE_SUCCESS = "UPDATE_SUCCESS"

	// 删除操作状态
	DELETE_FAILED  = "DELETE_FAILED"
	DELETE_SUCCESS = "DELETE_SUCCESS"

	// 登录相关
	LOGIN_ERROR       = "LOGIN_ERROR"       // 登录失败
	LOGIN_SUCCESS     = "LOGIN_SUCCESS"     // 登录成功
	CODE_ERROR        = "CODE_ERROR"        // 验证码错误
	USER_DISABLED     = "USER_DISABLED"     // 用户被禁用
	SIGNATURE_MISSING = "SIGNATURE_MISSING" //缺少签名信息
	INVALID_TIMESTAMP = "INVALID_TIMESTAMP" //无效的时间戳
	TIMESTAMP_ERROR   = "TIMESTAMP_ERROR"   //时间戳错误
	ERR_READ_BODY     = "ERR_READ_BODY"     //读取请求体失败
	INVALID_JSON      = "INVALID_JSON"      // 无效的JSON
	INVALID_SIGNATURE = "INVALID_SIGNATURE" // 无效的签名

	CAPTCHA_ERROR = "CAPTCHA_ERROR" // 验证码错误
	// 认证相关
	AUTH_ERROR        = "AUTH_ERROR"        // 认证失败
	AUTH_DEVICE_ERROR = "AUTH_DEVICE_ERROR" // 设备不存在
	AUTH_USER_ERROR   = "AUTH_USER_ERROR"   // 用户不存在
	PASSWORD_ERROR    = "PASSWORD_ERROR"    // 密码错误
	USERNAME_EXISTS   = "USERNAME_EXISTS"   // 用户名已存在

	// 短信相关
	SMS_SEND_ERROR   = "SMS_SEND_ERROR"   // 短信发送失败
	SMS_SEND_SUCCESS = "SMS_SEND_SUCCESS" // 短信发送成功
	SMS_SEND_REPEAT  = "SMS_SEND_REPEAT"  //短信重复发送
	SMS_SEND_LIMIT   = "SMS_SEND_LIMIT"   // 短信发送次数限制

	REPEAT_USER                = "REPEAT_USER"                //重复用户
	USER_NOT_FOUND             = "USER_NOT_FOUND"             //用户不存在
	NOT_MATCH_MERCHANT         = "NOT_MATCH_MERCHANT"         //未匹配商家
	REPEAT_MATCH_MERCHANT      = "REPEAT_MATCH_MERCHANT"      //重复匹配商家
	AGE_NOT_MATCH              = "AGE_NOT_MATCH"              //年龄不匹配
	INVALID_SESAME_SCORE_RANGE = "INVALID_SESAME_SCORE_RANGE" //无效的芝麻分数范围
	ID_CARD_ERROR              = "ID_CARD_ERROR"              //身份证号错误
	CHANNEL_CODE_ERROR         = "CHANNEL_CODE_ERROR"         //上游渠道码错误
	ID_CARD_BACK_ERROR         = "ID_CARD_BACK_ERROR"         //身份证人像反面错误
	CHANNEL_NOT_FOUND          = "CHANNEL_NOT_FOUND"          //渠道不存在
)

// 全局映射表
var Messages = map[string]string{
	"INVALID_REQUEST_PARAMETERS": "请求参数无效",
	"SUCCESS":                    "操作成功",
	"FAILED":                     "操作失败",

	"CREATION_FAILED":  "创建操作失败",
	"CREATION_SUCCESS": "创建操作成功",

	"UPDATE_FAILED":  "更新操作失败",
	"UPDATE_SUCCESS": "更新操作成功",

	"DELETE_FAILED":  "删除操作失败",
	"DELETE_SUCCESS": "删除操作成功",

	"LOGIN_ERROR":       "登录失败，请检查输入信息",
	"LOGIN_SUCCESS":     "登录成功",
	"CODE_ERROR":        "验证码错误，请重新输入",
	"USER_DISABLED":     "用户已被禁用，无法登录",
	"SIGNATURE_MISSING": "缺少签名信息",
	"INVALID_TIMESTAMP": "无效的时间戳",
	"TIMESTAMP_ERROR":   "时间戳错误",
	"ERR_READ_BODY":     "读取请求体失败",
	"INVALID_JSON":      "无效的JSON", //SON
	"INVALID_SIGNATURE": "无效的签名",   //签名

	"CAPTCHA_ERROR": "验证码输入错误",

	"AUTH_ERROR":        "登录已过期,为确保顺畅使用,请重新登录。",
	"AUTH_DEVICE_ERROR": "设备信息不存在",
	"AUTH_USER_ERROR":   "用户信息不存在",
	"PASSWORD_ERROR":    "密码错误，请重新输入",
	"USERNAME_EXISTS":   "用户名已存在，请尝试其他手机号或用户名",

	"SMS_SEND_ERROR":   "短信发送失败，请稍后重试",
	"SMS_SEND_SUCCESS": "短信发送成功",
	"SMS_SEND_REPEAT":  "短信重复发送，请稍后再试",
	"SMS_SEND_LIMIT":   "短信发送次数已达上限，请明日再试",

	"REPEAT_USER":                "重复用户信息",
	"USER_NOT_FOUND":             "用户不存在",
	"NOT_MATCH_MERCHANT":         "未匹配到相关商家信息",
	"REPEAT_MATCH_MERCHANT":      "重复匹配商家信息",
	"INVALID_SESAME_SCORE_RANGE": "无效的芝麻分数范围，请检查输入",
	"ID_CARD_ERROR":              "您提供的身份证号、姓名或手机号信息不匹配，请检查并重新输入。",
	"CHANNEL_CODE_ERROR":         "渠道码错误，请检查渠道是否正确",
	"ID_CARD_BACK_ERROR":         "身份证人像反面识别失败，请重新上传",
	"CHANNEL_NOT_FOUND":          "渠道信息不存在，请检查渠道码是否正确",
	"AGE_NOT_MATCH":              "年龄不匹配，请检查年龄是否符合要求",
}

func GetMessage(code string) string {
	if msg, exists := Messages[code]; exists {
		return msg
	}
	return "未知错误，请联系客服解决。"
}

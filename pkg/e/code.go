package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	// tag
	ERROR_TAT_EXIST               = 10101
	ERROR_TAG_NOT_EXIST           = 10102
	ERROR_TAG_ID_INVALID          = 10103
	ERROR_TAG_NAME_OVERSIZE       = 10104
	ERROR_TAG_STATE_INVALID       = 10105
	ERROR_TAG_MODIFIED_BY_INVALID = 10106

	// article
	ERROR_ARTICLE_NOT_EXIST = 10201

	// auth
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 10301
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10302
	ERROR_AUTH_TOKEN               = 10303
	ERROR_AUTH                     = 10304
	ERROR_AUTH_REQ_PARAM           = 10305
	ERROR_AUTH_VALIDATE_FAIL       = 10306
)

var CodeMsgs = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_TAT_EXIST:               "tag已存在",
	ERROR_TAG_NOT_EXIST:           "tag不存在",
	ERROR_TAG_ID_INVALID:          "没有tag id",
	ERROR_TAG_NAME_OVERSIZE:       "tag太长",
	ERROR_TAG_STATE_INVALID:       "tag的state非法",
	ERROR_TAG_MODIFIED_BY_INVALID: "tag的modified by非法",
	ERROR_ARTICLE_NOT_EXIST:       "文章不存在",

	// auth
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token鉴权错误",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token鉴权超时",
	ERROR_AUTH_TOKEN:               "token生成失败",
	ERROR_AUTH:                     "token错误",
	ERROR_AUTH_REQ_PARAM:           "auth请求参数错误",
	ERROR_AUTH_VALIDATE_FAIL:       "信息有误，校验不通过",
}

func GetMsg(code int) string {
	if msg, ok := CodeMsgs[code]; ok {
		return msg
	}

	return CodeMsgs[ERROR]
}

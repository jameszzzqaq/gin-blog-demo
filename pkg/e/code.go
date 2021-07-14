package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_TAT_EXIST               = 10001
	ERROR_TAG_NOT_EXIST           = 10002
	ERROR_TAG_ID_INVALID          = 10003
	ERROR_TAG_NAME_OVERSIZE       = 10004
	ERROR_TAG_STATE_INVALID       = 10005
	ERROR_TAG_MODIFIED_BY_INVALID = 10006
	ERROR_ARTICLE_NOT_EXIST       = 10020

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
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

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token鉴权错误",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token鉴权超时",
	ERROR_AUTH_TOKEN:               "token生成失败",
	ERROR_AUTH:                     "token错误",
}

func GetMsg(code int) string {
	if msg, ok := CodeMsgs[code]; ok {
		return msg
	}

	return CodeMsgs[ERROR]
}

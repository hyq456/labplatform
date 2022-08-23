package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	//CODE 100X ..user model err
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	//code 2XXX article model err
	ERROR_ART_NOT_EXITS = 2001
	//CODE 3XXX category model err
	ERROR_CATENAME_USER      = 3001
	ERROR_CATENAME_NOT_EXIST = 3002
	//CODE 4XXX EMAIL MODEL ERR
	ERROR_SEND_FAIL        = 4001
	ERROR_SAVE_VCODE       = 4002
	ERROR_VCODE_NOT_EXIT   = 4003
	ERROR_VCODE_TYPE_WRONG = 4004
)

var codemsg = map[int]string{
	SUCCSE:                   "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "用户名已存在",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIST:     "用户名不存在",
	ERROR_TOKEN_NOT_EXIST:    "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:      "TOKEN已过去",
	ERROR_TOKEN_WRONG:        "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:   "TOKEN格式错误",
	ERROR_CATENAME_USER:      "分类已存在",
	ERROR_CATENAME_NOT_EXIST: "分类不存在",
	ERROR_ART_NOT_EXITS:      "文字不存在",
	ERROR_USER_NO_RIGHT:      "用户无权限",
	ERROR_SEND_FAIL:          "发送邮件失败",
	ERROR_SAVE_VCODE:         "储存验证码失败",
	ERROR_VCODE_NOT_EXIT:     "验证码错误或已失效",
	ERROR_VCODE_TYPE_WRONG:   "验证码格式错误",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}

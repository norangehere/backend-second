package dto

//用户注册
type UserRegisterReq struct { //Body Params
	Email    string `json:"email"`    // 用户邮箱
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 用户密码
}

//获取登录验证码
type GetVeriResp struct { //Responses
	CAPTCHAID  string `json:"captcha_id"`  // 验证码 ID
	CAPTCHAURL string `json:"captcha_url"` // 验证码图片地址
}

//用户登录
type UserLoginReq struct { //Body Params
	CAPTCHAID    string `json:"captcha_id"`
	CAPTCHAValue string `json:"captcha_value"`
	Email        string `json:"email"`    // 用户邮箱
	Password     string `json:"password"` // 用户密码
}

//获取用户信息
type GetUserInfoResp struct { //Responses
	Email string `json:"email"` // 用户邮箱
	ID    int64  `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
}

//修改用户信息
type UserModifyInfoReq struct { //Body Params
	Email string `json:"email"` // 用户邮箱
	ID    int64  `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
}

//修改密码
type UserModifyPwdReq struct { //Body Params
	NewPwd string `json:"new_pwd"` // 新密码
	OldPwd string `json:"old_pwd"` // 旧密码
}

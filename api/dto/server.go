package dto

//短链接server
type ServerLinkReq struct {
	Short string `json:"short"` // 短链接唯一 ID
}

//验证码图片server
type ServerVeriReq struct {
	Target string `json:"target"` // 验证码对应的图片资源
}

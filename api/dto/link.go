package dto

import "gopkg.in/guregu/null.v4"

// 创建短链接
type LinkCreateReq struct { //Body Params
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}

type LinkCreateResp struct { //Responses
	Active    bool      `json:"active"`     // 服务状态
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}

// 删除短链接
type LinkDeleteReq struct { //Body Params
	Short string `json:"short"` // 短链接的唯一 ID
}

// 获取短链接信息
type GetLinkInfoReq struct { //Body Params
	Short string `json:"short"` // 短链接的唯一 ID
}

type GetLinkInfoResp struct { //Responses
	Active    bool      `json:"active"`     // 服务状态
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}

// 更新短链接信息
type LinkUpdateReq struct { //Body Params
	Active    bool      `json:"active"`     // 服务状态
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}

// 获取短链接信息
type GetLinkListReq struct { //Body & Header Params
	PageNumber  int64   `json:"page_number"` // 请求页码
	PageSize    int64   `json:"page_size"`   // 每页 item 数，-1 时代表全部返回
	AccessToken *string `json:"access_token,omitempty"`
}

type GetLinkListResp struct {
	Links []ShortLinkModel `json:"links"`
	Total int64            `json:"total"` // 总 item 数量
}

// Link
type ShortLinkModel struct {
	Active    bool      `json:"active"`     // 服务状态
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}

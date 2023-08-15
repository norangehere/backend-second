package controller

import (
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/model"

	"github.com/gin-gonic/gin"
)

// >>>>>>>>>>>>>>>>>> Interface  >>>>>>>>>>>>>>>>>>

type IFooController interface {
	GetFoo(*gin.Context, *dto.GetFooReq) (*dto.GetFooResp, error)
}

// >>>>>>>>>>>>>>>>>> Controller >>>>>>>>>>>>>>>>>>

// check interface implementation
var _ IFooController = (*FooController)(nil)

var NewFooController = func() *FooController {
	return &FooController{}
}

type FooController struct {
	// maybe some logic config to read from viper
	// or a service dependency
}

// link controller
type ILinkController interface { //定义tag link 下的操作
	Create(*gin.Context, *dto.LinkCreateReq) (*dto.LinkCreateResp, error)
	Delete(*gin.Context, *dto.LinkDeleteReq) error
	GetInfo(*gin.Context, *dto.GetLinkInfoReq) (*dto.GetLinkInfoResp, error)
	Update(*gin.Context, *dto.LinkUpdateReq) error
	GetList(*gin.Context, *dto.GetLinkListReq) (*dto.GetLinkListResp, error)
}

// user controller
type IUserController interface { //定义tag user 下的操作
	Register(*gin.Context, *dto.UserRegisterReq) error
	GetVeri(*gin.Context) (*dto.GetVeriResp, error)
	Login(*gin.Context, *dto.UserLoginReq) error
	GetInfo(*gin.Context) (*dto.GetUserInfoResp, error)
	ModifyInfo(*gin.Context, *dto.UserModifyInfoReq) error
	ModifyPwd(*gin.Context, *dto.UserModifyPwdReq) error
}

// server controller
type IServerController interface {
	Link(*gin.Context, *dto.ServerLinkReq) error
	Veri(*gin.Context, *dto.ServerVeriReq) error
}

// ---------------------- GetFoo ----------------------

func (c *FooController) GetFoo(ctx *gin.Context, req *dto.GetFooReq) (*dto.GetFooResp, error) {
	var resp dto.GetFooResp
	dao.DB(ctx).Model(model.Foo{Name: req.Name}).First(&resp)
	return &resp, nil
}

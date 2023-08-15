package route

import (
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"

	"github.com/gin-gonic/gin"
)

// 设置路由
func setupUserController(r *gin.RouterGroup) {
	lcw := UserCtlWrapper{
		ctl: controller.NewUserController(), //Factory method of ILinkController.
		// Implemented in controller package.
	}
	p := r.Group("/user")
	p.POST("/register", controller.ParseTokenMidware(), lcw.Register)
	p.GET("/getveri", controller.ParseTokenMidware(), lcw.GetVeri)
	p.POST("/login", controller.ParseTokenMidware(), lcw.Login)
	p.GET("/getinfo", controller.ParseTokenMidware(), lcw.GetInfo)
	p.POST("/modifyinfo", controller.ParseTokenMidware(), lcw.ModifyInfo)
	p.POST("/modifypwd", controller.ParseTokenMidware(), lcw.ModifyPwd)
}

type UserCtlWrapper struct { //Wrapper类隔离接口具体逻辑
	ctl controller.IUserController
}

// register user
func (w *UserCtlWrapper) Register(c *gin.Context) {
	var req dto.UserRegisterReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.Register(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, "register successfully")
}

// get verification
func (w *UserCtlWrapper) GetVeri(c *gin.Context) {
	//var req dto.UserRegisterReq
	//if err := dto.BindReq(c, &req); err != nil {
	//	dto.ResponseFail(c, err)
	//	return
	//}
	resp, err := w.ctl.GetVeri(c)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

// login
func (w *UserCtlWrapper) Login(c *gin.Context) {
	var req dto.UserLoginReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.Login(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, "login successfully")
}

// get user information
func (w *UserCtlWrapper) GetInfo(c *gin.Context) {
	//var req dto.UserRegisterReq
	//if err := dto.BindReq(c, &req); err != nil {
	//	dto.ResponseFail(c, err)
	//	return
	//}
	resp, err := w.ctl.GetInfo(c)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

// modify user information
func (w *UserCtlWrapper) ModifyInfo(c *gin.Context) {
	var req dto.UserModifyInfoReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.ModifyInfo(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, "modify information successfully")
}

// modify user password
func (w *UserCtlWrapper) ModifyPwd(c *gin.Context) {
	var req dto.UserModifyPwdReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.ModifyPwd(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, "modify password successfully")
}

package route

import (
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"

	"github.com/gin-gonic/gin"
)

// 设置路由
func setupServerController(r *gin.RouterGroup) {
	lcw := ServerCtlWrapper{
		ctl: controller.NewServerController(), //Factory method of ILinkController.
		// Implemented in controller package.
	}
	p := r.Group("/server")
	p.GET("/getlink", controller.ParseTokenMidware(), lcw.Link)
	p.GET("/getveri", controller.ParseTokenMidware(), lcw.Veri)
}

type ServerCtlWrapper struct { //Wrapper类隔离接口具体逻辑
	ctl controller.IServerController
}

// get link
func (w *ServerCtlWrapper) Link(c *gin.Context) {
	var req dto.ServerLinkReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.Link(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, "get link server successfully")
}

// get verification
func (w *ServerCtlWrapper) Veri(c *gin.Context) {
	var req dto.ServerVeriReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.Veri(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, "get verification server successfully")
}

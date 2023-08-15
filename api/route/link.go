// route/link.go
package route

import (
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"

	"github.com/gin-gonic/gin"
)

// 设置路由
func setupLinkController(r *gin.RouterGroup) {
	lcw := LinkCtlWrapper{
		ctl: controller.NewLinkController(), //Factory method of ILinkController.
		// Implemented in controller package.
	}
	p := r.Group("/link")
	p.POST("/create", controller.ParseTokenMidware(), lcw.Create)
	p.POST("/delete", controller.ParseTokenMidware(), lcw.Delete)
	p.GET("/getinfo", controller.ParseTokenMidware(), lcw.GetInfo)
	p.POST("/update", controller.ParseTokenMidware(), lcw.Update)
	p.GET("/getlist", controller.ParseTokenMidware(), lcw.GetList)
}

type LinkCtlWrapper struct { //Wrapper类隔离接口具体逻辑
	ctl controller.ILinkController
}

// create link
func (w *LinkCtlWrapper) Create(c *gin.Context) {
	var req dto.LinkCreateReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.Create(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

// delete link
func (w *LinkCtlWrapper) Delete(c *gin.Context) {
	var req dto.LinkDeleteReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.Delete(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, "delete successfully")
}

// get information
func (w *LinkCtlWrapper) GetInfo(c *gin.Context) {
	var req dto.GetLinkInfoReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.GetInfo(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

// update link
func (w *LinkCtlWrapper) Update(c *gin.Context) {
	var req dto.LinkUpdateReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.Update(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, "update successfully")
}

// get list
func (w *LinkCtlWrapper) GetList(c *gin.Context) {
	var req dto.GetLinkListReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.GetList(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

// 单独解析path
//func (w *SomeCtlWrapper) Get(c *gin.Context) {
//	var req dto.SomeGetReq
//	if err := dto.BindReq(c, &req); err != nil {
//		dto.ResponseFail(c, err)
//		return
//	}
//	path := c.Param("name")
//	resp, err := w.ctl.Get(c, path, &req)
//}

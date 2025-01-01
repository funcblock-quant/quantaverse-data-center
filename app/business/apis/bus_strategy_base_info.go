package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"quanta-admin/app/business/models"
	"quanta-admin/app/business/service"
	"quanta-admin/app/business/service/dto"
	"quanta-admin/common/actions"
)

type BusStrategyBaseInfo struct {
	api.Api
}

// GetPage 获取策略注册列表
// @Summary 获取策略注册列表
// @Description 获取策略注册列表
// @Tags 策略注册
// @Param strategyName query string false "策略名称"
// @Param strategyCategory query string false "策略交易类型"
// @Param status query string false "策略运行状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BusStrategyBaseInfo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/stragegy-base [get]
// @Security Bearer
func (e BusStrategyBaseInfo) GetPage(c *gin.Context) {
	req := dto.BusStrategyBaseInfoGetPageReq{}
	s := service.BusStrategyBaseInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.BusStrategyBaseInfo, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取策略注册失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取策略注册
// @Summary 获取策略注册
// @Description 获取策略注册
// @Tags 策略注册
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BusStrategyBaseInfo} "{"code": 200, "data": [...]}"
// @Router /api/v1/stragegy-base/{id} [get]
// @Security Bearer
func (e BusStrategyBaseInfo) Get(c *gin.Context) {
	req := dto.BusStrategyBaseInfoGetReq{}
	s := service.BusStrategyBaseInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.BusStrategyBaseInfo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取策略注册失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建策略注册
// @Summary 创建策略注册
// @Description 创建策略注册
// @Tags 策略注册
// @Accept application/json
// @Product application/json
// @Param data body dto.BusStrategyBaseInfoInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/stragegy-base [post]
// @Security Bearer
func (e BusStrategyBaseInfo) Insert(c *gin.Context) {
	req := dto.BusStrategyBaseInfoInsertReq{}
	s := service.BusStrategyBaseInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建策略注册失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改策略注册
// @Summary 修改策略注册
// @Description 修改策略注册
// @Tags 策略注册
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BusStrategyBaseInfoUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/stragegy-base/{id} [put]
// @Security Bearer
func (e BusStrategyBaseInfo) Update(c *gin.Context) {
	req := dto.BusStrategyBaseInfoUpdateReq{}
	s := service.BusStrategyBaseInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改策略注册失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除策略注册
// @Summary 删除策略注册
// @Description 删除策略注册
// @Tags 策略注册
// @Param data body dto.BusStrategyBaseInfoDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/stragegy-base [delete]
// @Security Bearer
func (e BusStrategyBaseInfo) Delete(c *gin.Context) {
	s := service.BusStrategyBaseInfo{}
	req := dto.BusStrategyBaseInfoDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除策略注册失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	models2 "quanta-admin/app/admin/models"
	"quanta-admin/app/grpc/client"
	"quanta-admin/app/grpc/proto/client/trigger_service"
	"strconv"

	"quanta-admin/app/business/models"
	"quanta-admin/app/business/service/dto"
	"quanta-admin/common/actions"
	cDto "quanta-admin/common/dto"
)

type BusPriceTriggerStrategyInstance struct {
	service.Service
}

// GetPage 获取BusPriceTriggerStrategyInstance列表
func (e *BusPriceTriggerStrategyInstance) GetPage(c *dto.BusPriceTriggerStrategyInstanceGetPageReq, p *actions.DataPermission, list *[]dto.BusPriceTriggerStrategyResp, count *int64) error {
	var err error
	var data models.BusPriceTriggerStrategyInstance
	var detail models.BusPriceMonitorForOptionHedging

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	for index, strategy := range *list {
		statistical := dto.BusPriceTriggerStrategyStatistical{}
		details := make([]models.BusPriceMonitorForOptionHedging, 0)
		err := e.Orm.Model(&detail).Where("strategy_instance_id = ?", strategy.Id).Order("id desc").Find(&details).Error
		if err != nil {
			e.Log.Errorf("BusPriceTriggerStrategyInstanceService Get details error:%s \r\n", err)
			return err
		}

		var apiConfig models.BusPriceTriggerStrategyApikeyConfig
		err = e.Orm.Unscoped().Model(&apiConfig).Where("id = ?", strategy.ApiConfig).First(&apiConfig).Error
		if err != nil {
			e.Log.Errorf("BusPriceTriggerStrategyInstanceService Get apiConfig error:%s \r\n", err)
			return err
		}
		sysUser := models2.SysUser{}
		err = e.Orm.Unscoped().Model(&sysUser).Where("user_id = ?", apiConfig.UserId).First(&sysUser).Error
		if err != nil {
			e.Log.Errorf("BusPriceTriggerStrategyInstanceService Get sysUser error:%s \r\n", err)
			return err
		}

		(*list)[index].ApiConfigData = apiConfig

		totalOrderNum := len(details)
		totalPnl := decimal.NewFromFloat(0)
		for _, d := range details {
			var pnl decimal.Decimal
			if d.Pnl == "" {
				pnl = decimal.NewFromFloat(0)
			} else {
				pnl, err = decimal.NewFromString(d.Pnl)
				if err != nil {
					e.Log.Errorf("BusPriceTriggerStrategyInstanceService Get details error:%s \r\n", err)
				}
			}
			totalPnl = totalPnl.Add(pnl)
		}
		statistical.OrderNum = totalOrderNum
		statistical.TotalPnl = totalPnl.StringFixed(8)
		(*list)[index].Statistical = statistical
		(*list)[index].Details = details
	}
	if err != nil {
		e.Log.Errorf("BusPriceTriggerStrategyInstanceService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取BusPriceTriggerStrategyInstance对象
func (e *BusPriceTriggerStrategyInstance) Get(d *dto.BusPriceTriggerStrategyInstanceGetReq, p *actions.DataPermission, model *models.BusPriceTriggerStrategyInstance) error {
	var data models.BusPriceTriggerStrategyInstance

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBusPriceTriggerStrategyInstance error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建BusPriceTriggerStrategyInstance对象
func (e *BusPriceTriggerStrategyInstance) Insert(c *dto.BusPriceTriggerStrategyInstanceInsertReq) error {
	var err error
	var data models.BusPriceTriggerStrategyInstance
	c.Generate(&data)
	data.Status = "created"
	//启动事务
	tx := e.Orm.Begin()
	if tx.Error != nil {
		e.Log.Errorf("BusPriceTriggerStrategyInstanceService error:%s \r\n", tx.Error)
		return tx.Error
	}

	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BusPriceTriggerStrategyInstanceService Insert error:%s \r\n", err)
		return err
	}
	e.Log.Infof("instance id : %d\r\n", data.Id)

	e.Log.Infof("api config key: %d", c.ApiConfig)
	var apiKeyConfig models.BusPriceTriggerStrategyApikeyConfig
	err = tx.First(&apiKeyConfig, c.ApiConfig).Error
	if err != nil {
		tx.Rollback()
		e.Log.Errorf("cannot found apikey config:%s \r\n", err)
	}

	//创建成功后， 自动通过grpc启动
	apiConfigReq := trigger_service.APIConfig{
		ApiKey:    apiKeyConfig.ApiKey,
		SecretKey: apiKeyConfig.SecretKey,
		Exchange:  apiKeyConfig.Exchange,
	}

	request := &trigger_service.StartTriggerRequest{
		InstanceId: strconv.Itoa(data.Id),
		OpenPrice:  c.OpenPrice,
		ClosePrice: c.ClosePrice,
		Side:       c.Side,
		Amount:     c.Amount,
		Symbol:     c.Symbol,
		StopTime:   strconv.FormatInt(c.CloseTime.UnixMilli(), 10),
		ApiConfig:  &apiConfigReq,
		UserId:     c.ExchangeUserId,
	}
	_, err = client.StartTriggerInstance(request)
	if err != nil {
		tx.Rollback()
		e.Log.Errorf("Service grpc start error:%s \r\n", err)
		return err
	}
	e.Log.Infof("instance id : %d grpc start success\r\n", data.Id)

	tx.Commit()

	err = e.Orm.Model(&data).Update("status", "started").Error
	if err != nil {
		e.Log.Errorf("start trigger update status failed:%s \r\n", err)
		return err
	}

	return nil
}

// StopInstance
// @Summary 暂停实例
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/stopInstance [post]
// @Security Bearer
func (e *BusPriceTriggerStrategyInstance) StopInstance(req *dto.StopTriggerInstanceRequest) error {
	var err error
	data := models.BusPriceTriggerStrategyInstance{}

	err = e.Orm.Model(&data).First(&data, req.InstanceId).Error
	if err != nil {
		e.Log.Errorf("Service StopInstance error:%s \r\n", err)
		return err
	}

	e.Log.Infof("stop instance id : %d\r\n", data.Id)
	request := trigger_service.StopTriggerRequest{
		InstanceId: strconv.Itoa(data.Id),
	}
	err = client.StopTriggerInstance(&request)
	err = nil
	if err != nil {
		e.Log.Errorf("Service StopInstance throw grpc error:%s \r\n", err)
		return err
	}
	err = e.Orm.Model(&data).
		Update("status", "stopped").
		Error

	if err != nil {
		e.Log.Errorf("Service StopInstance throw db error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改BusPriceTriggerStrategyInstance对象
func (e *BusPriceTriggerStrategyInstance) Update(c *dto.BusPriceTriggerStrategyInstanceUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.BusPriceTriggerStrategyInstance{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("BusPriceTriggerStrategyInstanceService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除BusPriceTriggerStrategyInstance
func (e *BusPriceTriggerStrategyInstance) Remove(d *dto.BusPriceTriggerStrategyInstanceDeleteReq, p *actions.DataPermission) error {
	var data models.BusPriceTriggerStrategyInstance

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveBusPriceTriggerStrategyInstance error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

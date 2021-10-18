package syncer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 资产列表
// @Description 分页获取资产列表
// @Tags 资产
// @accept json
// @Produce  json
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"执行成功"}
// @Failure 200 {object} Response {"code":!200,"data":null,"msg":"","detail"：""}
// @Router /assets [get]
func (a *api) GetAssets(c *gin.Context) {
	request := PageRequest{
		PageNum:  1,
		PageSize: defaultPageSize,
	}
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	if err := c.BindQuery(&request); err != nil {
		response.Code = RequestCode
		response.Msg = FailedRequest
		response.Detail = err.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	offset := (request.PageNum - 1) * request.PageSize
	var assets []*Asset
	if result := a.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&assets); result.Error != nil {
		a.logger.Error("GetAssets", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain BlockChain
	if result := a.db.First(&blockchain); result.Error != nil {
		a.logger.Error("GetAssets", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := blockchain.AssetCnt
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &List{
		PageResponse: PageResponse{
			NextPageNum: request.PageNum,
			PageSize:    request.PageSize,
			PageTotal:   pageTotal,
			Total:       total,
		},
		Items: assets,
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 资产信息
// @Description 获取资产信息
// @Tags 资产
// @accept json
// @Produce  json
// @Param name path string true "资产名称"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /assets/{name} [get]
func (a *api) GetAsset(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	cond := map[string]interface{}{}
	cond["name"] = c.Param("name")
	var asset Asset
	if result := a.db.Find(&asset, cond); result.Error != nil {
		a.logger.Error("GetAsset", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
	} else if result.RowsAffected > 0 {
		response.Data = asset
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 资产相关相关交易
// @Description 分页获取资产相关交易
// @Tags 资产
// @accept json
// @Produce  json
// @Param name path string true "资产名称"
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /assets/{name}/transactions [get]
func (a *api) GetAssetTransactions(c *gin.Context) {
	request := PageRequest{
		PageNum:  1,
		PageSize: defaultPageSize,
	}
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	if err := c.BindQuery(&request); err != nil {
		response.Code = RequestCode
		response.Msg = FailedRequest
		response.Detail = err.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	offset := (request.PageNum - 1) * request.PageSize
	var total int64
	var transactions []*Transaction
	if result := a.db.Joins("Asset").Count(&total).Order("transactions.id desc").Offset(offset).Limit(request.PageSize).Find(&transactions, "assets.name = ?", c.Param("name")); result.Error != nil {
		a.logger.Error("GetAssetTransactions", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}

	var blockchain BlockChain
	if result := a.db.Find(&blockchain); result.Error != nil {
		a.logger.Error("GetAssetTransactions", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = &List{
		PageResponse: PageResponse{
			NextPageNum: request.PageNum + 1,
			PageSize:    request.PageSize,
			PageTotal:   pageTotal,
			Total:       total,
		},
		Items: func() []*Transaction {
			for _, tx := range transactions {
				tx.Confirmed = blockchain.BlockCnt - tx.Height + 1
			}
			return transactions
		}(),
	}
	c.JSON(http.StatusOK, response)
}

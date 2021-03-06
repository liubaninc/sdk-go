package syncer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 账户地址列表
// @Description 分页获取账户地址列表
// @Tags 账户地址
// @accept json
// @Produce  json
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /addresses [get]
func (a *api) GetAddresses(c *gin.Context) {
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
	var addresses []*Address
	if result := a.db.Order("id desc").Offset(offset).Limit(request.PageSize).Find(&addresses); result.Error != nil {
		a.logger.Error("GetAddresses", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain BlockChain
	if result := a.db.Find(&blockchain); result.Error != nil {
		a.logger.Error("GetAddresses", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := blockchain.AddressCnt
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &List{
		PageResponse: PageResponse{
			Total:       total,
			NextPageNum: request.PageNum + 1,
			PageSize:    request.PageSize,
			PageTotal:   pageTotal,
		},
		Items: addresses,
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 账户地址信息
// @Description 获取账户地址信息
// @Tags 账户地址
// @accept json
// @Produce  json
// @Param name path string true "账户地址"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /addresses/{name} [get]
func (a *api) GetAddress(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	cond := map[string]interface{}{}
	cond["address"] = c.Param("name")
	var address Address
	if result := a.db.Find(&address, cond); result.Error != nil {
		a.logger.Error("GetAddress", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
	} else if result.RowsAffected > 0 {
		response.Data = address
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 账户地址相关交易
// @Description 分页获取账户地址相关交易
// @Tags 账户地址
// @accept json
// @Produce  json
// @Param name path string true "账户地址"
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /addresses/{name}/transactions [get]
func (a *api) GetAddressTransactions(c *gin.Context) {
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
	if result := a.db.Joins("Address").Count(&total).Order("transactions.id desc").Offset(offset).Limit(request.PageSize).Find(&transactions, "addresses.address = ?", c.Param("name")); result.Error != nil {
		a.logger.Error("GetAddressTransactions", "error", result.Error)
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
		a.logger.Error("GetTransactions", "error", result.Error)
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

// @Summary 账户地址相关资产
// @Description 分页获取账户地址创建的相关资产
// @Tags 账户地址
// @accept json
// @Produce  json
// @Param name path string true "账户地址"
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /addresses/{name}/assets [get]
func (a *api) GetAddressAssets(c *gin.Context) {
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
	var assets []*Asset
	if result := a.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&assets, "creator = ?", c.Param("name")); result.Error != nil {
		a.logger.Error("GetAddressAssets", "error", result.Error)
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

	response.Data = &List{
		PageResponse: PageResponse{
			NextPageNum: request.PageNum + 1,
			PageSize:    request.PageSize,
			PageTotal:   pageTotal,
			Total:       total,
		},
		Items: assets,
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 账户地址相关合约
// @Description 分页获取账户地址创建的相关合约
// @Tags 账户地址
// @accept json
// @Produce  json
// @Param name path string true "账户地址"
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /addresses/{name}/contracts [get]
func (a *api) GetAddressContracts(c *gin.Context) {
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
	var items []*Contract
	if result := a.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&items, "creator = ?", c.Param("name")); result.Error != nil {
		a.logger.Error("GetAddressAssets", "error", result.Error)
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

	response.Data = &List{
		PageResponse: PageResponse{
			NextPageNum: request.PageNum + 1,
			PageSize:    request.PageSize,
			PageTotal:   pageTotal,
			Total:       total,
		},
		Items: items,
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 获取账户地址相关交易
// @Description 分页获取账户地址提交的相关合约代码
// @Tags 账户地址
// @accept json
// @Produce  json
// @Param name path string true "账户地址"
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /addresses/{name}/contractcodes [get]
func (a *api) GetAddressContractCodes(c *gin.Context) {
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
	var items []*ContractCode
	if result := a.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&items, "creator = ?", c.Param("name")); result.Error != nil {
		a.logger.Error("GetAddressAssets", "error", result.Error)
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

	response.Data = &List{
		PageResponse: PageResponse{
			NextPageNum: request.PageNum + 1,
			PageSize:    request.PageSize,
			PageTotal:   pageTotal,
			Total:       total,
		},
		Items: items,
	}
	c.JSON(http.StatusOK, response)
}

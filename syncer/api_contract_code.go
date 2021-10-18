package syncer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 合约代码列表
// @Description 分页获取合约代码列表
// @Tags 合约代码
// @accept json
// @Produce  json
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"执行成功"}
// @Failure 200 {object} Response {"code":!200,"data":null,"msg":"","detail"：""}
// @Router /contractcodes [get]
func (a *api) GetContractCodes(c *gin.Context) {
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
	var contracts []*ContractCode
	if result := a.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&contracts); result.Error != nil {
		a.logger.Error("GetContractCodes", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain BlockChain
	if result := a.db.First(&blockchain); result.Error != nil {
		a.logger.Error("GetContractCodes", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := blockchain.ContractCnt
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
		Items: contracts,
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 合约代码信息
// @Description 获取合约代码信息
// @Tags 合约代码
// @accept json
// @Produce  json
// @Param name path string true "合约代码哈希"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /contractcodess/{name} [get]
func (a *api) GetContractCode(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	cond := map[string]interface{}{}
	cond["hash"] = c.Param("name")
	var contract ContractCode
	if result := a.db.Find(&contract, cond); result.Error != nil {
		a.logger.Error("GetContractCode", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
	} else if result.RowsAffected > 0 {
		response.Data = contract
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 合约代码相关交易
// @Description 分页获取合约代码相关交易
// @Tags 合约代码
// @accept json
// @Produce  json
// @Param name path string true "合约代码哈希"
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /contractcodes/{name}/transactions [get]
func (a *api) GetContractCodeTransactions(c *gin.Context) {
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
	if result := a.db.Joins("ContractCode").Count(&total).Order("transactions.id desc").Offset(offset).Limit(request.PageSize).Find(&transactions, "contractcodes.hash = ?", c.Param("name")); result.Error != nil {
		a.logger.Error("GetContractCodeTransactions", "error", result.Error)
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
		a.logger.Error("GetContractCodeTransactions", "error", result.Error)
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

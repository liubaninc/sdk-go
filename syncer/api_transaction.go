package syncer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 交易列表
// @Description 分页获取交易列表
// @Tags 交易
// @accept json
// @Produce  json
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"执行成功"}
// @Failure 200 {object} Response {"code":!200,"data":null,"msg":"","detail"：""}
// @Router /transactions [get]
func (a *api) GetTransactions(c *gin.Context) {
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
	var transactions []*Transaction
	if result := a.db.Order("id desc").Offset(offset).Limit(request.PageSize).Find(&transactions); result.Error != nil {
		a.logger.Error("GetTransactions", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
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

	total := blockchain.TxCnt
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
		Items: func() []*Transaction {
			for _, tx := range transactions {
				tx.Confirmed = blockchain.BlockCnt - tx.Height + 1
			}
			return transactions
		}(),
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 交易信息
// @Description 获取交易信息
// @Tags 交易
// @accept json
// @Produce  json
// @Param hash path string true "交易哈希"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /transactions/{name} [get]
func (a *api) GetTransaction(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	var tx Transaction
	cond := map[string]interface{}{}
	cond["hash"] = c.Param("hash")
	if result := a.db.Preload("Msgs").Find(&tx, cond); result.Error != nil {
		a.logger.Error("GetTransaction", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
	} else if result.RowsAffected > 0 {
		var blockchain BlockChain
		if result := a.db.Last(&blockchain); result.Error != nil {
			a.logger.Error("GetTransactions", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = FailedExecute
			response.Detail = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}

		tx.Confirmed = blockchain.BlockCnt - tx.Height + 1
		response.Data = tx
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

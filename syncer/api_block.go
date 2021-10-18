package syncer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (a *api) GetBlockChain(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	var blockchain BlockChain
	if result := a.db.First(&blockchain); result.Error != nil {
		a.logger.Error("GetBlockChain", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
	} else {
		response.Data = blockchain
	}
	c.JSON(http.StatusOK, response)
}

func (a *api) GetBlocks(c *gin.Context) {
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
	var blocks []*Block
	if result := a.db.Order("id desc").Offset(offset).Limit(request.PageSize).Find(&blocks); result.Error != nil {
		a.logger.Error("GetBlocks", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain BlockChain
	if result := a.db.Find(&blockchain); result.Error != nil {
		a.logger.Error("GetBlocks", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := blockchain.BlockCnt
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
		Items: blocks,
	}
	c.JSON(http.StatusOK, response)
}

func (a *api) GetBlock(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	var block Block
	cond := map[string]interface{}{}
	id := c.Param("id")
	if height, err := strconv.ParseInt(id, 10, 64); err != nil {
		cond["hash"] = id
	} else {
		cond["height"] = height
	}

	if result := a.db.Preload("Txs").Find(&block, cond); result.Error != nil {
		a.logger.Error("GetBlock", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
	} else if result.RowsAffected > 0 {
		response.Data = &block
	}

	c.JSON(http.StatusOK, response)
}

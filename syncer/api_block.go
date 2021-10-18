package syncer

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 区块链TPS统计
// @Description 获取区块链TPS统计信息
// @Tags 区块
// @accept json
// @Produce  json
// @Success 200 {object} Response {"code":200,"data":null,"msg":"执行成功"}
// @Failure 200 {object} Response {"code":!200,"data":null,"msg":"","detail"：""}
// @Router /charts [get]
func (a *api) Charts(c *gin.Context) {
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

	var blockChainChart BlockChainChart
	if result := a.db.Last(&blockChainChart); result.Error != nil {
		a.logger.Error("Charts", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	offset := (request.PageNum - 1) * request.PageSize
	var blocks map[string]interface{}
	if result := a.db.Model(&BlockChainChart{}).Order("ID desc").Offset(offset).Limit(request.PageSize).Select("time", "block_num as number").Find(&blocks); result.Error != nil {
		a.logger.Error("Charts", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := int64(blockChainChart.ID)
	pageTotal := total / int64(request.PageSize)
	if int64(blockChainChart.ID)%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &List{
		PageResponse: PageResponse{
			Total:       total,
			PageSize:    request.PageSize,
			NextPageNum: request.PageNum + 1,
			PageTotal:   pageTotal,
		},
		Items: blocks,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 搜索
// @Description 搜索区块高度、区块哈希、交易哈希、账户地址、资产、合约、合约代码
// @Tags 区块
// @accept json
// @Produce  json
// @Param name path string true "未知"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /search/{name} [get]
func (a *api) Search(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	content := c.Param("content")
	if _, err := strconv.ParseInt(content, 10, 64); err == nil {
		response.Data = "blocks/" + content
		c.JSON(http.StatusOK, response)
		return
	}

	if _, err := sdk.AccAddressFromBech32(content); err == nil {
		response.Data = "addresses/" + content
		c.JSON(http.StatusOK, response)
		return
	}

	if _, err := hex.DecodeString(content); err == nil {
		var block Block
		if result := a.db.Find(&block, "hash = ?", content); result.RowsAffected == 1 {
			response.Data = "blocks/" + content
			c.JSON(http.StatusOK, response)
			return
		}

		var tx Transaction
		if result := a.db.Find(&tx, "hash = ?", content); result.RowsAffected == 1 {
			response.Data = "transactions/" + content
			c.JSON(http.StatusOK, response)
			return
		}

		var contractCode ContractCode
		if result := a.db.Find(&contractCode, "hash = ?", content); result.RowsAffected == 1 {
			response.Data = "contractcodes/" + content
			c.JSON(http.StatusOK, response)
			return
		}
	}

	var asset Asset
	if result := a.db.Find(&asset, "name = ?", content); result.RowsAffected == 1 {
		response.Data = "assets/" + content
		c.JSON(http.StatusOK, response)
		return
	}

	var contract Contract
	if result := a.db.Find(&contract, "name = ?", content); result.RowsAffected == 1 {
		response.Data = "contracts/" + content
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

// @Summary 区块链总览
// @Description 获取区块链总览信息
// @Tags 区块
// @accept json
// @Produce  json
// @Success 200 {object} Response {"code":200,"data":null,"msg":"执行成功"}
// @Failure 200 {object} Response {"code":!200,"data":null,"msg":"","detail"：""}
// @Router /blockchain [get]
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

// @Summary 区块列表
// @Description 分页获取区块列表
// @Tags 区块
// @accept json
// @Produce  json
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"执行成功"}
// @Failure 200 {object} Response {"code":!200,"data":null,"msg":"","detail"：""}
// @Router /blocks [get]
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

// @Summary 区块信息
// @Description 获取区块信息
// @Tags 区块
// @accept json
// @Produce  json
// @Param id path string true "区块高度/区块哈希"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /blocks/{name} [get]
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

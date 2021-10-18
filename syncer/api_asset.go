package syncer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

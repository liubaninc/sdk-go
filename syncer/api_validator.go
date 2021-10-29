package syncer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 验证者列表
// @Description 分页获验证者列表
// @Tags 验证者
// @accept json
// @Produce  json
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"执行成功"}
// @Failure 200 {object} Response {"code":!200,"data":null,"msg":"","detail"：""}
// @Router /validators [get]
func (a *api) GetValidators(c *gin.Context) {
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
	var validators []*Validator
	if result := a.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&validators); result.Error != nil {
		a.logger.Error("GetValidators", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain BlockChain
	if result := a.db.First(&blockchain); result.Error != nil {
		a.logger.Error("GetValidators", "error", result.Error)
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
		Items: validators,
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 验证者信息
// @Description 获取验证者信息
// @Tags 验证者
// @accept json
// @Produce  json
// @Param name path string true "账户地址"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /validators/{name} [get]
func (a *api) GetValidator(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	cond := map[string]interface{}{}
	cond["creator"] = c.Param("name")
	var validator Validator
	if result := a.db.Find(&validator, cond); result.Error != nil {
		a.logger.Error("GetValidator", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
	} else if result.RowsAffected > 0 {
		response.Data = validator
	}
	c.JSON(http.StatusOK, response)
}

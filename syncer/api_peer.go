package syncer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 节点列表
// @Description 分页获取节点列表
// @Tags 节点
// @accept json
// @Produce  json
// @Param page_num query int false "页码"
// @Param page_size query int false "个数"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"执行成功"}
// @Failure 200 {object} Response {"code":!200,"data":null,"msg":"","detail"：""}
// @Router /peers [get]
func (a *api) GetPeers(c *gin.Context) {
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
	var peers []*Peer
	if result := a.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&peers); result.Error != nil {
		a.logger.Error("GetPeers", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain BlockChain
	if result := a.db.First(&blockchain); result.Error != nil {
		a.logger.Error("GetPeers", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := blockchain.PeerCnt
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
		Items: peers,
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 节点信息
// @Description 获取节点信息
// @Tags 节点
// @accept json
// @Produce  json
// @Param name path string true "节点ID"
// @Success 200 {object} Response {"code":200,"data":null,"msg":"","detail"：""}
// @Router /peers/{name} [get]
func (a *api) GetPeer(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	cond := map[string]interface{}{}
	cond["node_id"] = c.Param("name")
	var peer Peer
	if result := a.db.Find(&peer, cond); result.Error != nil {
		a.logger.Error("GetPeer", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = FailedExecute
		response.Detail = result.Error.Error()
	} else if result.RowsAffected > 0 {
		response.Data = peer
	}
	c.JSON(http.StatusOK, response)
}

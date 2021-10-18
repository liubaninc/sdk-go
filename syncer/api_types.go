package syncer

var (
	// OKMsg msg
	OKMsg         = "执行成功"
	FailedRequest = "请求参数错误"
	FailedExecute = "执行错误"
	// OKCode Ok
	OKCode      = 200
	RequestCode = 3001
	ExecuteCode = 3002
)

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}

type PageRequest struct {
	PageNum  int `form:"page_num" json:"page_num" xml:"page_num"`
	PageSize int `form:"page_size" json:"page_size" xml:"page_size"`
}

type PageResponse struct {
	Total       int64 `json:"total"`
	PageTotal   int64 `json:"page_total"`
	NextPageNum int   `json:"next_page_num"`
	PageSize    int   `json:"page_size"`
}

type List struct {
	PageResponse PageResponse
	Items        interface{} `json:"items"`
}

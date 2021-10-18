package syncer

import (
	"github.com/gin-gonic/gin"
	"github.com/tendermint/tendermint/libs/log"
	"gorm.io/gorm"
)

const (
	defaultPageSize = 10
)

type api struct {
	db     *gorm.DB
	logger log.Logger
}

func NewApi(log log.Logger, db *gorm.DB) *api {
	return &api{
		logger: log.With("module", "api"),
		db:     db,
	}
}

func (a *api) RegisterApi(router *gin.Engine) {
	router.GET("/search/:content", a.Search)
	router.GET("/charts", a.Charts)
	router.GET("/blockchain", a.GetBlockChain)
	router.GET("/blocks", a.GetBlocks)
	router.GET("/blocks/:id", a.GetBlock)
	router.GET("/transactions", a.GetTransactions)
	router.GET("/transactions/:hash", a.GetTransaction)
	router.GET("/addresses", a.GetAddresses)
	router.GET("/addresses/:name", a.GetAddress)
	router.GET("/addresses/:name/assets", a.GetAddressAssets)
	router.GET("/addresses/:name/contracts", a.GetAddressContracts)
	router.GET("/addresses/:name/contractcodes", a.GetAddressContractCodes)
	router.GET("/addresses/:name/transactions", a.GetAddressTransactions)
	router.GET("/assets", a.GetAssets)
	router.GET("/assets/:name", a.GetAsset)
	router.GET("/assets/:name/transactions", a.GetAssetTransactions)
	router.GET("/contractcodes", a.GetContractCodes)
	router.GET("/contractcodes/:name", a.GetContractCode)
	router.GET("/contractcodes/:name/transactions", a.GetContractCodeTransactions)
	router.GET("/contracts", a.GetContracts)
	router.GET("/contracts/:name", a.GetContract)
	router.GET("/contracts/:name/transactions", a.GetContractTransactions)
	router.GET("/peers", a.GetPeers)
	router.GET("/peers/:name", a.GetPeer)
	router.GET("/validators", a.GetValidators)
	router.GET("/validators/:name", a.GetValidator)
}

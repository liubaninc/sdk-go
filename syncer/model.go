package syncer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&BlockChain{})
	db.AutoMigrate(&Block{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&Msg{})
	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Asset{})
	db.AutoMigrate(&ContractCode{})
	db.AutoMigrate(&Contract{})
	db.AutoMigrate(&Peer{})
	db.AutoMigrate(&Validator{})
}

type BlockChainChart struct {
	gorm.Model      `json:"-"`
	Time            string `json:"time" gorm:"unique"`
	BlockCnt        int64  `json:"block_num"`
	TxCnt           int64  `json:"tx_num"`
	MsgCnt          int64  `json:"msg_num"`
	AssetCnt        int64  `json:"asset_num"`
	ContractCnt     int64  `json:"contract_num"`
	ContractCodeCnt int64  `json:"contract_code_num"`
}

type BlockChain struct {
	gorm.Model      `json:"-"`
	BlockCnt        int64 `json:"block_num"`
	TxCnt           int64 `json:"tx_num"`
	MsgCnt          int64 `json:"msg_num"`
	AddressCnt      int64 `json:"address_num"`
	AssetCnt        int64 `json:"asset_num"`
	ContractCodeCnt int64 `json:"contract_code_num"`
	ContractCnt     int64 `json:"contract_num"`
	ValidatorCnt    int64 `json:"validator_num"`
	PeerCnt         int64 `json:"peer_num"`
}

type Block struct {
	gorm.Model `json:"-"`
	Hash       string    `gorm:"uniqueIndex" json:"hash"`
	Height     int64     `gorm:"uniqueIndex" json:"height"`
	PrevHash   string    `gorm:"uniqueIndex" json:"prev_hash"`
	Proposer   string    `json:"proposer"`
	Time       time.Time `json:"time"`
	Size       int       `json:"size"`
	TxCnt      int       `json:"tx_num"`

	Txs []*Transaction `gorm:"foreignKey:Height;references:Height" json:"txs"`

	assetMap        map[string]*Asset
	addressMap      map[string]*Address
	contractMap     map[string]*Contract
	contractCodeMap map[string]*ContractCode
	peerMap         map[string]*Peer
	validatorMap    map[string]*Validator
}

type Transaction struct {
	gorm.Model `json:"-"`
	Hash       string    `gorm:"uniqueIndex" json:"hash"`
	Height     int64     `gorm:"index" json:"height"`
	Status     bool      `json:"status" `
	RawLog     string    `json:"raw_log" `
	Memo       string    `json:"memo" `
	Fees       string    `json:"fee" `
	Time       time.Time `json:"time" `
	Size       int       `json:"size"`
	MsgCnt     int       `json:"msg_num" `
	Raw        string    `json:"-" `
	Type       string    `json:"type"`

	Msgs          []*Msg          `gorm:"foreignKey:Hash;references:Hash" json:"-" `
	Addresses     []*Address      `gorm:"many2many:tx_addresses;" json:"-" `
	Assets        []*Asset        `gorm:"many2many:tx_assets;" json:"-" `
	Contracts     []*Contract     `gorm:"many2many:tx_contracts;" json:"-" `
	ContractCodes []*ContractCode `gorm:"many2many:tx_contract_codes;" json:"-" `
	Peers         []*Peer         `gorm:"many2many:tx_peers;" json:"-" `
	Validators    []*Validator    `gorm:"many2many:tx_validators;" json:"-" `

	Confirmed int64 `gorm:"-" json:"confirmed" `

	assetMap        map[string]*Asset
	addressMap      map[string]*Address
	contractMap     map[string]*Contract
	contractCodeMap map[string]*ContractCode
	peerMap         map[string]*Peer
	validatorMap    map[string]*Validator
}

type Msg struct {
	gorm.Model
	Hash  string `gorm:"index"`
	Route string `gorm:"index"`
	Type  string `gorm:"index"`
	Raw   string
}

type Address struct {
	gorm.Model
	Address       string `gorm:"uniqueIndex"`
	AccountNumber uint64
	Sequence      uint64
	Balance       string
	Roles         string
	Height        int64
	Time          time.Time

	Txs []*Transaction `gorm:"many2many:tx_addresses;"`

	amounts sdk.Coins `gorm:"-"`
}

type Asset struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Creator     string
	Display     string
	Meta        string
	IssueAmount string
	BurnAmount  string
	Height      int64
	Time        time.Time

	Txs []*Transaction `gorm:"many2many:tx_assets;"`

	issueAmount sdk.Coin `gorm:"-"`
	burnAmount  sdk.Coin `gorm:"-"`
}

type ContractCode struct {
	gorm.Model
	Hash      string `gorm:"uniqueIndex"`
	Creator   string
	Name      string
	Code      string
	SRC       string
	ABI       string
	Approvers string
	Approved  string
	Status    string
	Runtime   string
	Module    string
	Height    int64
	Time      time.Time

	Txs       []*Transaction `gorm:"many2many:tx_contract_codes;"`
	Contracts *[]Contract    `gorm:"many2many:tx_contract_code_contracs;"`

	approved []string
}

type Contract struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex"`
	Creator  string
	Hash     string `gorm:"index"`
	InitArgs string
	Status   string
	Version  uint32
	Height   int64
	Time     time.Time

	ContractCodes []*ContractCode `gorm:"many2many:tx_contract_code_contracs;"`
	Txs           []*Transaction  `gorm:"many2many:tx_contracts;"`
}

type Peer struct {
	gorm.Model
	NodeID  string `gorm:"unique"`
	Addr    string
	Moniker string
	Height  int64
	Time    time.Time

	Delete bool
}

type Validator struct {
	gorm.Model
	Creator string `gorm:"unique"`
	Name    string
	PubKey  string `gorm:"unique"`
	Pow     int64
	NodeID  string
	Height  int64
	Time    time.Time

	Delete bool
}

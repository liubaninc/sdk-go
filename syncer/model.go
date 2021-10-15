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

type BlockChain struct {
	gorm.Model
	BlockCnt        int64
	TxCnt           int64
	MsgCnt          int64
	AssetCnt        int64
	ContractCodeCnt int64
	ContractCnt     int64
	ValidatorCnt    int64
	PeerCnt         int64
}

type Block struct {
	gorm.Model
	Hash     string `gorm:"uniqueIndex"`
	Height   int64  `gorm:"uniqueIndex"`
	PrevHash string
	Proposer string
	Time     time.Time
	Size     int
	TxCnt    int

	Txs []*Transaction `gorm:"foreignKey:Height;references:Height"`

	assetMap        map[string]*Asset
	addressMap      map[string]*Address
	contractMap     map[string]*Contract
	contractCodeMap map[string]*ContractCode
	peerMap         map[string]*Peer
	validatorMap    map[string]*Validator
}

type Transaction struct {
	gorm.Model
	Hash   string `gorm:"uniqueIndex"`
	Height int64  `gorm:"index" `
	Status bool
	RawLog string
	Memo   string
	Fees   string
	Time   time.Time
	Size   int
	MsgCnt int
	Raw    string

	Msgs          []*Msg          `gorm:"foreignKey:Hash;references:Hash"`
	Addresses     []*Address      `gorm:"many2many:tx_addresses;"`
	Assets        []*Asset        `gorm:"many2many:tx_assets;"`
	Contracts     []*Contract     `gorm:"many2many:tx_contracts;"`
	ContractCodes []*ContractCode `gorm:"many2many:tx_contract_codes;"`
	Peers         []*Peer         `gorm:"many2many:tx_peers;"`
	Validators    []*Validator    `gorm:"many2many:tx_validators;"`

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
}

type Validator struct {
	gorm.Model
	Name   string
	PubKey string `gorm:"unique"`
	Pow    int64
	NodeID string
	Height int64
	Time   time.Time
}

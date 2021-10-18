package syncer

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	peertypes "github.com/liubaninc/sdk-go/modules/peer/types"
	storagetypes "github.com/liubaninc/sdk-go/modules/storage/types"
	validatortypes "github.com/liubaninc/sdk-go/modules/validator/types"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/tendermint/tendermint/crypto/tmhash"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/golang/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/sdk-go/client"
	accountypes "github.com/liubaninc/sdk-go/modules/account/types"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
	wasmtypes "github.com/liubaninc/sdk-go/modules/wasm/types"
	"github.com/liubaninc/sdk-go/params"
	"github.com/tendermint/tendermint/libs/log"
	"gorm.io/gorm"
)

var (
	encodingConfig = params.MakeEncodingConfig()
)

var _ client.SyncedHandler = &handler{}

type HandleTxFunc func(db *gorm.DB, tx *tx.GetTxResponse) error

func hash(bytes []byte) string {
	return fmt.Sprintf("%X", tmhash.Sum(bytes))
}

func NewHander(log log.Logger, db *gorm.DB, client *client.Client, handle ...HandleTxFunc) *handler {
	return &handler{
		log:           log.With("module", "handler"),
		db:            db,
		client:        client,
		handleTxFuncs: handle,
	}
}

type handler struct {
	log           log.Logger
	db            *gorm.DB
	client        *client.Client
	handleTxFuncs []HandleTxFunc
}

func (h handler) GetLastBlockHeight() int64 {
	var mBlock Block
	if err := h.db.Order("id desc").Find(&mBlock).Error; err != nil {
		panic(err)
	}
	return mBlock.Height
}

func (h handler) HandleGenesisTxs(tx []*tx.GetTxResponse) error {
	h.log.Info("HandleGenesis", "height", 0)
	return h.db.Transaction(func(tx *gorm.DB) error {
		c := h.client.WithHeight(1)
		resAccount, err := c.GetAccounts(nil, 0, 1000, false)
		if err != nil {
			h.log.Error("HandleGenesisTxs", "GetAccounts", err)
			return err
		}
		var addresses []*Address
		for _, account := range resAccount.Accounts {
			var acc authtypes.BaseAccount
			if err := proto.Unmarshal(account.Value, &acc); err != nil {
				continue
			}
			res, err := c.GetAccountACL(acc.Address)
			if err != nil {
				h.log.Error("HandleGenesisTxs", "GetAccountACL", err)
				return err
			}
			res2, err := h.client.GetAccountAllBalances(acc.Address)
			if err != nil {
				h.log.Error("HandleGenesisTxs", "GetAccountAllBalances", err)
				return err
			}
			addresses = append(addresses, &Address{
				Address:       acc.Address,
				AccountNumber: acc.AccountNumber,
				Sequence:      acc.Sequence,
				Roles:         strings.Join(res.Account.Roles, ","),
				Balance:       res2.Balances.String(),
			})
		}

		resAsset, err := c.GetTokenAll(nil, 0, 1000, false)
		if err != nil {
			h.log.Error("HandleGenesisTxs", "GetTokenAll", err)
			return err
		}
		var assets []*Asset
		for _, token := range resAsset.Token {
			meta, _ := json.Marshal(token)
			assets = append(assets, &Asset{
				Name:        token.Metadata.Base,
				Display:     token.Metadata.Display,
				Meta:        string(meta),
				Creator:     token.Creator,
				IssueAmount: token.Issued.String(),
				BurnAmount:  token.Burned.String(),
			})
		}

		resPeer, err := c.GetPeerAll(nil, 0, 1000, false)
		if err != nil {
			h.log.Error("HandleGenesisTxs", "GetPeerAll", err)
			return err
		}
		var peers []*Peer
		for _, peer := range resPeer.Peer {
			peers = append(peers, &Peer{
				NodeID: peer.Index,
				Addr:   peer.Addr,
			})
		}

		resValidtor, err := c.GetValidatorAll(nil, 0, 1000, false)
		if err != nil {
			h.log.Error("HandleGenesisTxs", "GetValidatorAll", err)
			return err
		}
		var validators []*Validator
		for _, validator := range resValidtor.Validator {
			validators = append(validators, &Validator{
				Creator: validator.Creator,
				PubKey:  validator.PubKey,
				Pow:     validator.Power,
			})
		}

		if err := tx.Create(addresses).Error; err != nil {
			h.log.Error("HandleGenesisTxs", "Create(addresses)", err)
			return err
		}
		if err := tx.Create(assets).Error; err != nil {
			h.log.Error("HandleGenesisTxs", "Create(assets)", err)
			return err
		}
		if err := tx.Create(peers).Error; err != nil {
			h.log.Error("HandleGenesisTxs", "Create(peers)", err)
			return err
		}
		if err := tx.Create(validators).Error; err != nil {
			h.log.Error("HandleGenesisTxs", "Create(validators)", err)
			return err
		}

		return nil
	})
}

func (h handler) HandlePrevBlock(block *tmservice.GetBlockByHeightResponse) error {
	h.log.Info("HandlePrevBlock", "height", block.Block.Header.Height)
	var mBlock Block
	if err := h.db.Last(&mBlock).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		h.log.Error("HandlePrevBlock db.Last(&mBlock)", "error", err)
		return err
	}

	prevHash := hash(block.Block.Header.LastBlockId.Hash)
	hash := hash(block.BlockId.Hash)
	height := block.Block.Header.Height
	if mBlock.Hash != hash && mBlock.PrevHash != prevHash && mBlock.Height != height {
		return fmt.Errorf("database mismatch error")
	}
	return nil
}
func (h handler) HandleBlock(block *tmservice.GetBlockByHeightResponse, txs []*tx.GetTxResponse) error {
	h.log.Info("HandleBlock", "height", block.Block.Header.Height)
	return h.db.Transaction(func(db *gorm.DB) error {
		var blockchain BlockChain
		if err := h.db.FirstOrInit(&blockchain).Error; err != nil {
			return err
		}
		blockchain.BlockCnt += 1

		mBlock := &Block{
			Hash:            hash(block.BlockId.Hash),
			Height:          block.Block.Header.Height,
			PrevHash:        hash(block.Block.Header.LastBlockId.Hash),
			Proposer:        sdk.AccAddress(block.Block.Header.ProposerAddress).String(),
			Time:            block.Block.Header.Time,
			Size:            block.Block.Size(),
			assetMap:        map[string]*Asset{},
			addressMap:      map[string]*Address{},
			contractCodeMap: map[string]*ContractCode{},
			contractMap:     map[string]*Contract{},
			peerMap:         map[string]*Peer{},
			validatorMap:    map[string]*Validator{},
		}

		mTxs := make([]*Transaction, len(txs))
		for itx, tx := range txs {
			h.Logger().Info("HandleBlock", "height", block.Block.Header.Height, "tx hash", tx.TxResponse.TxHash)
			blockchain.TxCnt += 1
			timestamp, _ := time.Parse(tx.TxResponse.Timestamp, time.RFC3339)
			rawTx := encodingConfig.Marshaler.MustMarshalJSON(tx)
			mTx := &Transaction{
				Hash:   tx.TxResponse.TxHash,
				Height: tx.TxResponse.Height,
				Status: tx.TxResponse.Code == 0,
				Memo:   tx.Tx.Body.Memo,
				Fees:   tx.Tx.AuthInfo.Fee.Amount.String(),
				Time:   timestamp,
				Size:   tx.Tx.Size(),
				MsgCnt: len(tx.Tx.GetMsgs()),
				Raw:    string(rawTx),

				assetMap:        map[string]*Asset{},
				addressMap:      map[string]*Address{},
				contractCodeMap: map[string]*ContractCode{},
				contractMap:     map[string]*Contract{},
				peerMap:         map[string]*Peer{},
				validatorMap:    map[string]*Validator{},
			}
			if tx.TxResponse.Code != 0 {
				mTx.RawLog = tx.TxResponse.RawLog
			}
			types := []string{}
			for _, msg := range tx.Tx.GetMsgs() {
				h.Logger().Info("HandleBlock", "height", block.Block.Header.Height, "tx hash", tx.TxResponse.TxHash, "msg", msg.Type())
				if err := h.handleMsg(db, msg, mTx, mBlock); err != nil {
					return err
				}
				types = append(types, msg.Type())
			}
			if len(types) != 0 {
				mTx.Type = strings.Join(types, ",")
			}
			for _, mAddress := range mTx.addressMap {
				mTx.Addresses = append(mTx.Addresses, mAddress)
			}
			for _, mAsset := range mTx.assetMap {
				mTx.Assets = append(mTx.Assets, mAsset)
			}
			for _, mContractCode := range mTx.contractCodeMap {
				mTx.ContractCodes = append(mTx.ContractCodes, mContractCode)
			}
			for _, mContract := range mTx.contractMap {
				mTx.Contracts = append(mTx.Contracts, mContract)
			}
			for _, mPeer := range mTx.peerMap {
				mTx.Peers = append(mTx.Peers, mPeer)
			}
			for _, mValidator := range mTx.validatorMap {
				mTx.Validators = append(mTx.Validators, mValidator)
			}
			mTxs[itx] = mTx

			for _, handleTx := range h.handleTxFuncs {
				if err := handleTx(db, tx); err != nil {
					return err
				}
			}
		}

		mBlock.TxCnt = len(mTxs)
		mBlock.Txs = mTxs

		for _, mAddr := range mBlock.addressMap {
			if mAddr.ID == 0 {
				blockchain.AddressCnt += 1
			}
			mAddr.Balance = mAddr.amounts.String()
		}
		for _, mAsset := range mBlock.assetMap {
			if mAsset.ID == 0 {
				blockchain.AssetCnt += 1
			}
			mAsset.IssueAmount = mAsset.issueAmount.String()
			mAsset.BurnAmount = mAsset.burnAmount.String()
		}
		for _, mContractCode := range mBlock.contractCodeMap {
			if mContractCode.ID == 0 {
				blockchain.ContractCodeCnt += 1
			}
			mContractCode.Approved = strings.Join(mContractCode.approved, ",")
		}
		for _, mContract := range mBlock.contractMap {
			if mContract.ID == 0 {
				blockchain.ContractCnt += 1
			}
		}
		// c := h.client.WithHeight(block.Block.Header.Height)
		//// 更新地址
		//for _, mAddress := range addressMap {
		//	if mAddress.ID == 0 {
		//		mAddress.Height = mBlock.Height
		//		mAddress.Time = mBlock.Time
		//		res, err := c.GetAccount(mAddress.Address)
		//		if err != nil {
		//			return fmt.Errorf("GetAccount %s %s", mAddress.Address, err)
		//		}
		//		var acc authtypes.BaseAccount
		//		if err := proto.Unmarshal(res.Account.Value, &acc); err != nil {
		//			return err
		//		}
		//		mAddress.AccountNumber = acc.AccountNumber
		//		mAddress.Sequence = acc.Sequence
		//	}
		//	res, err := c.GetAccountACL(mAddress.Address)
		//	if err != nil {
		//		return fmt.Errorf("GetAccountACL %s %s", mAddress.Address, err)
		//	}
		//	mAddress.Roles = strings.Join(res.Account.Roles, ",")
		//	res2, err := c.GetAccountAllBalances(mAddress.Address)
		//	if err != nil {
		//		return fmt.Errorf("GetAccountAllBalances %s %s", mAddress.Address, err)
		//	}
		//	mAddress.Balance = res2.Balances.String()
		//}
		//// 更新资产
		//for _, mAsset := range assetMap {
		//	if mAsset.ID == 0 {
		//		mAsset.Height = mBlock.Height
		//		mAsset.Time = mBlock.Time
		//	}
		//	res, err := c.GetToken(mAsset.Name)
		//	if err != nil {
		//		return fmt.Errorf("GetToken %s %s", mAsset.Name, err)
		//	}
		//	mAsset.Creator = res.Token.Creator
		//	mAsset.IssueAmount = res.Token.Issued.String()
		//	mAsset.BurnAmount = res.Token.Burned.String()
		//}
		//// 更新合约代码
		//for _, mContractCode := range contractCodeMap {
		//	if mContractCode.ID == 0 {
		//		mContractCode.Height = mBlock.Height
		//		mContractCode.Time = mBlock.Time
		//	}
		//	res, err := c.GetContractCode(mContractCode.Hash)
		//	if err != nil {
		//		return fmt.Errorf("GetContractCode %s %s", mContractCode.Hash, err)
		//	}
		//	mContractCode.Creator = res.ContractCode.Creator
		//	mContractCode.Name = res.ContractCode.Name
		//	mContractCode.Code = hex.EncodeToString(res.ContractCode.Code)
		//	mContractCode.SRC = res.ContractCode.Src
		//	mContractCode.ABI = res.ContractCode.Abi
		//	mContractCode.Approvers = strings.Join(res.ContractCode.Approvers, ",")
		//	mContractCode.Status = res.ContractCode.Status.String()
		//	mContractCode.Runtime = res.ContractCode.Desc.Runtime
		//	mContractCode.Module = res.ContractCode.Desc.ContractType
		//}
		//// 更新合约
		//for name, mContract := range contractMap {
		//	if mContract.ID == 0 {
		//		mContract.Height = mBlock.Height
		//		mContract.Time = mBlock.Time
		//	}
		//	res, err := c.GetContract(mContract.Name)
		//	if err != nil {
		//		return fmt.Errorf("GetContract %s %s", name, err)
		//	}
		//	mContract.Creator = res.Contract.Creator
		//	mContract.Hash = res.Contract.Hash
		//	mContract.InitArgs = res.Contract.InitArgs
		//	mContract.Status = res.Contract.Status.String()
		//	mContract.Version = res.Contract.Version
		//	if mContract.ID == 0 {
		//		if _, ok := contractCodeMap[mContract.Hash]; !ok {
		//			var cc ContractCode
		//			if err := db.First(&cc, "hash = ?", mContract.Hash).Error; err != nil {
		//				h.log.Error("HandleBlock db.First", "key", mContract.Hash, "error", err)
		//				return err
		//			}
		//			contractCodeMap[mContract.Hash] = &cc
		//		}
		//		mContract.ContractCode = contractCodeMap[mContract.Hash]
		//	}
		//}

		if err := db.Save(&blockchain).Error; err != nil {
			h.log.Error("HandleBlock db.Create", "error", err)
			return err
		}

		if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Create(mBlock).Error; err != nil {
			h.log.Error("HandleBlock db.Create", "error", err)
			return err
		}
		return nil
	})
}

func amountChanged(base *utxotypes.Base) map[string]sdk.Coins {
	changes := make(map[string]sdk.Coins)
	for _, input := range base.Inputs {
		if _, ok := changes[input.Addr]; !ok {
			changes[input.Addr] = sdk.NewCoins()
		}
		changes[input.Addr], _ = changes[input.Addr].SafeSub(sdk.NewCoins(input.Amount))
	}
	for _, output := range base.Outputs {
		if output.Purpose == utxotypes.Output_BURN {
			continue
		}
		if _, ok := changes[output.Addr]; !ok {
			changes[output.Addr] = sdk.NewCoins()
		}
		changes[output.Addr] = changes[output.Addr].Add(output.Amount)
	}
	return changes
}

func (h handler) handleMsg(db *gorm.DB, msg sdk.Msg, mtx *Transaction, mblk *Block) error {
	switch msg := msg.(type) {
	case *accountypes.MsgSetRole:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		maddr, err = getAddress(msg.Address, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr
		maddr.Roles = strings.Join(msg.Roles, ",")
	case *utxotypes.MsgAsset:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		masset, err := getAsset(msg.Metadata.Base, mblk, db)
		if err != nil {
			return err
		}
		mtx.assetMap[masset.Name] = masset
		meta, _ := json.Marshal(msg.Metadata)
		masset.Creator = msg.Creator
		masset.Display = msg.Metadata.Display
		masset.Meta = string(meta)
	case *utxotypes.MsgIssue:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		for _, output := range msg.Base.Outputs {
			if output.Purpose == utxotypes.Output_ISSUE {
				masset, err := getAsset(output.Amount.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
				masset.issueAmount = masset.issueAmount.Add(output.Amount)
			}
		}
	case *utxotypes.MsgTransfer:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr
	case *utxotypes.MsgBurn:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		for _, output := range msg.Base.Outputs {
			if output.Purpose == utxotypes.Output_BURN {
				masset, err := getAsset(output.Amount.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
				masset.burnAmount = masset.burnAmount.Add(output.Amount)
			}
		}
	case *wasmtypes.MsgSubmit:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mContractCode, err := getContractCode(hash(msg.Code), mblk, db)
		if err != nil {
			return err
		}
		mtx.contractCodeMap[mContractCode.Hash] = mContractCode
		mContractCode.Creator = msg.Creator
		mContractCode.Name = msg.Name
		mContractCode.Code = hex.EncodeToString(msg.Code)
		mContractCode.SRC = msg.Src
		mContractCode.ABI = msg.Abi
		mContractCode.Approvers = strings.Join(msg.Approvers, ",")
		if len(msg.Approvers) == 0 {
			mContractCode.Status = wasmtypes.ContractCode_APPROVE.String()
		} else {
			mContractCode.Status = wasmtypes.ContractCode_SUBMIT.String()
		}
		mContractCode.Runtime = msg.Desc.Runtime
		mContractCode.Module = msg.Desc.ContractType
	case *wasmtypes.MsgApprove:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mContractCode, err := getContractCode(msg.CodeHash, mblk, db)
		if err != nil {
			return err
		}
		mtx.contractCodeMap[mContractCode.Hash] = mContractCode
		mContractCode.approved = append(mContractCode.approved, msg.Creator)
		if len(mContractCode.approved) == len(strings.Split(mContractCode.Approvers, ",")) {
			mContractCode.Status = wasmtypes.ContractCode_APPROVE.String()
		} else {
			mContractCode.Status = wasmtypes.ContractCode_SUBMIT.String()
		}
	case *wasmtypes.MsgDeploy:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mContract, err := getContract(msg.Name, mblk, db)
		if err != nil {
			return err
		}
		mContractCode, err := getContractCode(msg.CodeHash, mblk, db)
		if err != nil {
			return err
		}
		mtx.contractMap[mContract.Name] = mContract
		mContract.Name = msg.Name
		mContract.Creator = msg.Creator
		mContract.Hash = msg.CodeHash
		mContract.InitArgs = msg.InitArgs
		mContract.Status = wasmtypes.Contract_NORMAL.String()
		mContract.Version = 1
		mContract.ContractCodes = append(mContract.ContractCodes, mContractCode)
	case *wasmtypes.MsgUpgrade:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mContract, err := getContract(msg.Name, mblk, db)
		if err != nil {
			return err
		}
		mtx.contractMap[mContract.Name] = mContract
		mContract.Version++
		mContractCode, err := getContractCode(msg.CodeHash, mblk, db)
		if err != nil {
			return err
		}
		mContract.ContractCodes = append(mContract.ContractCodes, mContractCode)
	case *wasmtypes.MsgInvoke:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		for _, r := range msg.Requests {
			mContract, err := getContract(r.ContractName, mblk, db)
			if err != nil {
				return err
			}
			mtx.contractMap[mContract.Name] = mContract
		}
	case *wasmtypes.MsgFreeze:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mContract, err := getContract(msg.Name, mblk, db)
		if err != nil {
			return err
		}
		mtx.contractMap[mContract.Name] = mContract
		mContract.Status = wasmtypes.Contract_FROZE.String()
	case *wasmtypes.MsgUnfreeze:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mContract, err := getContract(msg.Name, mblk, db)
		if err != nil {
			return err
		}
		mtx.contractMap[mContract.Name] = mContract
		mContract.Status = wasmtypes.Contract_NORMAL.String()
	case *wasmtypes.MsgDestroy:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mContract, err := getContract(msg.Name, mblk, db)
		if err != nil {
			return err
		}
		mtx.contractMap[mContract.Name] = mContract
		mContract.Status = ""
	case *peertypes.MsgCreatePeer:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mPeer, err := getPeer(msg.Index, mblk, db)
		if err != nil {
			return err
		}
		mtx.peerMap[mPeer.NodeID] = mPeer
		mPeer.Addr = msg.Addr
		mPeer.Moniker = msg.Moniker
	case *peertypes.MsgUpdatePeer:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mPeer, err := getPeer(msg.Index, mblk, db)
		if err != nil {
			return err
		}
		mtx.peerMap[mPeer.NodeID] = mPeer
		mPeer.Addr = msg.Addr
	case *peertypes.MsgDeletePeer:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mPeer, err := getPeer(msg.Index, mblk, db)
		if err != nil {
			return err
		}
		mtx.peerMap[mPeer.NodeID] = mPeer
		mPeer.Delete = true
	case *validatortypes.MsgCreateValidator:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mValidator, err := getValidator(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.validatorMap[mValidator.Creator] = mValidator
		mValidator.PubKey = msg.PubKey
	case *validatortypes.MsgEditValidator:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr
		mValidator, err := getValidator(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.validatorMap[mValidator.Creator] = mValidator
	case *validatortypes.MsgLeaveValidator:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr

		mValidator, err := getValidator(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.validatorMap[mValidator.Creator] = mValidator
		mValidator.Delete = true
	case *storagetypes.MsgCreateData:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr
	case *storagetypes.MsgUpdateData:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr
	case *storagetypes.MsgDeleteData:
		addrChanges := amountChanged(msg.Base)
		for addr, changes := range addrChanges {
			maddr, err := getAddress(addr, mblk, db)
			if err != nil {
				return err
			}
			mtx.addressMap[addr] = maddr
			maddr.amounts = maddr.amounts.Add(changes...)

			for _, change := range changes {
				masset, err := getAsset(change.Denom, mblk, db)
				if err != nil {
					return err
				}
				mtx.assetMap[masset.Name] = masset
			}
		}
		maddr, err := getAddress(msg.Creator, mblk, db)
		if err != nil {
			return err
		}
		mtx.addressMap[maddr.Address] = maddr
	default:
		return fmt.Errorf("not support route %v, type %v", msg.Route(), msg.Type())
	}

	rawMsg := encodingConfig.Marshaler.MustMarshalJSON(msg)
	mMsg := &Msg{
		Hash:  mtx.Hash,
		Route: msg.Route(),
		Type:  msg.Type(),
		Raw:   string(rawMsg),
	}
	mtx.Msgs = append(mtx.Msgs, mMsg)
	return nil
}

func (h handler) Logger() log.Logger {
	if h.log == nil {
		h.log = log.NewTMLogger(os.Stdout)
	}
	return h.log
}

func getAddress(address string, mblk *Block, db *gorm.DB) (*Address, error) {
	if v, ok := mblk.addressMap[address]; ok {
		return v, nil
	}

	var mAddress Address
	if result := db.Find(&mAddress, "address = ?", address); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		mAddress.Address = address
		mAddress.Height = mblk.Height
		mAddress.Time = mblk.Time
		mAddress.amounts = sdk.NewCoins()
	} else {
		var err error
		mAddress.amounts, err = sdk.ParseCoinsNormalized(mAddress.Balance)
		if err != nil {
			return nil, err
		}
	}
	mblk.addressMap[address] = &mAddress
	return &mAddress, nil
}

func getAsset(name string, mblk *Block, db *gorm.DB) (*Asset, error) {
	if v, ok := mblk.assetMap[name]; ok {
		return v, nil
	}

	var mAsset Asset
	if result := db.Find(&mAsset, "name = ?", name); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		mAsset.Name = name
		mAsset.Height = mblk.Height
		mAsset.Time = mblk.Time
		mAsset.IssueAmount = sdk.NewCoin(mAsset.Name, sdk.ZeroInt()).String()
		mAsset.BurnAmount = sdk.NewCoin(mAsset.Name, sdk.ZeroInt()).String()
	}
	var err error
	mAsset.issueAmount, err = sdk.ParseCoinNormalized(mAsset.IssueAmount)
	if err != nil {
		return nil, err
	}
	mAsset.burnAmount, err = sdk.ParseCoinNormalized(mAsset.BurnAmount)
	if err != nil {
		return nil, err
	}
	mblk.assetMap[name] = &mAsset
	return &mAsset, nil

}

func getContract(name string, mblk *Block, db *gorm.DB) (*Contract, error) {
	if v, ok := mblk.contractMap[name]; ok {
		return v, nil
	}

	var mContract Contract
	if result := db.Find(&mContract, "name = ?", name); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		mContract.Name = name
		mContract.Height = mblk.Height
		mContract.Time = mblk.Time
	}
	mblk.contractMap[name] = &mContract
	return &mContract, nil
}

func getContractCode(hash string, mblk *Block, db *gorm.DB) (*ContractCode, error) {
	if v, ok := mblk.contractCodeMap[hash]; ok {
		return v, nil
	}

	var mContractCode ContractCode
	if result := db.Find(&mContractCode, "hash = ?", hash); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		mContractCode.Hash = hash
		mContractCode.Height = mblk.Height
		mContractCode.Time = mblk.Time
	}
	if len(mContractCode.approved) != 0 {
		mContractCode.approved = strings.Split(mContractCode.Approved, ",")
	}
	mblk.contractCodeMap[hash] = &mContractCode
	return &mContractCode, nil
}

func getPeer(name string, mblk *Block, db *gorm.DB) (*Peer, error) {
	if v, ok := mblk.peerMap[name]; ok {
		return v, nil
	}

	var mPeer Peer
	if result := db.Find(&mPeer, "node_id = ?", name); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		mPeer.NodeID = name
		mPeer.Height = mblk.Height
		mPeer.Time = mblk.Time
	}
	mblk.peerMap[name] = &mPeer
	return &mPeer, nil
}

func getValidator(name string, mblk *Block, db *gorm.DB) (*Validator, error) {
	if v, ok := mblk.validatorMap[name]; ok {
		return v, nil
	}

	var mValidator Validator
	if result := db.Find(&mValidator, "creator = ?", name); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		mValidator.Creator = name
		mValidator.Height = mblk.Height
		mValidator.Time = mblk.Time
	}
	mblk.validatorMap[name] = &mValidator
	return &mValidator, nil
}

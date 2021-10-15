package client

import (
	"fmt"
	"testing"
)

func TestGetGenesisTxs(t *testing.T) {
	resp, err := testClient.GetGenesisTxs()
	if err != nil {
		t.Fatal("GetGenesisTxs", err)
	}
	fmt.Println(resp)
}

func TestGetNodeInfo(t *testing.T) {
	resp, err := testClient.GetNodeInfo()
	if err != nil {
		t.Fatal("GetNodeInfo", err)
	}
	fmt.Println(resp)
}

func TestGetSyncing(t *testing.T) {
	resp, err := testClient.GetSyncing()
	if err != nil {
		t.Fatal("GetSyncing", err)
	}
	fmt.Println(resp)
}

func TestGetLatestBlock(t *testing.T) {
	resp, err := testClient.GetLatestBlock()
	if err != nil {
		t.Fatal("GetLatestBlock", err)
	}
	fmt.Println(resp)
}

func TestGetBlockByHeight(t *testing.T) {
	resp, err := testClient.GetBlockByHeight(400000000)
	if err != nil {
		t.Fatal("GetBlockByHeight", err)
	}
	fmt.Println(resp)
}

func TestGetLatestValidatorSet(t *testing.T) {
	resp, err := testClient.GetLatestValidatorSet()
	if err != nil {
		t.Fatal("GetLatestValidatorSet", err)
	}
	bts, err := encodingConfig.Marshaler.MarshalJSON(resp)
	fmt.Println(string(bts), err)
}

func TestGetValidatorSetByHeight(t *testing.T) {
	resp, err := testClient.GetValidatorSetByHeight(4)
	if err != nil {
		t.Fatal("GetValidatorSetByHeight", err)
	}
	bts, err := encodingConfig.Marshaler.MarshalJSON(resp)
	fmt.Println(string(bts), err)
}

func TestGetTx(t *testing.T) {
	resp, err := testClient.GetTx("91FBB9CDB7924B75FDDFD12A22D4936284ACACC6C03BA08E10232D14A2754A9E")
	if err != nil {
		t.Fatal("GetTx", err)
	}
	bts, err := encodingConfig.Marshaler.MarshalJSON(resp)
	fmt.Println(string(bts), err)
}

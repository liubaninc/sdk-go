package client

import (
	"fmt"
	"testing"
)

func TestTxPeerCreate(t *testing.T) {
	builder := &TxPeerCreateBuilder{}
	builder.SetIndex("node-id")
	builder.SetAddr(address1)
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxPeerUpdate(t *testing.T) {
	builder := &TxPeerUpdateBuilder{}
	builder.SetIndex("node-id")
	builder.SetAddr(address1)

	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxPeerDelete(t *testing.T) {
	builder := &TxPeerDeleteBuilder{}
	builder.SetIndex("node-id")

	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

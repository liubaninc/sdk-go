package client

import (
	"fmt"
	"testing"
)

func TestTxAsset(t *testing.T) {
	builder := &TxAssetBuilder{}
	builder.SetAsset("wei", "ueth", 18, "a basic token")
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxIssue(t *testing.T) {
	builder := &TxIssueBuilder{}
	builder.AppendReceiver("mc19dzfuxxv8vjeajjq475ahgrl0meudwexdmrnye", "1000000wei")
	builder.AppendReceiver("mc19dzfuxxv8vjeajjq475ahgrl0meudwexdmrnye", "1000000000wei")
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxTransfer(t *testing.T) {
	builder := &TxTransferBuilder{}
	builder.AppendReceiver("mc1f2zv38nm0nfzz5x08hrjr9wgr6ru9anu4s65l3", "1000000wei")
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxBurn(t *testing.T) {
	builder := &TxBurnBuilder{}
	builder.SetAmounts("40wei")
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

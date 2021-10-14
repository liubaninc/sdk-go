package client

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var codeFile = "/Users/admin/GoPro/m1/x/wasm/xmodel/contractsdk/cpp/build/counter.wasm"

func TestTxSubmit(t *testing.T) {
	builder := &TxSubmitBuilder{}
	code, err := ioutil.ReadFile(codeFile)
	if err != nil {
		t.Fatalf("invalid code file (%v)/n", err)
	}
	builder.SetCode(code)
	//builder.SetAbi("")
	src, err := ioutil.ReadFile("/Users/admin/GoPro/m1/x/wasm/xmodel/contractsdk/cpp/example/counter.cc")
	if err != nil {
		t.Fatalf("invalid src file (%v)/n", err)
	}
	builder.SetSrc(string(src))
	builder.SetName("计数器")
	builder.SetModule("wasm")
	builder.SetRuntime("c")
	builder.SetApprovers([]string{address2})

	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxApprove(t *testing.T) {
	builder := &TxApproveBuilder{}
	code, err := ioutil.ReadFile(codeFile)
	if err != nil {
		t.Fatalf("invalid code file (%v)", err)
	}
	builder.SetCode(code)
	resp, err := testClient.GenerateAndBroadcastTx(privateKey1, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxDeploy(t *testing.T) {
	builder := &TxDeployBuilder{}
	code, err := ioutil.ReadFile(codeFile)
	if err != nil {
		t.Fatalf("invalid code file (%v)", err)
	}
	builder.SetName("ccc1")
	builder.SetCode(code)
	builder.SetInitArgs("{}")
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxUpgrade(t *testing.T) {
	builder := &TxUpgradeBuilder{}
	code, err := ioutil.ReadFile(codeFile)
	if err != nil {
		t.Fatalf("invalid code file (%v)", err)
	}
	builder.SetName("ccc1")
	builder.SetCode(code)
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxInvoke(t *testing.T) {
	builder := &TxInvokeBuilder{}
	builder.SetName("ccc1")
	builder.SetMethod("increase")
	builder.SetArgs("{\"key\":\"a\"}")
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxFreeze(t *testing.T) {
	builder := &TxFreezeBuilder{}
	builder.SetName("ccc1")
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxUnFreeze(t *testing.T) {
	builder := &TxUnFreezeBuilder{}
	builder.SetName("ccc1")
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestTxDestroy(t *testing.T) {
	builder := &TxDestroyBuilder{}
	builder.SetName("ccc1")
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

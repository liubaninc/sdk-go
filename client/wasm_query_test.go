package client

import (
	"fmt"
	"testing"
)

func TestGetContract(t *testing.T) {
	resp, err := testClient.GetContract("ccc1")
	if err != nil {
		t.Fatal("GetContract", err)
	}
	fmt.Println(resp)
}

func TestGetContractAll(t *testing.T) {
	resp, err := testClient.GetContractAll(nil, 0, 1, true)
	if err != nil {
		t.Fatal("GetContractAll", err)
	}
	fmt.Println(resp)
}

func TestGetContractCode(t *testing.T) {
	resp, err := testClient.GetContractCode("6D9AD0999683B3BA7FEA98B3C33691D7CEC8AEE535AEAB2BE3A2B66AC527A9AF")
	if err != nil {
		t.Fatal("GetContractCode", err)
	}
	fmt.Println(resp)
}

func TestGetContractCodeAll(t *testing.T) {
	resp, err := testClient.GetContractCodeAll(nil, 0, 10, true)
	if err != nil {
		t.Fatal("GetContractCodeAll", err)
	}
	fmt.Println(resp)
}

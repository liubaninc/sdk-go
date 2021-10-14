package client

import (
	"fmt"
	"testing"
)

func TestGetTokenAll(t *testing.T) {
	resp, err := testClient.GetTokenAll(nil, 0, 10, true)
	if err != nil {
		t.Fatal("GetTokenAll", err)
	}
	fmt.Println(resp)
}

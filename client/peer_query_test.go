package client

import (
	"fmt"
	"testing"
)

func TestGetPeer(t *testing.T) {
	resp, err := testClient.GetPeer("node-id")
	if err != nil {
		t.Fatal("GetPeer", err)
	}
	fmt.Println(resp)
}

func TestGetPeerAll(t *testing.T) {
	resp, err := testClient.GetPeerAll(nil, 0, 10, true)
	if err != nil {
		t.Fatal("GetPeerAll", err)
	}
	fmt.Println(resp)
}

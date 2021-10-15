package client

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	testClient  = MustNew("http://localhost:26657")
	mnemonic    = "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear"
	mnemonic1   = "chapter useful holiday pumpkin that game lion vendor glove engage clarify old palm bubble august plastic either song trouble guilt gasp sea make deal"
	address     = "mc19dzfuxxv8vjeajjq475ahgrl0meudwexdmrnye"
	address1    = "mc1f2zv38nm0nfzz5x08hrjr9wgr6ru9anu4s65l3"
	address2    = "mc1kczun6k6nt4p8fyvaxz5h3jfsyesuuead084eq"
	privateKey  cryptotypes.PrivKey
	privateKey1 cryptotypes.PrivKey
)

func init() {
	kr := keyring.NewInMemory()

	_, err := kr.NewAccount("test", mnemonic, keyring.DefaultBIP39Passphrase, sdk.GetConfig().GetFullFundraiserPath(), hd.Secp256k1)
	if err != nil {
		panic(err)
	}
	_, err = kr.NewAccount("wang", mnemonic1, keyring.DefaultBIP39Passphrase, sdk.GetConfig().GetFullFundraiserPath(), hd.Secp256k1)
	if err != nil {
		panic(err)
	}
	privateKeyArmor, err := kr.ExportPrivKeyArmor("test", "111")
	if err != nil {
		panic(err)
	}
	privateKey, _, err = crypto.UnarmorDecryptPrivKey(privateKeyArmor, "111")
	privateKeyArmor1, err := kr.ExportPrivKeyArmor("wang", "111")
	if err != nil {
		panic(err)
	}
	privateKey1, _, err = crypto.UnarmorDecryptPrivKey(privateKeyArmor1, "111")
}

func TestSetRole(t *testing.T) {
	builder := &SetRoleBuilder{}
	builder.SetAddress("mc1f2zv38nm0nfzz5x08hrjr9wgr6ru9anu4s65l3")
	builder.SetRoles([]string{"admin"})
	resp, err := testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
	builder = &SetRoleBuilder{}
	builder.SetAddress("mc1kczun6k6nt4p8fyvaxz5h3jfsyesuuead084eq")
	builder.SetRoles([]string{"admin"})
	resp, err = testClient.GenerateAndBroadcastTx(privateKey, builder, "block")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", err)
	}
	fmt.Println(resp)
}

func TestGetAccountAll(t *testing.T) {
	resp, err := testClient.GetAccountACLAll(nil, 0, 10, true)
	if err != nil {
		t.Fatal("GetAccountAll", err)
	}
	fmt.Println(resp)
}

func TestGetAccount(t *testing.T) {
	resp, err := testClient.GetAccountACL("mc1f2zv38nm0nfzz5x08hrjr9wgr6ru9anu4s65l3")
	if err != nil {
		t.Fatal("GetAccount", err)
	}
	fmt.Println(resp)
}

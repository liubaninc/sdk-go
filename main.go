package main

import (
	"context"
	"fmt"
	"github.com/liubaninc/sdk-go/client"
	"github.com/liubaninc/sdk-go/syncer"
	"github.com/tendermint/tendermint/libs/log"
	"os"
)

func main() {
	client := client.MustNew("http://localhost:26657")

	db, err := syncer.NewDB("sqlite", "synced.db")
	if err != nil {
		panic(fmt.Errorf("failed to open database"))
	}

	syncer.AutoMigrate(db)

	syncer := syncer.NewHander(log.NewTMLogger(os.Stdout), db, client)
	start := syncer.GetLastBlockHeight() + 1

	client.OffChainSyncing(context.Background(), start, syncer)

}

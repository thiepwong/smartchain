package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/thiepwong/smartchain/core/types"
)

func main() {

	_h := &types.Header{}
	_h.Difficulty = 5
	_h.Height = 0
	_h.Miner = types.Address{}
	_h.Nonce = types.BlockNonce{}
	_h.Time = 12313123123
	_h.Version = 3

	_ts := &types.Transactions{}

	bl := types.NewBlock(_h, *_ts)

	be, e := json.Marshal(bl)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(*_h, *bl, string(be))
}

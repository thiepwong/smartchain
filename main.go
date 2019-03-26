package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/thiepwong/smartchain/core/types"
)

func main() {

	_h := &types.Header{}
	_h.Difficulty = big.NewInt(5)
	_h.Height = big.NewInt(100)
	_h.Miner = types.Address{}
	_h.Nonce = types.BlockNonce{}
	_h.Time = big.NewInt(time.Now().Unix())
	_h.Version = 3

	_ts := &types.Transactions{types.Transaction{byte{0xaa, 0x22, 0x44}}}

	bl := types.NewBlock(_h, *_ts)

	be, e := json.Marshal(bl)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(*_h, *bl, string(be))
}

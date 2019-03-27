package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/thiepwong/smartchain/core"

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

	_ts := &types.Transactions{&types.Transaction{}}

	bl := types.NewBlock(_h, *_ts)

	be, e := json.Marshal(bl)
	if e != nil {
		log.Fatal(e)
	}

	// db, e := leveldb.New("local-data", 4000, 4000, "thongbao")
	// if e != nil {
	// 	fmt.Println(e.Error())
	// 	os.Exit(3)
	// }
	// cf := &params.ChainConfig{ChainID: 3}
	// bc, e := core.NewBlockChain(db, *cf, bl)

	// //bc1, e := core.GetlastBlock()
	// if e != nil {
	// 	fmt.Println(e)
	// 	os.Exit(10)
	// }
	bc1, e := core.GetLocalChain()
	fmt.Println("Loi goi block ", e.Error())
	//pc, e := bc.PullChain()
	//fmt.Printf("day la pull ve :%x", string(pc))

	b, e := core.GetlastBlock(bc1)

	fmt.Println("du lieu: ", bc1, b)
	fmt.Println(*_h, *bl, string(be))
}

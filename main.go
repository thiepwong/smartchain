package main

import (
	"fmt"
	"os"

	"github.com/thiepwong/smartchain/smartdb/mongodb"
)

func main() {

	// _h := &types.Header{}
	// _h.Difficulty = big.NewInt(5)
	// _h.Height = big.NewInt(100)

	// _h.Miner = types.Address{}
	// _h.Nonce = types.BlockNonce{}
	// _h.Time = big.NewInt(time.Now().Unix())
	// _h.Version = 3

	// data := &types.TxData{}
	// data.Mode = 4
	// data.P = []byte{0x22, 0x55}
	// data.SmartID = 99883728288393
	// data.User = []byte{0x55, 0x77, 0x11}

	// tx, e := types.NewTransaction(data)

	// _ts := &types.Transactions{tx}

	// bl := types.NewBlock(_h, *_ts)

	// be, e := json.Marshal(bl)
	// if e != nil {
	// 	log.Fatal(e)
	// }
	db, e := mongodb.New("mongodb://171.244.49.164:2688", "mainchain")
	//conf := &params.ChainConfig{}
	// conf.ChainID = 44
	// bc, e := core.NewBlockChain(db, *conf, bl)
	if e != nil {
		fmt.Println(e)
		os.Exit(10)
	}

	zz := db.Load()

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

	fmt.Println("Da tao block ", zz)
	//pc, e := bc.PullChain()
	//fmt.Printf("day la pull ve :%x", string(pc))

	//	fmt.Println("du lieu: ", bc1, b)
	//	fmt.Println(*_h, *bl, string(be))
}

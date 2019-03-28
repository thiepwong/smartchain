package main

import (
	"fmt"
	"log"
	"time"

	"github.com/thiepwong/smartchain/core"
	"github.com/thiepwong/smartchain/core/params"
	"github.com/thiepwong/smartchain/core/types"

	"github.com/thiepwong/smartchain/smartdb/mongodb"
)

func main() {

	_h := &types.Header{}
	_dif := new(uint64)
	_heg := new(uint64)
	*_heg = 23258
	*_dif = 444
	_h.Difficulty = _dif
	_h.Height = _heg
	_h.ParentHash = []byte{0x05, 0x1f, 0x0f, 0x0c, 0x3b}
	_h.Miner = types.Address{}
	_h.Nonce = types.BlockNonce{}
	_ti := new(uint64)
	*_ti = uint64(time.Now().Unix())
	fmt.Println("Thoi gian: ", *_ti)
	_h.Time = _ti
	_h.Version = 3
	//_a := [types.HashLength]byte{0x9f, 0x86, 0xd0, 0x81, 0x88, 0x4c, 0x7d, 0x65, 0x9a, 0x2f, 0xea, 0xa0, 0xc5, 0x5a, 0xd0, 0x15, 0xa3, 0xbf, 0x4f, 0x1b, 0x2b, 0x0b, 0x82, 0x2c, 0xd1, 0x5d, 0x6c, 0x15, 0xb0, 0xf0, 0x0a, 0x08}

	data := &types.TxData{}
	data.Mode = 4
	data.P = []byte{0x22, 0x55}
	data.SmartID = 99883728288393
	data.User = []byte{0x55, 0x77, 0x11}

	tx, e := types.NewTransaction(data)

	//	_ts := &types.Transactions{tx}

	bl := types.NewBlock(_h, []types.Transaction{*tx})

	//	be, e := json.Marshal(bl)
	if e != nil {
		log.Fatal(e)
	}
	db, e := mongodb.New("mongodb://171.244.49.164:2688", "mainchain")
	conf := &params.ChainConfig{}
	conf.ChainID = 44
	//bc, e := core.NewBlockChain(db, *conf, bl)
	// if e != nil {
	// 	fmt.Println(e)
	// 	os.Exit(10)
	// }

	trBL := bl.String()

	fmt.Println(trBL)
	fmt.Println("===============")
	//	fmt.Println(bc)

	core.GetLastBlock(db)
	//	zz := db.Load()
	// fmt.Println("===============")
	// fmt.Println(*bll)

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

	//	fmt.Println("Da tao block ", bc)
	//pc, e := bc.PullChain()
	//fmt.Printf("day la pull ve :%x", string(pc))

	//	fmt.Println("du lieu: ", bc1, b)
	//	fmt.Println(*_h, *bl, string(be))
}

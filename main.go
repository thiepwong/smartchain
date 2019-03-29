package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/thiepwong/smartchain/core"
	"github.com/thiepwong/smartchain/core/common"
	"github.com/thiepwong/smartchain/core/params"
	"github.com/thiepwong/smartchain/core/types"

	"github.com/thiepwong/smartchain/smartdb/mongodb"
)

func main() {
	interval := int64(10)
	for {
		time := int64(time.Now().Unix())
		if time%interval == 0 {
			Loop()
		}
	}

}

func Loop() {

	_h := &types.Header{
		Version:    1000,
		Difficulty: big.NewInt(10000),
		Extra:      []byte{0x05, 0xff, 0x35, 0x77},
		Height:     big.NewInt(int64(time.Now().Unix())),
		MerkleHash: common.Hash{0x33, 0xf0, 0xa5, 0x66},
		Miner:      common.Address{0xff, 0xfa, 0xfc},
		Nonce:      types.BlockNonce{0xfc, 0xfb, 0xfa},
		ParentHash: common.Hash{0xcf, 0xcc, 0xca},
		Time:       big.NewInt(int64(time.Now().Unix())),
		TxHash:     common.Hash{0x03, 0x04, 0x10, 0x85},
	}

	//_a := [types.HashLength]byte{0x9f, 0x86, 0xd0, 0x81, 0x88, 0x4c, 0x7d, 0x65, 0x9a, 0x2f, 0xea, 0xa0, 0xc5, 0x5a, 0xd0, 0x15, 0xa3, 0xbf, 0x4f, 0x1b, 0x2b, 0x0b, 0x82, 0x2c, 0xd1, 0x5d, 0x6c, 0x15, 0xb0, 0xf0, 0x0a, 0x08}

	//	ld, e := types.DeserializeHeader(_h)
	fmt.Println()

	data := &types.TxData{}
	data.Mode = 4
	data.P = []byte{0x22, 0x55}
	data.SmartID = 99883728288393
	data.User = []byte{0x55, 0x77, 0x11}

	tx, e := types.NewTransaction(data, []byte{0x00, 0x01, 0x05})
	fmt.Println("=============== Chu ky so ===============")
	key, _ := crypto.GenerateKey()
	types.TxSign(tx, key)
	//	_ts := &types.Transactions{tx}

	bl := types.NewBlock(_h, []types.Transaction{*tx})

	//	be, e := json.Marshal(bl)
	if e != nil {
		log.Fatal(e)
	}
	db, e := mongodb.New("mongodb://171.244.49.164:2688", "mainchain")
	conf := &params.ChainConfig{}
	conf.ChainID = 44
	core.NewBlockChain(db, *conf, bl)
	if e != nil {
		fmt.Println(e)
		os.Exit(10)
	}
	fmt.Println("===============")

	fmt.Printf("Xem lai json: %s", string(types.ToBytes(core.GetLastBlock(db))))
}

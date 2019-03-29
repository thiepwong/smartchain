package main

import (
	"fmt"
	"math/big"
	"reflect"
	"time"
)

type Header struct {
	Version    int      `json:"version"`
	ParentHash [32]byte `json:"parentHash"`
	MerkleHash [32]byte `json:"merkleHash"`
	TxHash     [32]byte `json:"txHash"`
	Height     *big.Int `json:"height"`
	Miner      [20]byte `json:"miner"`
	Time       *big.Int `json:"timestamp"`
	Difficulty *big.Int `json:"difficulty"`
	Nonce      [8]byte  `json:"nonce"`
	Extra      []byte   `json:"extra"`
}

func main() {
	fmt.Println("Hello, playground")

	var headerSize = uint64(reflect.TypeOf(Header{}).Size())
	// _h := new(uint64)
	// *_h = 3455
	// _t := new(uint64)
	// *_t = uint64(time.Now().Unix())

	h := &Header{
		Version:    1000,
		Difficulty: big.NewInt(10000),
		Extra:      []byte{0x05, 0xff, 0x35, 0x77},
		Height:     big.NewInt(102001),
		MerkleHash: [32]byte{0x33, 0xf0, 0xa5, 0x66},
		Miner:      [20]byte{0xff, 0xfa, 0xfc},
		Nonce:      [8]byte{0xfc, 0xfb, 0xfa},
		ParentHash: [32]byte{0xcf, 0xcc, 0xca},
		Time:       big.NewInt(int64(time.Now().Unix())),
		TxHash:     [32]byte{0x03, 0x04, 0x10, 0x85},
	}

	headerSize += uint64(len(h.Extra) + (h.Difficulty.BitLen()+h.Height.BitLen()+h.Time.BitLen())/8)

	//	uPnt := (unsafe.Sizeof(h.Difficulty) + unsafe.Sizeof(h.Height) + unsafe.Sizeof(h.Time))
	//	af := headerSize + uint64(len(h.Extra)+int(uPnt))/8
	//	fmt.Println(len(h.Extra))
	//	fmt.Println(int(uPnt))
	//	fmt.Println("=========== \r\n ", uint64((h.Height.BitLen())/8))
	fmt.Println(headerSize)

	//	fmt.Print("================\r\n", len(h.Extra), float64((h.Difficulty.BitLen()+h.Height.BitLen()+h.Time.BitLen())/8))

}

package types

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	// HashLength is the expected length of the hash
	HashLength = 32
	// AddressLength is the expected length of the address
	AddressLength = 20
	//NonceLength is length of nonce
	NonceLength = 8
)

//BlockNonce is 64 bit hash
type BlockNonce [NonceLength]byte

//Hash is encode of an object
type Hash []byte

//Address is wallet of smartchain
type Address []byte

//EncodeNonce conver the given integer to hash
func EncodeNonce(i uint64) BlockNonce {
	var n BlockNonce
	binary.BigEndian.PutUint64(n[:], i)
	return n
}

// Uint64 return integer value of block nonce
func (n BlockNonce) Uint64() uint64 {
	return binary.BigEndian.Uint64(n[:])
}

//MarshalText encodes n as a hex string with 0x prefix
func (n BlockNonce) MarshalText() ([]byte, error) {
	return hexutil.Bytes(n[:]).MarshalText()
}

//UnMarshalText function implements
func (n BlockNonce) UnMarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("BlockNonce", input, n[:])
}

// Header represents a block header in smartchain
type Header struct {
	Version    int        `json:"version"`
	ParentHash Hash       `json:"parentHash"`
	MerkleHash Hash       `json:"merkleHash"`
	TxHash     Hash       `json:"txHash"`
	Height     uint64     `json:"height"`
	Miner      Address    `json:"miner"`
	Time       uint64     `json:"timestamp"`
	Difficulty int        `json:"difficulty"`
	Nonce      BlockNonce `json:"nonce"`
	Extra      []byte     `json:"extra"`
}

//Block structure of chain
type Block struct {
	Header       *Header
	Transactions Transactions
	Hash         [HashLength]byte
	Size         int
	ReceiverAt   time.Time
	ReceiverFrom interface{}
}

func NewBlock(header *Header, txs Transactions) *Block {
	b := &Block{Header: CopyHeader(header)}

	if len(txs) == 0 {
		b.Header.TxHash = Hash{}
	} else {
		_txs, err := json.Marshal(txs)
		if err != nil {
			log.Fatal(err)
		}
		b.Header.TxHash = _txs
		b.Transactions = make(Transactions, len(txs))
		b.ReceiverAt = time.Now()
		copy(b.Transactions, txs)
	}

	return b

}

// CopyHeader creates a deep copy of a block header to prevent side effects from
// modifying a header variable.
func CopyHeader(h *Header) *Header {
	cpy := *h
	if cpy.Time = *new(uint64); h.Time > 0 {
		cpy.Time = h.Time
	}
	if cpy.Difficulty = *new(int); h.Difficulty > 0 {
		cpy.Difficulty = h.Difficulty
	}
	if cpy.Height = *new(uint64); h.Height > 0 {
		cpy.Height = h.Height
	}
	if len(h.Extra) > 0 {
		cpy.Extra = make([]byte, len(h.Extra))
		copy(cpy.Extra, h.Extra)
	}
	return &cpy
}

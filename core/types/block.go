package types

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
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
type Hash [HashLength]byte

//Address is wallet of smartchain
type Address [AddressLength]byte

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

var headerSize = float64(reflect.TypeOf(Header{}).Size())

// Header represents a block header in smartchain
type Header struct {
	Version    int        `json:"version"`
	ParentHash Hash       `json:"parentHash"`
	MerkleHash Hash       `json:"merkleHash"`
	TxHash     Hash       `json:"txHash"`
	Height     *big.Int   `json:"height"`
	Miner      Address    `json:"miner"`
	Time       *big.Int   `json:"timestamp"`
	Difficulty *big.Int   `json:"difficulty"`
	Nonce      BlockNonce `json:"nonce"`
	Extra      []byte     `json:"extra"`
}

//Block structure of chain
type Block struct {
	Header       *Header
	Transactions Transactions
	Hash         Hash    //Hash of Header
	Size         float64 // Size of header
	ReceiverAt   *big.Int
	ReceiverFrom interface{}
}

func (h *Header) Size() float64 {
	return headerSize + float64(len(h.Extra)+(h.Difficulty.BitLen()+h.Height.BitLen()+h.Time.BitLen())/8)
}

//Hash of Header
func (h *Header) Hash() (hash Hash) {
	b, err := json.Marshal(h)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return hash
	}
	_hash := sha256.New()
	_hash.Write(b)
	_hash.Sum(hash[:0])
	return hash
}

//Hash of block
// func (b *Block) Hash() (hash Hash) {
// 	b.Hash

// }
func NewBlock(header *Header, txs Transactions) *Block {
	b := &Block{Header: CopyHeader(header)}

	if len(txs) == 0 {
		b.Header.TxHash = Hash{}
	} else {

		b.Transactions = make(Transactions, len(txs))

		fmt.Println()
		b.ReceiverAt = big.NewInt(time.Now().Unix())
		copy(b.Transactions, txs)
	}
	b.Hash = b.Header.Hash()
	b.Size = b.Header.Size()
	return b

}

// CopyHeader creates a deep copy of a block header to prevent side effects from
// modifying a header variable.
func CopyHeader(h *Header) *Header {
	cpy := *h
	if cpy.Time = new(big.Int); h.Time != nil {
		cpy.Time = h.Time
	}
	if cpy.Difficulty = new(big.Int); h.Difficulty != nil {
		cpy.Difficulty = h.Difficulty
	}
	if cpy.Height = new(big.Int); h.Height != nil {
		cpy.Height = h.Height
	}
	if len(h.Extra) > 0 {
		cpy.Extra = make([]byte, len(h.Extra))
		copy(cpy.Extra, h.Extra)
	}
	return &cpy
}

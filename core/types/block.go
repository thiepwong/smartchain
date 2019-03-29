package types

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/thiepwong/smartchain/core/common"
)

const (
	//NonceLength is length of nonce
	NonceLength = 8
)

//BlockNonce is 64 bit hash
type BlockNonce [NonceLength]byte

//Hash is encode of an object
//type Hash [HashLength]byte

var headerSize = common.StorageSize(reflect.TypeOf(Header{}).Size())

type writeCounter common.StorageSize

// Header represents a block header in smartchain
type Header struct {
	Version    int            `json:"version"`
	ParentHash common.Hash    `json:"parentHash"`
	MerkleHash common.Hash    `json:"merkleHash"`
	TxHash     common.Hash    `json:"txHash"`
	Height     *big.Int       `json:"height"`
	Miner      common.Address `json:"miner"`
	Time       *big.Int       `json:"timestamp"`
	Difficulty *big.Int       `json:"difficulty"`
	Nonce      BlockNonce     `json:"nonce"`
	Extra      []byte         `json:"extra"`
}

//Block structure of chain
type Block struct {
	Heigth       uint64        `bson:"_id,omitempty"`
	Header       *Header       `json:"header"`
	Transactions []Transaction `json:"transactions"`
	Hash         common.Hash   `json:"hash"` //Hash of Header
	Size         atomic.Value  `json:"size"` // Size of header
	ReceiverAt   uint64        `bson:"received_at" json:"received_at"`
	ReceiverFrom interface{}   `json:"received_from"`
}

//EncodeNonce conver the given integer to hash
func EncodeNonce(i uint64) BlockNonce {
	var n BlockNonce
	binary.BigEndian.PutUint64(n[:], i)
	return n
}

//EncodeHash for encoding
func EncodeHash(i uint64) common.Hash {
	var n common.Hash
	binary.BigEndian.PutUint64(n[:], i)
	return n
}

// Uint64 return integer value of block nonce
// func (h common.Hash) Uint64() uint64 {
// 	return binary.BigEndian.Uint64(h[:])
// }

//MarshalText encodes n as a hex string with 0x prefix
// func (h Hash) MarshalText() ([]byte, error) {
// 	return hexutil.Bytes(h[:]).MarshalText()
// }

// //UnMarshalText function implements
// func (h Hash) UnMarshalText(input []byte) error {
// 	return hexutil.UnmarshalFixedText("Hash", input, h[:])
// }

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

//Size func to get size of header
func (h *Header) Size() common.StorageSize {
	uPnt := (unsafe.Sizeof(h.Difficulty) + unsafe.Sizeof(h.Height) + unsafe.Sizeof(h.Time))
	return headerSize + common.StorageSize(len(h.Extra)+int(uPnt))/8
}

//Hash of Header
func (h *Header) setHash() (hash common.Hash) {
	_byte, err := h.Serialize()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return hash
	}
	_hash := sha256.New()
	_hash.Write(_byte)
	_hash.Sum(hash[:0])
	return hash
}

func (b *Block) setHash() error {
	_byte, err := b.serialize()
	hash := sha256.Sum256(_byte)
	b.Hash = hash
	return err
}

//NewBlock for blockchain
func NewBlock(header *Header, txs []Transaction) *Block {
	b := &Block{Header: CopyHeader(header)}

	if len(txs) == 0 {
		b.Header.TxHash = common.Hash{}
	} else {

		b.Transactions = txs // = make(Transactions, len(txs))

		fmt.Println()
		b.ReceiverAt = uint64(time.Now().Unix())
		copy(b.Transactions, txs)
	}
	b.Heigth = b.Header.Height.Uint64()
	b.Hash = b.Header.setHash()
	//	b.Size = b.Header.Size()
	_nonce := BlockNonce{5, 2, 3, 5, 6, 8, 6}

	b.Header.Nonce = _nonce
	b.setHash()
	return b

}

// CopyHeader creates a deep copy of a block header to prevent side effects from
// modifying a header variable.
func CopyHeader(h *Header) *Header {
	cpy := *h
	if cpy.Time = new(big.Int); h.Time != nil {
		cpy.Time.Set(h.Time)
	}
	if cpy.Difficulty = new(big.Int); h.Difficulty != nil {
		cpy.Difficulty.Set(h.Difficulty)
	}
	if cpy.Height = new(big.Int); h.Height != nil {
		cpy.Height.Set(h.Height)
	}
	if len(h.Extra) > 0 {
		cpy.Extra = make([]byte, len(h.Extra))
		copy(cpy.Extra, h.Extra)
	}
	return &cpy
}

//Size of block
func (b *Block) GetSize() common.StorageSize {
	if size := b.Size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, b)
	b.Size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

//Serialize the block
func (b *Block) serialize() ([]byte, error) {
	return json.Marshal(b)
}

//String convert a block to string json
func (b Block) String() string {
	// var strBlock string
	// strBlock += fmt.Sprintf("Prev hash: %x\n", b.Header.ParentHash)
	// strBlock += fmt.Sprintf("Transactions: \n")
	// for idx, tx := range b.Transactions {
	// 	strBlock += fmt.Sprintf("  Tx[%d] : %x\n", idx, tx)
	// }
	// strBlock += fmt.Sprintf("Hash: %x\n", b.Header.Hash)
	// strBlock += fmt.Sprintf("Hash Block: %x\n", b.Hash)
	// strBlock += fmt.Sprintf("Nonce: %d\n", b.Header.Nonce)
	// strBlock += fmt.Sprintf("Height: %d\n", *b.Header.Height)
	// strBlock += fmt.Sprintf("Timestamp: %d\n", *b.Header.Time)

	_byte, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(_byte))

	// return strBlock
	return string(b.Hash[:])
}

func (h Header) String() string {
	_byte, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(_byte))

	// return strBlock
	return string(h.TxHash[:])
}

//DeserializeBlock func for de-serialize the block info
func DeserializeBlock(data []byte) (*Block, error) {
	b := new(Block)
	err := json.Unmarshal(data, b)
	return b, err
}

//Serialize to serialize the header of block
func (h *Header) Serialize() ([]byte, error) {
	return json.Marshal(h)

}

//DeserializeHeader to to deserialize the header of the header
func DeserializeHeader(data []byte) (*Header, error) {
	h := new(Header)
	err := json.Unmarshal(data, h)
	return h, err
}

func (c *writeCounter) Write(b []byte) (int, error) {
	*c += writeCounter(len(b))
	return len(b), nil
}

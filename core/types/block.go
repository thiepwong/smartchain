package types

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"
	"unsafe"

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
type BlockNonce []byte

//Hash is encode of an object
type Hash []byte

//Address is wallet of smartchain
type Address [AddressLength]byte

//EncodeNonce conver the given integer to hash
func EncodeNonce(i uint64) BlockNonce {
	var n BlockNonce
	binary.BigEndian.PutUint64(n[:], i)
	return n
}

//EncodeHash for encoding
func EncodeHash(i uint64) Hash {
	var n Hash
	binary.BigEndian.PutUint64(n[:], i)
	return n
}

// Uint64 return integer value of block nonce
func (h Hash) Uint64() uint64 {
	return binary.BigEndian.Uint64(h[:])
}

//MarshalText encodes n as a hex string with 0x prefix
func (h Hash) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

//UnMarshalText function implements
func (h Hash) UnMarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("Hash", input, h[:])
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
	Hash       Hash       `json:"hash"`
	TxHash     Hash       `json:"txHash"`
	Height     *uint64    `json:"height"`
	Miner      Address    `json:"miner"`
	Time       *uint64    `json:"timestamp"`
	Difficulty *uint64    `json:"difficulty"`
	Nonce      BlockNonce `json:"nonce"`
	Extra      []byte     `json:"extra"`
}

//Block structure of chain
type Block struct {
	Heigth       uint64        `bson:"_id,omitempty"`
	Header       *Header       `json:"header"`
	Transactions []Transaction `json:"transactions"`
	Hash         Hash          `json:"hash"` //Hash of Header
	Size         float64       `json:"size"` // Size of header
	ReceiverAt   uint64        `bson:"received_at" json:"received_at"`
	ReceiverFrom interface{}   `json:"received_from"`
}

//Size func to get size of header
func (h *Header) Size() float64 {
	uPnt := (unsafe.Sizeof(h.Difficulty) + unsafe.Sizeof(h.Height) + unsafe.Sizeof(h.Time))

	return headerSize + float64(len(h.Extra)+int(uPnt))/8
}

//Hash of Header
func (h *Header) setHash() (hash Hash) {
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

func (b *Block) setHash() {
	hash := sha256.Sum256(b.Serialize())

	b.Header.Hash = hash[:]
	b.Hash = hash[:]
	fmt.Println("Hash cua block la: ", b.Header.Hash)
}

//Hash of block
// func (b *Block) Hash() (hash Hash) {
// 	b.Hash

// }
func NewBlock(header *Header, txs []Transaction) *Block {
	b := &Block{Header: CopyHeader(header)}

	if len(txs) == 0 {
		b.Header.TxHash = []byte{}
	} else {

		b.Transactions = txs // = make(Transactions, len(txs))

		fmt.Println()
		b.ReceiverAt = uint64(time.Now().Unix())
		copy(b.Transactions, txs)
	}
	b.Heigth = *b.Header.Height
	b.Hash = b.Header.setHash()
	b.Size = b.Header.Size()
	_nonce := &[]byte{5, 2, 3, 5, 6, 8, 6}

	b.Header.Nonce = *_nonce
	b.setHash()
	return b

}

// CopyHeader creates a deep copy of a block header to prevent side effects from
// modifying a header variable.
func CopyHeader(h *Header) *Header {
	cpy := *h
	if cpy.Time = new(uint64); h.Time != nil {
		cpy.Time = h.Time
	}
	if cpy.Difficulty = new(uint64); h.Difficulty != nil {
		cpy.Difficulty = h.Difficulty
	}
	if cpy.Height = new(uint64); h.Height != nil {
		cpy.Height = h.Height
	}
	if len(h.Extra) > 0 {
		cpy.Extra = make([]byte, len(h.Extra))
		copy(cpy.Extra, h.Extra)
	}
	return &cpy
}

func (b *Block) Serialize() []byte {
	data, err := json.Marshal(b)

	if err != nil {
		fmt.Printf("Marshal block fail\n")
		os.Exit(1)
	}
	return data
}

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

//DeserializeBlock func for de-serialize the block info
func DeserializeBlock(data []byte) *Block {
	b := new(Block)
	err := json.Unmarshal(data, b)

	if err != nil {
		fmt.Printf("Marshal block fail\n")
		os.Exit(1)
	}

	return b
}

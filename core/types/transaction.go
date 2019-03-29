package types

import (
	"crypto/sha256"
	"encoding/json"
	"unsafe"

	"github.com/thiepwong/smartchain/core/common"
)

//Transactions type
type Transactions []*Transaction

//Transaction  type
type Transaction struct {
	Hash common.Hash `json:"hash" bson:"hash"`
	Size int         `json:"size"`
	From []byte      `json:"from"`
	Data TxData      `json:"data"`
}

//TxData  for transaction
type TxData struct {
	//is register new or update data
	Mode    int    `json:"mode"`
	SmartID uint64 `json:"smartid"`
	User    []byte `json:"user"`
	Payload []byte `json:"payload"`
	P       []byte `json:"publicKey"`
}

//NewTransaction func
func NewTransaction(data *TxData, from []byte) (*Transaction, error) {
	_size := unsafe.Sizeof(data)
	_byte := Bytes(data)
	_hash := sha256.Sum256(_byte)
	transaction := &Transaction{
		Size: int(_size),
		Data: *data,
		Hash: _hash,
		From: from,
	}

	return transaction, nil

}

func (t *Transaction) setHash() error {
	_byte, err := t.serialize()
	hash := sha256.Sum256(_byte)
	t.Hash = hash
	return err
}

//SerializeTransaction	serialize the transaction
func (t *Transaction) serialize() ([]byte, error) {
	return json.Marshal(t)
}

//DeserializeTransaction de-serialize the transaction
func SerializeTransaction(data []byte) (*Transaction, error) {
	t := new(Transaction)
	err := json.Unmarshal(data, t)
	return t, err
}

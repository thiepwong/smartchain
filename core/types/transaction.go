package types

import (
	"crypto/sha256"
	"unsafe"
)

//Transactions type
type Transactions []*Transaction

//Transaction  type
type Transaction struct {
	Hash Hash   `json:"hash" bson:"hash"`
	Size int    `json:"size"`
	From []byte `json:"from"`
	Data TxData `json:"data"`
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
		Hash: _hash[:],
		From: from,
	}

	return transaction, nil

}

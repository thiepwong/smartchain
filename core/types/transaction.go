package types

import (
	"unsafe"
)

//Transactions type
type Transactions []*Transaction

//Transaction  type
type Transaction struct {
	Hash []byte `json:"hash" bson:"hash"`
	Size int    `json:"size"`
	From []byte `json:"from"`
	Data TxData `json:"data"`
}

type TxData struct {
	//is register new or update data
	Mode    int    `json:"mode"`
	SmartID uint64 `json:"smartid"`
	User    []byte `json:"user"`
	Payload []byte `json:"payload"`
	P       []byte `json:"public"`
}

func NewTransaction(data *TxData) (*Transaction, error) {
	_size := unsafe.Sizeof(data)
	transaction := &Transaction{
		Size: int(_size),
		Data: *data,
	}

	return transaction, nil

}

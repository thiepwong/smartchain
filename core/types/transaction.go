package types

import (
	"unsafe"
)

//Transactions type
type Transactions []*Transaction

//Transaction  type
type Transaction struct {
	hash []byte
	size int
	from Address
	data TxData
}

type TxData struct {
	//is register new or update data
	Mode    int    `json: "mode"`
	SmartID uint64 `json: "smartid"`
	User    []byte `json: "user"`
	Payload []byte `json: "payload"`
	P       []byte `json: "public"`
}

func NewTransaction(data *TxData) (*Transaction, error) {
	_size := unsafe.Sizeof(data)
	transaction := &Transaction{
		size: int(_size),
		data: *data,
	}

	return transaction, nil

}

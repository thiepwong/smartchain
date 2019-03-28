package types

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

//TxSign func to sign the transaction
func TxSign(transaction *Transaction, prv *ecdsa.PrivateKey) (*Transaction, error) {

	tx := TxData{Mode: transaction.Data.Mode, SmartID: transaction.Data.SmartID, User: transaction.Data.User, Payload: transaction.Data.Payload, P: transaction.Data.P}
	_h := ToBytes(tx)
	hash := crypto.Keccak256Hash(_h)

	sig, err := crypto.Sign(hash[:], prv)
	if err != nil {
		return nil, err
	}
	fmt.Println("========== Da ky xong chu ky ==========")
	fmt.Printf("===== OK ====== \r\n %x", sig)

	return transaction, nil

}

package types

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

func Bytes(t interface{}) []byte {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, t)

	if err != nil {
		fmt.Println(err.Error())
	}

	return buf.Bytes()
}

//ToBytes function convert the interface to byte array
func ToBytes(inter interface{}) []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(inter)
	fmt.Println(reqBodyBytes.Bytes())         // this is the []byte
	fmt.Println(string(reqBodyBytes.Bytes())) // converted back to show it's your original object
	return reqBodyBytes.Bytes()
}

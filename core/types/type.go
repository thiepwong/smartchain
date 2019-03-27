package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Bytes(t interface{}) []byte {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, t)

	if err != nil {
		fmt.Println(err)
	}

	return buf.Bytes()
}

package types

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/thiepwong/smartchain/core/common"
)

func TestBlockEncoding(t *testing.T) {
	_dif := new(uint64)
	_heg := new(uint64)
	_time := new(uint64)
	*_time = uint64(time.Now().Unix())
	*_heg = 1025
	*_dif = 45655

	header := &Header{Version: 45, Difficulty: _dif, Height: _heg, Time: _time}
	tx, e := NewTransaction(&TxData{}, nil)
	if e != nil {
		fmt.Println(e.Error())
	}
	bl := NewBlock(header, []Transaction{*tx})
	s := bl.Header.Size()
	size := bl.GetSize()
	str := bl.String()
	fmt.Println("Kich thuoc header", size, s)
	fmt.Println(" block  ", str)

	lbs, e := bl.Header.serialize()
	check := func(f string, got, want interface{}) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s mismatch: got %v, want %v", f, got, want)
		}
	}

	//	serial, e := bl.serialize()

	check("Size", bl.Header.Size(), common.StorageSize(len(lbs)))

}

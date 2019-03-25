package types

//Transactions type
type Transactions []*Transaction

//Transaction  type
type Transaction struct {
	hash []byte
	size int
	from Address
	data txdata
}

type txdata struct {
	//is register new or update data
	Mode    int    `json: "mode"`
	SmartID uint64 `json: "smartid"`
	User    []byte `json: "user"`
	Payload []byte `json: "payload"`
	P       []byte `json: "public"`
}

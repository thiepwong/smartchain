package mongodb

import (
	"errors"

	"github.com/thiepwong/smartchain/core/types"
	"gopkg.in/mgo.v2"
)

//Database for mongodb
type Database struct {
	url string //dataname
	db  *mgo.Database
}

func getSession(url string) (*mgo.Session, error) {
	dialInfo, err := mgo.ParseURL(url)
	s, err := mgo.DialWithInfo(dialInfo)
	return s, err
}

//New func Create a new Database
func New(url string, dn string) (*Database, error) {
	if url == "" {
		return nil, errors.New("Url of server not found")
	}

	session, err := getSession(url)
	if err != nil {
		return nil, err
	}

	db := &Database{url: url, db: session.DB(dn)}
	return db, nil
}

//Add function to add new document to db
func (db *Database) Add(collection string, ojb interface{}) error {
	err := db.db.C(collection).Insert(ojb)
	if err != nil {
		return err
	}
	return nil
}

//Load func
func (db *Database) Load() *types.Block {
	bl := &types.Block{}
	db.db.C("mainchain").FindId(23258).One(bl)
	return bl
}

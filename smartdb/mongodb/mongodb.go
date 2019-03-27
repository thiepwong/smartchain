package mongodb

import (
	"errors"

	"gopkg.in/mgo.v2"
)

//Database for mongodb
type Database struct {
	url string //dataname
	db  *mgo.Database
}

//New func Create a new Database
func New(url string, dn string) (*Database, error) {
	if url == "" {
		return nil, errors.New("Url of server not found")
	}

	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	db := &Database{url: url, db: session.DB(dn)}
	return db, nil
}

func (db *Database) Insert(collection string, ojb interface{}) error {
	err := db.db.C(collection).Insert(ojb)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Load() interface{} {
	return db.db.C("smartchain").Find(nil)
}

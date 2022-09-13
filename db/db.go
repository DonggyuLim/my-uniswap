package db

import (
	"bytes"

	"log"

	"github.com/dgraph-io/badger/v3"
)

var db *badger.DB

func NewDB() {
	if db == nil {
		dbPointer, err := badger.Open(badger.DefaultOptions("./badger/contract"))
		if err != nil {
			log.Fatal(err)
		}
		db = dbPointer
	}
}

func Close() {
	db.Close()
}

func Add(key string, value []byte) {
	db.Update(func(txn *badger.Txn) error {
		txn.Set([]byte(key), []byte(value))
		return nil
	})
}

func Get(key string) ([]byte, bool) {
	var buf bytes.Buffer
	ok := false
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err == nil {
			item.Value(func(val []byte) error {
				buf.Write(val)
				ok = true
				return nil
			})
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes(), ok
}

func Remove(key string) {
	db.Update(func(txn *badger.Txn) error {
		txn.Delete([]byte(key))
		return nil
	})
}

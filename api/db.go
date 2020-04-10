package main

import (
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func init() {
	// Due to assignment issues:
	var err error
	db, err = bbolt.Open("wiking.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	go db.Update(func(tx *bbolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("Users"))
		_, err = tx.CreateBucketIfNotExists([]byte("Superusers"))
		return err
	})
}

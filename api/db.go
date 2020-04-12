package main

import (
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

// Registered users, user permissions, logged in users,
var buckets = []string{"Users", "Superusers", "Authorized", "Pages", "Content"}

func init() {
	// Due to assignment issues:
	var err error
	db, err = bbolt.Open("wiking.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	go db.Update(func(tx *bbolt.Tx) error {
		for _, v := range buckets {
			_, err = tx.CreateBucketIfNotExists([]byte(v))
		}
		return err
	})
}

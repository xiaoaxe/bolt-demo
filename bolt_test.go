package bolt_test

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/boltdb/bolt"
)

func init() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

/*
bucket、cursor、db、freelist、node、page、tx
*/
func TestBolt(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		db.Close()
	}()

	// write
	db.Update(func(tx *bolt.Tx) error {
		bkt, err := tx.CreateBucketIfNotExists(b("my.bkt"))
		if err != nil {
			log.Fatalf("create bucket err: %v", err)
			return err
		}

		// bkt.Put(b("answer"), b("42"))
		// bkt.Put(b("1+2"), b("3"))
		bkt.Put(b("current"), b(strconv.FormatInt(time.Now().Unix(), 10)))

		return nil
	})

	// read
	db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(b("my.bkt"))
		c := bkt.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("Got: %s => %s\n", k, v)
		}

		return nil
	})

}

func b(s string) []byte {
	return []byte(s)
}

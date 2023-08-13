package db

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

type LevelDbDriver struct {
	Path string
	db   *leveldb.DB
}

func (lDriver *LevelDbDriver) Init() {
	var err error
	lDriver.db, err = leveldb.OpenFile(lDriver.Path, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (lDriver *LevelDbDriver) Close() {
	if err := lDriver.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func (lDriver *LevelDbDriver) Get(key string) string {
	data, err := lDriver.db.Get([]byte(key), nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(data)
}

func (lDriver *LevelDbDriver) Set(key string, value string) error {
	return lDriver.db.Put([]byte(key), []byte(value), nil)
}

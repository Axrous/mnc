package app

import "github.com/sonyarouje/simdb"

func NewDB() *simdb.Driver {

	db, err := simdb.New("database")
	if err != nil {
		panic(err)
	}

	return db
}
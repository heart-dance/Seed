package db

type DB interface{}

type db struct{}

func NewDB() DB {
	return &db{}
}

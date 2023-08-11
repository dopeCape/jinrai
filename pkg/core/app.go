package core

type DB interface {
	write() error
}

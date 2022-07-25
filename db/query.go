package db

type QueryType int

const (
	UserExists = iota
)

type Database interface {
	QueryUserInformation(query QueryType, uuid string) (bool, error)
}

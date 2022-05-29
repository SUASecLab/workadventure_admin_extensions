package db

type QueryType int

const (
	UserExists = iota
	UserIsAdmin
)

type Database interface {
	QueryUserInformation(query QueryType, uuid string) (bool, error)
}

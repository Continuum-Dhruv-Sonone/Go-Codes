package db

type DbAccessor interface {
	Read() string
	Write(a int)
}

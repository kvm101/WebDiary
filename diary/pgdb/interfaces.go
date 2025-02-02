package pgdb

type Entry interface {
	Get() []byte
	Add() string
	Delete() string
	Update() string
}

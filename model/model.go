package model

const (
	female = iota
	male   = iota
	secret = iota
)

type User struct {
	AuthUser
	phone  string
	gender int8
}

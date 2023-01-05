package entity

type UserId string

type User struct {
	Id            UserId
	Age           int
	Name          string
	Authenticated bool
	Mail          string
	Icon          string
}

package entity

type UserId string

type User struct {
	Id            UserId
	Name          string
	Authenticated bool
	Mail          string
	Icon          string
}

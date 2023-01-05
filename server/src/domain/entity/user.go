package entity

type UserId string

type User struct {
	Id            UserId
	Name          string
	Authenticated bool
	Mail          string
	Icon          string
}

// TODO: ageってなんだw
type Participant struct {
	Id      UserId
	Name    string
	Icon    string
	Comment *EComment
}

package entity

type EventId string

type Event struct {
	Id           EventId
	Name         string
	Detail       string
	Location     string
	Admin        *User
	State        *EState
	Type         *EType
	Participants []*Participant
}

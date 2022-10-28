package model

type User struct {
	Id          int
	Name        string
	Age         int
	PhoneNumber string
}

func (User) TableName() string {
	return "user"
}

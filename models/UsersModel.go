package models

type User struct {
	ID        int
	Username  string
	Password  string
	FirstName string
	LastName  string
}

func (b *User) TableName() string {
	return "users"
}

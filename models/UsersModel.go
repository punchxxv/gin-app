package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func (b *User) TableName() string {
	return "users"
}

package model

type User struct {
	Id        int64
	Name      string
	Info      string
	Password  string
	Deleted   int8
	CreatedAt int64
	UpdatedAt int64
}

func (*User) TableName() string {
	return "user"
}

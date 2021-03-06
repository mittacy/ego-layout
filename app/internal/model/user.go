package model

type User struct {
	Id        int64
	Name      string
	Introduce string
	Password  string
	Deleted   int8
	CreatedAt int64
	UpdatedAt int64
}

func (*User) TableName() string {
	return "user"
}

const (
	// 状态
	UserStateDeleted = 100 // 删除
)

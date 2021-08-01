package userValidator

type CreateReq struct {
	Name     string `json:"name" binding:"required,min=1,max=20"`
	Info     string `json:"info" binding:"required,min=1,max=100"`
	Password string `json:"password" binding:"required,min=1,max=20"`
}
type CreateReply struct{}

type DeleteReq struct {
	Id int64 `json:"id" binding:"required,min=1"`
}
type DeleteReply struct{}

type UpdateReq struct {
	UpdateType int `json:"update_type" binding:"required,oneof=1 2"`
}

type UpdateInfoReq struct {
	Id   int64  `json:"id" binding:"required,min=1"`
	Name string `json:"name" binding:"required,min=1,max=20"`
	Info string `json:"info" binding:"required,min=1,max=100"`
}
type UpdateInfoReply struct{}

type UpdatePasswordReq struct {
	Id       int64  `json:"id" binding:"required,min=1"`
	Password string `json:"password" binding:"required,min=1,max=20"`
}
type UpdatePasswordReply struct{}

type GetReq struct {
	Id int64 `form:"id" binding:"required,min=1"`
}
type GetReply struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type ListReq struct {
	Page     int `json:"page" binding:"required,min=1"`
	PageSize int `json:"page_size" binding:"required,min=1,max=50"`
}
type ListReply struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

package userValidator

type GetReq struct {
	Id int64 `form:"id" json:"id" binding:"required,min=1"`
}
type GetReply struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

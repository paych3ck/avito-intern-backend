package models

type GetUserTendersRequest struct {
	Limit    int    `form:"limit" binding:"omitempty,min=0,max=50"`
	Offset   int    `form:"offset" binding:"omitempty,min=0"`
	Username string `form:"username" binding:"required"`
}

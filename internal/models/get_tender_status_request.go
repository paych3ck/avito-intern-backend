package models

type GetTenderStatusRequest struct {
	TenderID string `uri:"tenderId" binding:"required,uuid"`
	Username string `form:"username" binding:"required"`
}

package models

type UpdateTenderStatusRequest struct {
	TenderID string `uri:"tenderId" binding:"required,uuid"`
	Status   string `form:"status" binding:"required"`
	Username string `form:"username" binding:"required"`
}

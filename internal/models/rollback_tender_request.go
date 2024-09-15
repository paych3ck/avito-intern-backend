package models

type RollbackTenderRequest struct {
	TenderID string `uri:"tenderId" binding:"required,uuid"`
	Version  int    `uri:"version" binding:"required,min=1"`
	Username string `form:"username" binding:"required"`
}

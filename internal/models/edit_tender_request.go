package models

type EditTenderRequest struct {
	TenderID    string `uri:"tenderId" binding:"required,uuid"`
	Username    string `form:"username" binding:"required"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
}

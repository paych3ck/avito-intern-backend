package models

type GetTendersRequest struct {
	Limit       int      `form:"limit" binding:"omitempty,min=0,max=50"`
	Offset      int      `form:"offset" binding:"omitempty,min=0"`
	ServiceType []string `form:"service_type"`
}

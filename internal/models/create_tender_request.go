package models

type CreateTenderRequest struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	ServiceType     string `json:"serviceType" binding:"required"`
	Status          string `json:"status" binding:"required"`
	OrganizationID  string `json:"organizationId" binding:"required"`
	CreatorUsername string `json:"creatorUsername" binding:"required"`
}

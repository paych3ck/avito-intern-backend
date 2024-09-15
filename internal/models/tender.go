package models

type Tender struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ServiceType    string `json:"serviceType"`
	Status         string `json:"status"`
	OrganizationID string `json:"organizationId"`
	Version        int    `json:"version"`
	CreatedAt      string `json:"createdAt"`
}

package tools

import (
	"avito-intern-backend/internal/database"
	"avito-intern-backend/internal/models"
	"database/sql"
	"strings"
)

func IsValidUser(username string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := database.DB.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func IsValidOrganization(organizationID string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM organizations WHERE id = $1)`
	err := database.DB.QueryRow(query, organizationID).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func CreateTender(tender *models.CreateTenderRequest) (*models.Tender, error) {
	var newTender models.Tender
	query := `
		INSERT INTO tenders (name, description, service_type, status, organization_id, creator_username)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	err := database.DB.QueryRow(query, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationID, tender.CreatorUsername).
		Scan(&newTender.ID, &newTender.CreatedAt)
	if err != nil {
		return nil, err
	}

	newTender.Name = tender.Name
	newTender.Description = tender.Description
	newTender.ServiceType = tender.ServiceType
	newTender.Status = tender.Status
	newTender.OrganizationID = tender.OrganizationID
	newTender.Version = 1

	return &newTender, nil
}

func GetTenders(serviceTypes []string, limit, offset int) ([]models.Tender, error) {
	var tenders []models.Tender

	queryBuilder := strings.Builder{}
	queryBuilder.WriteString(`SELECT id, name, description, service_type, status, organization_id, version, created_at FROM tenders`)

	var queryParams []interface{}
	if len(serviceTypes) > 0 {
		queryBuilder.WriteString(` WHERE service_type = ANY($1)`)
		queryParams = append(queryParams, serviceTypes)
	}

	queryBuilder.WriteString(` ORDER BY name ASC LIMIT $2 OFFSET $3`)
	queryParams = append(queryParams, limit, offset)

	query := queryBuilder.String()
	rows, err := database.DB.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tender models.Tender
		if err := rows.Scan(&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationID, &tender.Version, &tender.CreatedAt); err != nil {
			return nil, err
		}
		tenders = append(tenders, tender)
	}

	return tenders, rows.Err()
}

func GetUserTenders(username string, limit, offset int) ([]models.Tender, error) {
	query := `
		SELECT id, name, description, service_type, status, organization_id, version, created_at
		FROM tenders
		WHERE creator_username = $1
		ORDER BY name ASC
		LIMIT $2 OFFSET $3
	`

	rows, err := database.DB.Query(query, username, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders []models.Tender
	for rows.Next() {
		var tender models.Tender
		if err := rows.Scan(&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationID, &tender.Version, &tender.CreatedAt); err != nil {
			return nil, err
		}
		tenders = append(tenders, tender)
	}

	return tenders, rows.Err()
}

func GetTenderStatus(tenderID string) (string, error) {
	var status string
	query := `SELECT status FROM tenders WHERE id = $1`
	err := database.DB.QueryRow(query, tenderID).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", sql.ErrNoRows
		}
		return "", err
	}
	return status, nil
}

func UpdateTenderStatus(tenderID string, status string) (*models.Tender, error) {
	var updatedTender models.Tender
	query := `
		UPDATE tenders
		SET status = $1
		WHERE id = $2
		RETURNING id, name, description, service_type, status, organization_id, version, created_at
	`
	err := database.DB.QueryRow(query, status, tenderID).Scan(
		&updatedTender.ID, &updatedTender.Name, &updatedTender.Description, &updatedTender.ServiceType,
		&updatedTender.Status, &updatedTender.OrganizationID, &updatedTender.Version, &updatedTender.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &updatedTender, nil
}

func RollbackTender(tenderID string, version int) (*models.Tender, error) {
	var tender models.Tender
	query := `SELECT id, name, description, service_type, status, organization_id, version, created_at FROM tender_versions WHERE id = $1 AND version = $2`
	err := database.DB.QueryRow(query, tenderID, version).Scan(
		&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType,
		&tender.Status, &tender.OrganizationID, &tender.Version, &tender.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	newVersion := tender.Version + 1
	query = `
		UPDATE tenders SET
		name = $1, description = $2, service_type = $3, version = $4
		WHERE id = $5
	`
	_, err = database.DB.Exec(query, tender.Name, tender.Description, tender.ServiceType, newVersion, tenderID)
	if err != nil {
		return nil, err
	}

	tender.Version = newVersion
	return &tender, nil
}

func GetTenderByID(tenderID string) (*models.Tender, error) {
	var tender models.Tender
	query := `
		SELECT id, name, description, service_type, status, organization_id, version, created_at 
		FROM tenders 
		WHERE id = $1
	`

	err := database.DB.QueryRow(query, tenderID).Scan(
		&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType,
		&tender.Status, &tender.OrganizationID, &tender.Version, &tender.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &tender, nil
}

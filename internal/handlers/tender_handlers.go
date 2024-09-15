package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"avito-intern-backend/internal/database"
	"avito-intern-backend/internal/models"
	"avito-intern-backend/internal/tools"

	"github.com/gin-gonic/gin"
)

func CreateTenderHandler(c *gin.Context) {
	var request models.CreateTenderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	if !tools.IsValidUser(request.CreatorUsername) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не существует или некорректен."})
		return
	}

	if !tools.IsValidOrganization(request.OrganizationID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Недостаточно прав для выполнения действия."})
		return
	}

	newTender, err := tools.CreateTender(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании тендера"})
		return
	}

	c.JSON(http.StatusOK, newTender)
}

func GetTendersHandler(c *gin.Context) {
	var request models.GetTendersRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса или параметры"})
		return
	}

	if request.Limit == 0 {
		request.Limit = 5
	}

	tenders, err := tools.GetTenders(request.ServiceType, request.Limit, request.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса к БД"})
		return
	}

	c.JSON(http.StatusOK, tenders)
}

func GetUserTendersHandler(c *gin.Context) {
	var request models.GetUserTendersRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса или параметры"})
		return
	}

	if !tools.IsValidUser(request.Username) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не существует или некорректен."})
		return
	}

	if request.Limit == 0 {
		request.Limit = 5
	}

	tenders, err := tools.GetUserTenders(request.Username, request.Limit, request.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса к БД"})
		return
	}

	c.JSON(http.StatusOK, tenders)
}

func GetTenderStatusHandler(c *gin.Context) {
	var request models.GetTenderStatusRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор тендера"})
		return
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса или параметры"})
		return
	}

	if !tools.IsValidUser(request.Username) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не существует или некорректен."})
		return
	}

	status, err := tools.GetTenderStatus(request.TenderID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Тендер с указанным идентификатором не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса к БД"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"tenderStatus": status})
}

func UpdateTenderStatusHandler(c *gin.Context) {
	var request models.UpdateTenderStatusRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор тендера"})
		return
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса или параметры"})
		return
	}

	if !tools.IsValidUser(request.Username) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не существует или некорректен."})
		return
	}

	updatedTender, err := tools.UpdateTenderStatus(request.TenderID, request.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Тендер не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса тендера"})
		}
		return
	}
	c.JSON(http.StatusOK, updatedTender)
}

func EditTenderHandler(c *gin.Context) {
	var request models.EditTenderRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор тендера"})
		return
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса или параметры"})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if !tools.IsValidUser(request.Username) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не существует или некорректен."})
		return
	}

	var queryBuilder strings.Builder
	var queryParams []interface{}
	queryBuilder.WriteString("UPDATE tenders SET")

	paramCounter := 1
	fieldsToUpdate := 0

	if request.Name != "" {
		if fieldsToUpdate > 0 {
			queryBuilder.WriteString(", ")
		}
		queryBuilder.WriteString(" name = $" + fmt.Sprint(paramCounter))
		queryParams = append(queryParams, request.Name)
		paramCounter++
		fieldsToUpdate++
	}

	if request.Description != "" {
		if fieldsToUpdate > 0 {
			queryBuilder.WriteString(", ")
		}
		queryBuilder.WriteString(" description = $" + fmt.Sprint(paramCounter))
		queryParams = append(queryParams, request.Description)
		paramCounter++
		fieldsToUpdate++
	}

	if request.ServiceType != "" {
		if fieldsToUpdate > 0 {
			queryBuilder.WriteString(", ")
		}
		queryBuilder.WriteString(" service_type = $" + fmt.Sprint(paramCounter))
		queryParams = append(queryParams, request.ServiceType)
		paramCounter++
		fieldsToUpdate++
	}

	if fieldsToUpdate == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не переданы параметры для изменения"})
		return
	}

	queryBuilder.WriteString(" WHERE id = $" + fmt.Sprint(paramCounter))
	queryParams = append(queryParams, request.TenderID)

	query := queryBuilder.String()
	_, err := database.DB.Exec(query, queryParams...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении тендера"})
		return
	}

	updatedTender, err := tools.GetTenderByID(request.TenderID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Тендер не найден"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении обновленной информации о тендере"})
		return
	}

	c.JSON(http.StatusOK, updatedTender)
}

func RollbackTenderHandler(c *gin.Context) {
	var request models.RollbackTenderRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат идентификатора тендера или версии"})
		return
	}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные параметры запроса"})
		return
	}

	if !tools.IsValidUser(request.Username) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не существует или некорректен"})
		return
	}

	tender, err := tools.RollbackTender(request.TenderID, request.Version)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Тендер или версия не найдены"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при откате версии тендера"})
		}
		return
	}

	c.JSON(http.StatusOK, tender)
}

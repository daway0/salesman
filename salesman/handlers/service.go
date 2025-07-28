package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceHandler struct {
	DB *sql.DB
}

func (h *ServiceHandler) CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.ID = uuid.New()
	service.CreatedAt = time.Now()

	query := `INSERT INTO Services (id, company_id, title, description, price, image_url, type, created_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := h.DB.QueryRow(query, service.ID, service.CompanyID, service.Title,
		service.Description, service.Price, service.ImageURL, service.Type, service.CreatedAt).Scan(&service.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, service)
}

func (h *ServiceHandler) GetService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	query := `SELECT s.id, s.company_id, s.title, s.description, s.price, s.image_url, s.type, s.created_at, s.updated_at, s.deleted_at, c.title
              FROM Services s INNER JOIN Companies c ON s.company_id = c.id WHERE s.id = $1 AND s.deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&service.ID, &service.CompanyID, &service.Title,
		&service.Description, &service.Price, &service.ImageURL, &service.Type, &service.CreatedAt, &service.UpdatedAt, &service.DeletedAt, &service.CompanyTitle)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}

func (h *ServiceHandler) GetServices(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT s.id, s.company_id, s.title, s.description, s.price, s.image_url, s.type, s.created_at, s.updated_at, s.deleted_at, c.title
                             FROM Services s INNER JOIN Companies c ON s.company_id = c.id WHERE s.deleted_at IS NULL`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(&service.ID, &service.CompanyID, &service.Title,
			&service.Description, &service.Price, &service.ImageURL, &service.Type, &service.CreatedAt, &service.UpdatedAt, &service.DeletedAt, &service.CompanyTitle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		services = append(services, service)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

func (h *ServiceHandler) UpdateService(c *gin.Context) {
	idStr := c.Param("id")

	serviceID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID format"})
		return
	}

	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.ID = serviceID
	now := time.Now()
	service.UpdatedAt = &now

	query := `UPDATE Services SET company_id = $1, title = $2, description = $3, price = $4, image_url = $5, type = $6, updated_at = $7
              WHERE id = $8 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, service.CompanyID, service.Title,
		service.Description, service.Price, service.ImageURL, service.Type, &now, service.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, service)
}

func (h *ServiceHandler) DeleteService(c *gin.Context) {
	id := c.Param("id")
	query := `UPDATE Services SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, time.Now(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})
}

package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CompanyHandler struct {
	DB *sql.DB
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company.ID = uuid.New()
	company.CreatedAt = time.Now()

	query := `INSERT INTO Companies (id, title, brand_name, cid, description, image_url, created_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := h.DB.QueryRow(query, company.ID, company.Title, company.BrandName, company.CID,
		company.Description, company.ImageURL, company.CreatedAt).Scan(&company.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, company)
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	query := `SELECT id, title, brand_name, cid, description, image_url, created_at, updated_at, deleted_at
              FROM Companies WHERE id = $1 AND deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&company.ID, &company.Title, &company.BrandName, &company.CID,
		&company.Description, &company.ImageURL, &company.CreatedAt, &company.UpdatedAt, &company.DeletedAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) GetCompanies(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, title, brand_name, cid, description, image_url, created_at, updated_at, deleted_at 
                             FROM Companies WHERE deleted_at IS NULL`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var companies []models.Company
	for rows.Next() {
		var company models.Company
		err := rows.Scan(&company.ID, &company.Title, &company.BrandName, &company.CID,
			&company.Description, &company.ImageURL, &company.CreatedAt, &company.UpdatedAt, &company.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		companies = append(companies, company)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	idStr := c.Param("id")

	companyID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID format"})
		return
	}

	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company.ID = companyID
	company.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	query := `UPDATE Companies SET title = $1, brand_name = $2, cid = $3, description = $4, image_url = $5, updated_at = $6
              WHERE id = $7 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, company.Title, company.BrandName, company.CID,
		company.Description, company.ImageURL, company.UpdatedAt, company.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	deletedAt := time.Now()
	query := `UPDATE Companies SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, deletedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}

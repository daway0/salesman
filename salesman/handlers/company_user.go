package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CompanyUserHandler struct {
	DB *sql.DB
}

func (h *CompanyUserHandler) CreateCompanyUser(c *gin.Context) {
	var companyUser models.CompanyUser
	if err := c.ShouldBindJSON(&companyUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that company exists
	var count int
	err := h.DB.QueryRow("SELECT COUNT(1) FROM Companies WHERE id = $1 AND deleted_at IS NULL", companyUser.CompanyID).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate company: " + err.Error()})
		return
	}
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company not found"})
		return
	}

	// Validate that user exists
	err = h.DB.QueryRow("SELECT COUNT(1) FROM Users WHERE id = $1 AND deleted_at IS NULL", companyUser.UserID).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate user: " + err.Error()})
		return
	}
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Check if company-user relationship already exists
	err = h.DB.QueryRow("SELECT COUNT(1) FROM CompanyUsers WHERE company_id = $1 AND user_id = $2 AND deleted_at IS NULL",
		companyUser.CompanyID, companyUser.UserID).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing relationship: " + err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is already associated with this company"})
		return
	}

	companyUser.ID = uuid.New()
	companyUser.CreatedAt = time.Now()

	query := `INSERT INTO CompanyUsers (id, company_id, user_id, subscription_no, created_at)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = h.DB.QueryRow(query, companyUser.ID, companyUser.CompanyID, companyUser.UserID,
		companyUser.SubscriptionNo, companyUser.CreatedAt).Scan(&companyUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, companyUser)
}

func (h *CompanyUserHandler) GetCompanyUser(c *gin.Context) {
	id := c.Param("id")
	var companyUser models.CompanyUser
	query := `SELECT id, company_id, user_id, subscription_no, created_at, updated_at, deleted_at
              FROM CompanyUsers WHERE id = $1 AND deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&companyUser.ID, &companyUser.CompanyID, &companyUser.UserID,
		&companyUser.SubscriptionNo, &companyUser.CreatedAt, &companyUser.UpdatedAt, &companyUser.DeletedAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companyUser)
}

func (h *CompanyUserHandler) GetCompanyUsers(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, company_id, user_id, subscription_no, created_at, updated_at, deleted_at 
                             FROM CompanyUsers WHERE deleted_at IS NULL`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var companyUsers []models.CompanyUser
	for rows.Next() {
		var companyUser models.CompanyUser
		err := rows.Scan(&companyUser.ID, &companyUser.CompanyID, &companyUser.UserID,
			&companyUser.SubscriptionNo, &companyUser.CreatedAt, &companyUser.UpdatedAt, &companyUser.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		companyUsers = append(companyUsers, companyUser)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companyUsers)
}

func (h *CompanyUserHandler) GetCompanyUsersByUser(c *gin.Context) {
	userID := c.Param("user_id")
	rows, err := h.DB.Query(`SELECT id, company_id, user_id, subscription_no, created_at, updated_at, deleted_at 
                             FROM CompanyUsers WHERE user_id = $1 AND deleted_at IS NULL`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var companyUsers []models.CompanyUser
	for rows.Next() {
		var companyUser models.CompanyUser
		err := rows.Scan(&companyUser.ID, &companyUser.CompanyID, &companyUser.UserID,
			&companyUser.SubscriptionNo, &companyUser.CreatedAt, &companyUser.UpdatedAt, &companyUser.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		companyUsers = append(companyUsers, companyUser)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companyUsers)
}

func (h *CompanyUserHandler) UpdateCompanyUser(c *gin.Context) {
	idStr := c.Param("id")

	var companyUser models.CompanyUser
	if err := c.ShouldBindJSON(&companyUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyUser.ID = uuid.MustParse(idStr)
	now := time.Now()
	companyUser.UpdatedAt = &now

	query := `UPDATE CompanyUsers SET subscription_no = $1, updated_at = $2
              WHERE id = $3 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, companyUser.SubscriptionNo, companyUser.UpdatedAt, companyUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company user not found"})
		return
	}

	c.JSON(http.StatusOK, companyUser)
}

func (h *CompanyUserHandler) DeleteCompanyUser(c *gin.Context) {
	idStr := c.Param("id")

	now := time.Now()
	query := `UPDATE CompanyUsers SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, now, idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company user deleted successfully"})
}

package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"salesman/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserRoleHandler struct {
	DB *sql.DB
}

func (h *UserRoleHandler) CreateUserRole(c *gin.Context) {
	var userRole models.UserRole
	
	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}

	log.Println("hello there")

	userRole.ID = uuid.New()

	query := `INSERT INTO UserRoles (id, user_id, role_id)
              VALUES ($1, $2, $3) RETURNING id`
	err := h.DB.QueryRow(query, userRole.ID, userRole.UserID, userRole.RoleID).Scan(&userRole.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userRole)
}

func (h *UserRoleHandler) GetUserRole(c *gin.Context) {
	id := c.Param("id")
	var userRole models.UserRole
	query := `SELECT id, user_id, role_id FROM UserRoles WHERE id = $1`
	err := h.DB.QueryRow(query, id).Scan(&userRole.ID, &userRole.UserID, &userRole.RoleID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User role assignment not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userRole)
}

func (h *UserRoleHandler) GetUserRoles(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, user_id, role_id FROM UserRoles`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var userRoles []models.UserRole
	for rows.Next() {
		var userRole models.UserRole
		err := rows.Scan(&userRole.ID, &userRole.UserID, &userRole.RoleID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userRoles = append(userRoles, userRole)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userRoles)
}

func (h *UserRoleHandler) GetUserRolesByUserID(c *gin.Context) {
	userID := c.Param("id")
	rows, err := h.DB.Query(`SELECT id, user_id, role_id FROM UserRoles WHERE user_id = $1`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var userRoles []models.UserRole
	for rows.Next() {
		var userRole models.UserRole
		err := rows.Scan(&userRole.ID, &userRole.UserID, &userRole.RoleID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userRoles = append(userRoles, userRole)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userRoles)
}



func (h *UserRoleHandler) DeleteUserRole(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM UserRoles WHERE id = $1`
	result, err := h.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User role assignment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User role assignment deleted"})
}


package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleHandler struct {
	DB *sql.DB
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.ID = uuid.New()
	role.CreatedAt = time.Now()

	query := `INSERT INTO Roles (id, title, description, created_at)
              VALUES ($1, $2, $3, $4) RETURNING id`
	err := h.DB.QueryRow(query, role.ID, role.Title, role.Description, role.CreatedAt).Scan(&role.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) GetRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	query := `SELECT id, title, description, created_at, updated_at, deleted_at
              FROM Roles WHERE id = $1 AND deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&role.ID, &role.Title, &role.Description,
		&role.CreatedAt, &role.UpdatedAt, &role.DeletedAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) GetRoles(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, title, description, created_at, updated_at, deleted_at 
                             FROM Roles WHERE deleted_at IS NULL`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Title, &role.Description,
			&role.CreatedAt, &role.UpdatedAt, &role.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		roles = append(roles, role)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")

	roleID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID format"})
		return
	}

	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.ID = roleID

	query := `UPDATE Roles SET title = $1, description = $2
              WHERE id = $3`
	result, err := h.DB.Exec(query, role.Title, role.Description, role.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM Roles WHERE id = $1`
	result, err := h.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}

func (h *RoleHandler) GetUserRoles(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	rows, err := h.DB.Query(`SELECT id, title, description FROM Roles WHERE deleted_at IS NULL`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch roles: " + err.Error()})
		return
	}
	defer rows.Close()

	var roles []models.UserRoleInfo
	for rows.Next() {
		var roleID uuid.UUID
		var title string
		var description *string

		if err := rows.Scan(&roleID, &title, &description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		
		var count int
		err = h.DB.QueryRow(`SELECT COUNT(1) FROM UserRoles WHERE user_id = $1 AND role_id = $2 AND deleted_at IS NULL`, userID, roleID).Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		roles = append(roles, models.UserRoleInfo{
			RoleID:      roleID,
			Title:       title,
			Description: description,
			UserHasRole: count > 0,
		})
	}

	c.JSON(http.StatusOK, roles)
}
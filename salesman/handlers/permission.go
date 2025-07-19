package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionHandler struct {
	DB *sql.DB
}

func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission.ID = uuid.New()
	permission.CreatedAt = time.Now()

	query := `INSERT INTO Permissions (id, content_type, action_type, created_at)
              VALUES ($1, $2, $3, $4) RETURNING id`
	err := h.DB.QueryRow(query, permission.ID, permission.ContentType, permission.ActionType, permission.CreatedAt).Scan(&permission.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, permission)
}

func (h *PermissionHandler) GetPermission(c *gin.Context) {
	id := c.Param("id")
	var permission models.Permission
	query := `SELECT id, content_type, action_type, created_at, updated_at, deleted_at
              FROM Permissions WHERE id = $1 AND deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&permission.ID, &permission.ContentType, &permission.ActionType,
		&permission.CreatedAt, &permission.UpdatedAt, &permission.DeletedAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, permission)
}

func (h *PermissionHandler) GetPermissions(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, content_type, action_type, created_at, updated_at, deleted_at 
                             FROM Permissions WHERE deleted_at IS NULL`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var permissions []models.Permission
	for rows.Next() {
		var permission models.Permission
		err := rows.Scan(&permission.ID, &permission.ContentType, &permission.ActionType,
			&permission.CreatedAt, &permission.UpdatedAt, &permission.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		permissions = append(permissions, permission)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, permissions)
}

func (h *PermissionHandler) UpdatePermission(c *gin.Context) {
	idStr := c.Param("id")

	permissionID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID format"})
		return
	}

	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission.ID = permissionID
	permission.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	query := `UPDATE Permissions SET content_type = $1, action_type = $2, updated_at = $3
              WHERE id = $4 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, permission.ContentType, permission.ActionType, permission.UpdatedAt, permission.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, permission)
}

func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	deletedAt := time.Now()
	query := `UPDATE Permissions SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, deletedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Permission deleted"})
}

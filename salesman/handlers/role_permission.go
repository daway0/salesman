package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RolePermissionHandler struct {
	DB *sql.DB
}

func (h *RolePermissionHandler) CreateRolePermission(c *gin.Context) {
	var rolePermission models.RolePermission
	if err := c.ShouldBindJSON(&rolePermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rolePermission.ID = uuid.New()

	query := `INSERT INTO RolePermissions (id, role_id, permission_id)
              VALUES ($1, $2, $3) RETURNING id`
	err := h.DB.QueryRow(query, rolePermission.ID, rolePermission.RoleID, rolePermission.PermissionID).Scan(&rolePermission.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, rolePermission)
}

func (h *RolePermissionHandler) GetRolePermission(c *gin.Context) {
	id := c.Param("id")
	var rolePermission models.RolePermission
	query := `SELECT id, role_id, permission_id FROM RolePermissions WHERE id = $1`
	err := h.DB.QueryRow(query, id).Scan(&rolePermission.ID, &rolePermission.RoleID, &rolePermission.PermissionID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role permission assignment not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rolePermission)
}

func (h *RolePermissionHandler) GetRolePermissions(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, role_id, permission_id FROM RolePermissions`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var rolePermissions []models.RolePermission
	for rows.Next() {
		var rolePermission models.RolePermission
		err := rows.Scan(&rolePermission.ID, &rolePermission.RoleID, &rolePermission.PermissionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rolePermissions)
}

func (h *RolePermissionHandler) GetRolePermissionsByRoleID(c *gin.Context) {
	roleID := c.Param("role_id")
	rows, err := h.DB.Query(`SELECT id, role_id, permission_id FROM RolePermissions WHERE role_id = $1`, roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var rolePermissions []models.RolePermission
	for rows.Next() {
		var rolePermission models.RolePermission
		err := rows.Scan(&rolePermission.ID, &rolePermission.RoleID, &rolePermission.PermissionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rolePermissions)
}

func (h *RolePermissionHandler) UpdateRolePermission(c *gin.Context) {
	idStr := c.Param("id")

	rolePermissionID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role permission ID format"})
		return
	}

	var rolePermission models.RolePermission
	if err := c.ShouldBindJSON(&rolePermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rolePermission.ID = rolePermissionID

	query := `UPDATE RolePermissions SET role_id = $1, permission_id = $2 WHERE id = $3`
	result, err := h.DB.Exec(query, rolePermission.RoleID, rolePermission.PermissionID, rolePermission.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role permission assignment not found"})
		return
	}

	c.JSON(http.StatusOK, rolePermission)
}

func (h *RolePermissionHandler) DeleteRolePermission(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM RolePermissions WHERE id = $1`
	result, err := h.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role permission assignment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role permission assignment deleted"})
}

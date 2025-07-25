package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	DB *sql.DB
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int
	err := h.DB.QueryRow("SELECT COUNT(1) FROM Users WHERE nsid = $1 ", user.NSID).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check NSID duplication: " + err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NSID already exists"})
		return
	}

	err = h.DB.QueryRow("SELECT COUNT(1) FROM Users WHERE email = $1", user.Email).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email duplication: " + err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	query := `INSERT INTO Users (id, first_name, last_name, NSID, birthdate, email, password, created_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err = h.DB.QueryRow(query, user.ID, user.FirstName, user.LastName, user.NSID, user.Birthdate, user.Email, user.Password, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	query := `SELECT id, first_name, last_name, NSID, birthdate, email, password, created_at, updated_at, deleted_at, last_login_at
              FROM Users WHERE id = $1 AND deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.NSID, &user.Birthdate, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.LastLoginAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, first_name, last_name, NSID, birthdate, email, password, created_at, updated_at, deleted_at, last_login_at 
                             FROM Users WHERE deleted_at IS NULL`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.NSID, &user.Birthdate, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.LastLoginAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")

	userID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = userID
	now := time.Now()
	user.UpdatedAt = &now

	query := `UPDATE Users SET first_name = $1, last_name = $2, NSID = $3, birthdate = $4, email = $5, password = $6, updated_at = $7
              WHERE id = $8 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, user.FirstName, user.LastName, user.NSID, user.Birthdate, user.Email, user.Password, user.UpdatedAt, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	now := time.Now()
	query := `UPDATE Users SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, now, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (h *UserHandler) AddUserRole(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	roleID, err := uuid.Parse(c.Param("role_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID format"})
		return
	}
	query := `INSERT INTO UserRoles (id, user_id, role_id, created_at) VALUES ($1, $2, $3, $4)`
	_, err = h.DB.Exec(query, uuid.New(), userID, roleID, time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User role added"})
}

func (h *UserHandler) RemoveUserRole(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	roleID, err := uuid.Parse(c.Param("role_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID format"})
		return
	}
	query := `UPDATE UserRoles SET deleted_at = $1 WHERE user_id = $2 AND role_id = $3 AND deleted_at IS NULL`
	_, err = h.DB.Exec(query, time.Now(), userID, roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User role removed"})
}

func (h *UserHandler) GetUserPermissions(c *gin.Context) {
	id := c.Param("id")
	query := `SELECT DISTINCT p.action FROM Users u
		INNER JOIN UserRoles ur ON ur.user_id = u.id AND ur.deleted_at IS NULL
		INNER JOIN RolePermissions rp ON rp.role_id = ur.role_id AND rp.deleted_at IS NULL
		INNER JOIN Permissions p ON rp.permission_id = p.id
		WHERE u.id = $1`

	rows, err := h.DB.Query(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var action string
		if err := rows.Scan(&action); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		permissions = append(permissions, action)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"permissions": permissions})
}

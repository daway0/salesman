package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	DB *sql.DB
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uuid.New()
	product.CreatedAt = time.Now()

	query := `INSERT INTO Products (id, company_id, title, has_commission, created_at)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := h.DB.QueryRow(query, product.ID, product.CompanyID, product.Title, product.HasCommission, product.CreatedAt).Scan(&product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.ProductWithCompany
	query := `SELECT p.id, p.company_id, c.title as company_title, p.title, p.has_commission, p.created_at, p.updated_at, p.deleted_at
              FROM Products p INNER JOIN Companies c ON p.company_id = c.id WHERE p.id = $1 AND p.deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&product.ID, &product.CompanyID, &product.CompanyTitle, &product.Title, &product.HasCommission, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT p.id, p.company_id, c.title as company_title, p.title, p.has_commission, p.created_at, p.updated_at, p.deleted_at
                             FROM Products p INNER JOIN Companies c ON p.company_id = c.id WHERE p.deleted_at IS NULL`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []models.ProductWithCompany
	for rows.Next() {
		var product models.ProductWithCompany
		err := rows.Scan(&product.ID, &product.CompanyID, &product.CompanyTitle, &product.Title, &product.HasCommission, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductsByCompany(c *gin.Context) {
	companyID := c.Param("company_id")
	rows, err := h.DB.Query(`SELECT id, company_id, title, has_commission, created_at, updated_at, deleted_at 
                             FROM Products WHERE company_id = $1 AND deleted_at IS NULL`, companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.CompanyID, &product.Title, &product.HasCommission, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")

	productID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = productID
	now := time.Now()
	product.UpdatedAt = &now

	query := `UPDATE Products SET company_id = $1, title = $2, has_commission = $3, updated_at = $4
              WHERE id = $5 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, product.CompanyID, product.Title, product.HasCommission, &now, product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	deletedAt := time.Now()
	query := `UPDATE Products SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, deletedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}


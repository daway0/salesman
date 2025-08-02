package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"
	"salesman/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerHandler struct {
	DB *sql.DB
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer models.CustomerCreate
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerID := uuid.New()
	createdAt := time.Now()
	birthdate := time.Now()
	email := customer.NSID + "@nit.ir"
	password, _ := utils.HashPasswordForDjango(customer.NSID)
	query := `INSERT INTO Users (id, first_name, last_name, NSID, birthdate, email, password, created_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := h.DB.QueryRow(query, customerID, customer.FirstName, customer.LastName, customer.NSID, birthdate, email, password, createdAt).Scan(&customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": customerID})
}

func (h *CustomerHandler) CreateCompanyCustomerProduct(c *gin.Context) {
	var customer models.CompanyCustomerProductCreate
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := h.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction: " + err.Error()})
		return
	}
	defer tx.Rollback()

	// First, get the company ID from the product ID
	var companyID uuid.UUID
	err = tx.QueryRow("SELECT company_id FROM Products WHERE id = $1 AND deleted_at IS NULL", customer.ProductID).Scan(&companyID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find product: " + err.Error()})
		}
		return
	}

	var companyUserID uuid.UUID
	var count int
	err = tx.QueryRow("SELECT COUNT(1) FROM CompanyUsers WHERE company_id = $1 AND user_id = $2 AND deleted_at IS NULL",
		companyID, customer.UserID).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing company user: " + err.Error()})
		return
	}

	if count == 0 {
		companyUserID = uuid.New()
		createdAt := time.Now()

		query := `INSERT INTO CompanyUsers (id, company_id, user_id, subscription_no, created_at)
                  VALUES ($1, $2, $3, $4, $5) RETURNING id`
		err = tx.QueryRow(query, companyUserID, companyID, customer.UserID, customer.SubscriptionNo, createdAt).Scan(&companyUserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company user: " + err.Error()})
			return
		}
	} else {
		err = tx.QueryRow("SELECT id FROM CompanyUsers WHERE company_id = $1 AND user_id = $2 AND deleted_at IS NULL",
			companyID, customer.UserID).Scan(&companyUserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get existing company user: " + err.Error()})
			return
		}
	}	

	companyUserProductID := uuid.New()
	createdAt := time.Now()

	query := `INSERT INTO CompanyUserProducts (id, company_user_id, product_id, is_active, number, created_at)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = tx.QueryRow(query, companyUserProductID, companyUserID, customer.ProductID, true, customer.Number, createdAt).Scan(&companyUserProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company user product: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": companyUserProductID})
}

func (h *CustomerHandler) GetCompanyUserProducts(c *gin.Context) {
	rows, err := h.DB.Query(`select CompanyUserProducts.id, Users.id as user_id, Users.first_name, Users.last_name, Users.NSID, Concat(Companies.title, ' - ', Products.title) as product_company_title, CompanyUsers.subscription_no, CompanyUserProducts.number, CompanyUserProducts.is_active, CompanyUserProducts.created_at
from CompanyUserProducts
inner join CompanyUsers on CompanyUserProducts.company_user_id = CompanyUsers.id
inner join Products on CompanyUserProducts.product_id = Products.id
inner join Companies on Products.company_id = Companies.id
inner join Users on CompanyUsers.user_id = Users.id
where CompanyUserProducts.deleted_at is null
and CompanyUsers.deleted_at is null
and Products.deleted_at is null
and Companies.deleted_at is null
and Users.deleted_at is null
and CompanyUserProducts.is_active = true`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var companyUserProducts []models.CompanyUserProductList
	for rows.Next() {
		var companyUserProduct models.CompanyUserProductList
		err := rows.Scan(&companyUserProduct.ID, &companyUserProduct.UserID, &companyUserProduct.FirstName, &companyUserProduct.LastName, &companyUserProduct.NSID, &companyUserProduct.ProductCompanyTitle, &companyUserProduct.SubscriptionNo, &companyUserProduct.Number, &companyUserProduct.IsActive, &companyUserProduct.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		companyUserProducts = append(companyUserProducts, companyUserProduct)
	}
	c.JSON(http.StatusOK, companyUserProducts)
}

func (h *CustomerHandler) GetCompanyUserProductsByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	rows, err := h.DB.Query(`select CompanyUserProducts.id, Users.id as user_id, Users.first_name, Users.last_name, Users.NSID, Concat(Companies.title, ' - ', Products.title) as product_company_title, CompanyUsers.subscription_no, CompanyUserProducts.number, CompanyUserProducts.is_active, CompanyUserProducts.created_at
from CompanyUserProducts
inner join CompanyUsers on CompanyUserProducts.company_user_id = CompanyUsers.id
inner join Products on CompanyUserProducts.product_id = Products.id
inner join Companies on Products.company_id = Companies.id
inner join Users on CompanyUsers.user_id = Users.id
where CompanyUserProducts.deleted_at is null
and CompanyUsers.deleted_at is null
and Products.deleted_at is null
and Companies.deleted_at is null
and Users.deleted_at is null
and CompanyUserProducts.is_active = true
and Users.id = $1`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var companyUserProducts []models.CompanyUserProductList
	for rows.Next() {
		var companyUserProduct models.CompanyUserProductList
		err := rows.Scan(&companyUserProduct.ID, &companyUserProduct.UserID, &companyUserProduct.FirstName, &companyUserProduct.LastName, &companyUserProduct.NSID, &companyUserProduct.ProductCompanyTitle, &companyUserProduct.SubscriptionNo, &companyUserProduct.Number, &companyUserProduct.IsActive, &companyUserProduct.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		companyUserProducts = append(companyUserProducts, companyUserProduct)
	}
	c.JSON(http.StatusOK, companyUserProducts)
}
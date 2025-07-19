package handlers

import (
	"database/sql"
	"net/http"
	"salesman/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SalesLedgerHandler struct {
	DB *sql.DB
}

func (h *SalesLedgerHandler) CreateSalesLedger(c *gin.Context) {
	var salesLedger models.SalesLedger
	if err := c.ShouldBindJSON(&salesLedger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesLedger.ID = uuid.New()
	salesLedger.CreatedAt = time.Now()

	approvedBy := models.NullUUID{UUID: salesLedger.ApprovedBy.UUID, Valid: salesLedger.ApprovedBy.Valid}
	approvedAt := sql.NullTime{Time: salesLedger.ApprovedAt.Time, Valid: salesLedger.ApprovedAt.Valid}

	query := `INSERT INTO SalesLedger (id, approved_by, service_id, price, sales_price, sales_discount, trn, approved_at, created_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	err := h.DB.QueryRow(query, salesLedger.ID, approvedBy, salesLedger.ServiceID,
		salesLedger.Price, salesLedger.SalesPrice, salesLedger.SalesDiscount, salesLedger.TRN,
		approvedAt, salesLedger.CreatedAt).Scan(&salesLedger.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, salesLedger)
}

func (h *SalesLedgerHandler) GetSalesLedger(c *gin.Context) {
	id := c.Param("id")
	var salesLedger models.SalesLedger
	query := `SELECT id, approved_by, service_id, price, sales_price, sales_discount, trn, approved_at, created_at, updated_at, deleted_at
              FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&salesLedger.ID, &salesLedger.ApprovedBy, &salesLedger.ServiceID,
		&salesLedger.Price, &salesLedger.SalesPrice, &salesLedger.SalesDiscount, &salesLedger.TRN,
		&salesLedger.ApprovedAt, &salesLedger.CreatedAt, &salesLedger.UpdatedAt, &salesLedger.DeletedAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, salesLedger)
}

func (h *SalesLedgerHandler) GetSalesLedgers(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, approved_by, service_id, price, sales_price, sales_discount, trn, approved_at, created_at, updated_at, deleted_at 
                             FROM SalesLedger WHERE deleted_at IS NULL`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var salesLedgers []models.SalesLedger
	for rows.Next() {
		var salesLedger models.SalesLedger
		err := rows.Scan(&salesLedger.ID, &salesLedger.ApprovedBy, &salesLedger.ServiceID,
			&salesLedger.Price, &salesLedger.SalesPrice, &salesLedger.SalesDiscount, &salesLedger.TRN,
			&salesLedger.ApprovedAt, &salesLedger.CreatedAt, &salesLedger.UpdatedAt, &salesLedger.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		salesLedgers = append(salesLedgers, salesLedger)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, salesLedgers)
}

func (h *SalesLedgerHandler) UpdateSalesLedger(c *gin.Context) {
	idStr := c.Param("id")

	salesLedgerID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sales ledger ID format"})
		return
	}

	var salesLedger models.SalesLedger
	if err := c.ShouldBindJSON(&salesLedger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesLedger.ID = salesLedgerID
	salesLedger.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	query := `UPDATE SalesLedger SET approved_by = $1, service_id = $2, price = $3, sales_price = $4, sales_discount = $5, trn = $6, approved_at = $7, updated_at = $8
              WHERE id = $9 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, salesLedger.ApprovedBy, salesLedger.ServiceID,
		salesLedger.Price, salesLedger.SalesPrice, salesLedger.SalesDiscount, salesLedger.TRN,
		salesLedger.ApprovedAt, salesLedger.UpdatedAt, salesLedger.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, salesLedger)
}

func (h *SalesLedgerHandler) DeleteSalesLedger(c *gin.Context) {
	id := c.Param("id")
	deletedAt := time.Now()
	query := `UPDATE SalesLedger SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, deletedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sales ledger entry deleted"})
}

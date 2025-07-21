package handlers

import (
	"database/sql"
	"encoding/json"
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

	initialWorkflow := models.Workflow{
		ActionAt: salesLedger.CreatedAt,
		Comment:  "",
		ActorId:  salesLedger.CreatedBy,
		Step:     "ایجاد فاکتور و ارسال به مالی",					
	}
	workflows := []models.Workflow{initialWorkflow}
	if data, err := json.Marshal(workflows); err == nil {
		msg := json.RawMessage(data)
		salesLedger.WorkflowHistory = &msg
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal workflow history: " + err.Error()})
		return
	}
	salesLedger.Status = "PENDING"

	query := `INSERT INTO SalesLedger (id, customer_id, service_id, created_by, referer_id, price, sales_price, sales_discount, trn, workflow_history, status, approved_by, approved_at, cancelled_by, cancelled_at, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING id`
	err := h.DB.QueryRow(query, salesLedger.ID, salesLedger.CustomerID, salesLedger.ServiceID, salesLedger.CreatedBy, salesLedger.RefererID,
		salesLedger.Price, salesLedger.SalesPrice, salesLedger.SalesDiscount, salesLedger.TRN,
		salesLedger.WorkflowHistory, salesLedger.Status, nil, nil, nil, nil, salesLedger.CreatedAt, salesLedger.CreatedAt).Scan(&salesLedger.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, salesLedger)
}

func (h *SalesLedgerHandler) ApproveSalesLedger(c *gin.Context) {
	var approval models.SalesLedgerApproval
	if err := c.ShouldBindJSON(&approval); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	approval.ApprovedAt = time.Now()
	recordUpdatedAt := time.Now()
	id := c.Param("id")


	query := `
	UPDATE SalesLedger SET status = 'APPROVED', approved_by = $1, approved_at = $2, trn = $3, updated_at = $5
	WHERE id = $4 AND
	deleted_at IS NULL AND
	status = 'PENDING' AND
	trn IS NULL AND
	approved_by IS NULL AND
	approved_at IS NULL`
	result, err := h.DB.Exec(query, approval.ApproverID, approval.ApprovedAt, approval.TRN, id, recordUpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}

	var workflowHistoryRaw json.RawMessage
	err = h.DB.QueryRow("SELECT workflow_history FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL", id).Scan(&workflowHistoryRaw)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var workflows []models.Workflow
	if len(workflowHistoryRaw) > 0 {
		if err := json.Unmarshal(workflowHistoryRaw, &workflows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal workflow history: " + err.Error()})
			return
		}
	}

	approvalWorkflow := models.Workflow{
		ActionAt: approval.ApprovedAt,
		Comment:  approval.Comment,
		ActorId:  approval.ApproverID,
		Step:     "تایید مالی",
	}
	workflows = append(workflows, approvalWorkflow)

	updatedWorkflowHistory, _ := json.Marshal(workflows)
	wf_update_query := `UPDATE SalesLedger SET  workflow_history = $1 WHERE id = $2 `
	_, err = h.DB.Exec(wf_update_query, updatedWorkflowHistory, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sales ledger approved"})
}

func (h *SalesLedgerHandler) RejectSalesLedger(c *gin.Context) {
	id := c.Param("id")
	var rejection models.SalesLedgerRejection
	if err := c.ShouldBindJSON(&rejection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rejection.RejectedAt = time.Now()
	recordUpdatedAt := time.Now()
	var workflowHistoryRaw json.RawMessage
	err := h.DB.QueryRow("SELECT workflow_history FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL", id).Scan(&workflowHistoryRaw)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var workflows []models.Workflow
	if len(workflowHistoryRaw) > 0 {
		if err := json.Unmarshal(workflowHistoryRaw, &workflows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal workflow history: " + err.Error()})
			return
		}
	}

	rejectionWorkflow := models.Workflow{
		ActionAt: rejection.RejectedAt,
		Comment:  rejection.Reason,
		ActorId:  rejection.RejectedBy,
		Step:     "رد مالی",
	}
	workflows = append(workflows, rejectionWorkflow)

	updatedWorkflowHistory, err := json.Marshal(workflows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal updated workflow history: " + err.Error()})
		return
	}

	query := `UPDATE SalesLedger SET status = 'REJECTED', workflow_history = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL AND status = 'PENDING' AND trn IS NULL AND approved_by IS NULL AND approved_at IS NULL`
	result, err := h.DB.Exec(query, updatedWorkflowHistory, recordUpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sales ledger rejected"})
}

func (h *SalesLedgerHandler) CancelSalesLedger(c *gin.Context) {
	id := c.Param("id")
	var cancellation models.SalesLedgerCancellation
	if err := c.ShouldBindJSON(&cancellation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recordUpdatedAt := time.Now()
	cancellation.CancelledAt = time.Now()
	query := `UPDATE SalesLedger SET status = 'CANCELLED', cancelled_by = $1, cancelled_at = $2, updated_at = $3 WHERE id = $4 AND deleted_at IS NULL AND status <> 'APPROVED' AND status <> 'CANCELLED'`
	result, err := h.DB.Exec(query, cancellation.CancelledBy, cancellation.CancelledAt, recordUpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}

	var workflowHistoryRaw json.RawMessage
	err = h.DB.QueryRow("SELECT workflow_history FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL", id).Scan(&workflowHistoryRaw)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var workflows []models.Workflow
	if len(workflowHistoryRaw) > 0 {
		if err := json.Unmarshal(workflowHistoryRaw, &workflows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal workflow history: " + err.Error()})
			return
		}
	}

	
	cancellationWorkflow := models.Workflow{	
		ActionAt: cancellation.CancelledAt,
		Comment:  cancellation.Reason,
		Step:     "لغو بررسی",
		ActorId:  cancellation.CancelledBy,
	}
	workflows = append(workflows, cancellationWorkflow)

	updatedWorkflowHistory, _ := json.Marshal(workflows)
	wf_update_query := `UPDATE SalesLedger SET  workflow_history = $1 WHERE id = $2 `
	_, err = h.DB.Exec(wf_update_query, updatedWorkflowHistory, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "Sales ledger cancelled"})
}

func (h *SalesLedgerHandler) ResendSalesLedger(c *gin.Context) {
	id := c.Param("id")
	var salesLedger models.SalesLedgerResend
	var createdBy uuid.UUID
	recordUpdatedAt := time.Now()
	if err := c.ShouldBindJSON(&salesLedger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var workflowHistoryRaw json.RawMessage
	err := h.DB.QueryRow("SELECT workflow_history, created_by FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL", id).Scan(&workflowHistoryRaw, &createdBy)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var workflows []models.Workflow
	if len(workflowHistoryRaw) > 0 {
		if err := json.Unmarshal(workflowHistoryRaw, &workflows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal workflow history: " + err.Error()})
			return
		}
	}
	resendWorkflow := models.Workflow{	
		ActionAt: time.Now(),
		Comment:  salesLedger.Comment,
		Step:     "ارسال مجدد به مالی",
		ActorId:  createdBy,
	}
	workflows = append(workflows, resendWorkflow)

	updatedWorkflowHistory, _ := json.Marshal(workflows)
	wf_update_query := `UPDATE SalesLedger SET  workflow_history = $1, status = 'PENDING', price = $2, sales_price = $3, sales_discount = $4, updated_at = $5 WHERE id = $6 AND deleted_at IS NULL AND status = 'REJECTED'`
	_, err = h.DB.Exec(wf_update_query, updatedWorkflowHistory, salesLedger.Price, salesLedger.SalesPrice, salesLedger.SalesDiscount, recordUpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Sales ledger resend"})
}
func (h *SalesLedgerHandler) GetSalesLedger(c *gin.Context) {
	id := c.Param("id")
	var salesLedger models.SalesLedger
	query := `SELECT id, customer_id, service_id, created_by, referer_id, price, sales_price, sales_discount, trn, workflow_history, approved_by, approved_at, cancelled_by, cancelled_at, created_at, status, updated_at, deleted_at
              FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&salesLedger.ID, &salesLedger.CustomerID, &salesLedger.ServiceID, &salesLedger.CreatedBy, &salesLedger.RefererID,
		&salesLedger.Price, &salesLedger.SalesPrice, &salesLedger.SalesDiscount, &salesLedger.TRN,
		&salesLedger.WorkflowHistory, &salesLedger.ApprovedBy,
		&salesLedger.ApprovedAt, &salesLedger.CancelledBy, &salesLedger.CancelledAt, &salesLedger.CreatedAt, &salesLedger.Status, &salesLedger.UpdatedAt, &salesLedger.DeletedAt)
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
	rows, err := h.DB.Query(`SELECT id, customer_id, service_id, created_by, referer_id, price, sales_price, sales_discount, trn, workflow_history, approved_by, approved_at, cancelled_by, cancelled_at, created_at, status, updated_at, deleted_at 
                             FROM SalesLedger WHERE deleted_at IS NULL ORDER BY updated_at DESC`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var salesLedgers []models.SalesLedger
	for rows.Next() {
		var salesLedger models.SalesLedger
		err := rows.Scan(&salesLedger.ID, &salesLedger.CustomerID, &salesLedger.ServiceID, &salesLedger.CreatedBy, &salesLedger.RefererID,
			&salesLedger.Price, &salesLedger.SalesPrice, &salesLedger.SalesDiscount, &salesLedger.TRN,
			&salesLedger.WorkflowHistory, &salesLedger.ApprovedBy, &salesLedger.ApprovedAt, &salesLedger.CancelledBy, &salesLedger.CancelledAt, &salesLedger.CreatedAt, &salesLedger.Status, &salesLedger.UpdatedAt, &salesLedger.DeletedAt)
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
	now := time.Now()
	salesLedger.UpdatedAt = &now

	query := `UPDATE SalesLedger SET customer_id = $1, service_id = $2, created_by = $3, referer_id = $4, price = $5, sales_price = $6, sales_discount = $7, trn = $8, approved_by = $9, approved_at = $10, updated_at = $11
              WHERE id = $12 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, salesLedger.CustomerID, salesLedger.ServiceID, salesLedger.CreatedBy, salesLedger.RefererID,
		salesLedger.Price, salesLedger.SalesPrice, salesLedger.SalesDiscount, salesLedger.TRN,
		salesLedger.ApprovedBy, salesLedger.ApprovedAt, now, salesLedger.ID)
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

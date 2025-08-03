package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"salesman/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type SalesLedgerHandler struct {
	DB *sql.DB
}

func (h *SalesLedgerHandler) CreateSalesLedger(c *gin.Context) {
	var salesLedgerCreate models.SalesLedger2Create
	if err := c.ShouldBindJSON(&salesLedgerCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := h.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction: " + err.Error()})
		return
	}
	defer tx.Rollback()

	salesLedgerID := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()

	var createdByName string
	err = tx.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", salesLedgerCreate.CreatedBy).Scan(&createdByName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get creator name: " + err.Error()})
		return
	}

	initialWorkflow := models.Workflow{
		ActionAt:  createdAt,
		Comment:   "",
		ActorId:   salesLedgerCreate.CreatedBy,
		Step:      "ایجاد فاکتور و ارسال به مالی",
		ActorName: createdByName,
	}
	workflows := []models.Workflow{initialWorkflow}
	workflowHistory, err := json.Marshal(workflows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal workflow history: " + err.Error()})
		return
	}

	query := `INSERT INTO SalesLedger (id, customer_product_id, created_by, marketer_id, payment_method, TRN, workflow_history, commission_level, status, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	err = tx.QueryRow(query, salesLedgerID, salesLedgerCreate.CustomerProductID, salesLedgerCreate.CreatedBy,
		salesLedgerCreate.MarketerID, salesLedgerCreate.PaymentMethod, salesLedgerCreate.TRN,
		workflowHistory, salesLedgerCreate.CommissionLevel, "PENDING", createdAt, updatedAt).Scan(&salesLedgerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sales ledger: " + err.Error()})
		return
	}

	for _, service := range salesLedgerCreate.Services {
		// i know its bad :)))
		var actualPrice float64
		err = tx.QueryRow("SELECT price FROM Services WHERE id = $1 AND deleted_at IS NULL", service.ServiceID).Scan(&actualPrice)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Service not found: " + service.ServiceID.String()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch service price: " + err.Error()})
			}
			return
		}

		ledgerServiceItemID := uuid.New()
		serviceQuery := `INSERT INTO LedgerServiceItems (id, ledger_id, service_id, price, created_at)
                         VALUES ($1, $2, $3, $4, $5) RETURNING id`
		err = tx.QueryRow(serviceQuery, ledgerServiceItemID, salesLedgerID, service.ServiceID, actualPrice, createdAt).Scan(&ledgerServiceItemID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ledger service item: " + err.Error()})
			return
		}
	}

	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": salesLedgerID})
}

func (h *SalesLedgerHandler) ApproveSalesLedger(c *gin.Context) {
	var approval models.SalesLedger2Approval
	if err := c.ShouldBindJSON(&approval); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if approval.CommissionLevel == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Commission level is required"})
		return
	}

	commissionLevel := approval.CommissionLevel
	
	approval.ApprovedAt = time.Now()
	recordUpdatedAt := time.Now()
	id := c.Param("id")

	tx, err := h.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction: " + err.Error()})
		return
	}
	defer tx.Rollback()

	query := `UPDATE SalesLedger SET status = 'APPROVED', approved_by = $1, approved_at = $2, updated_at = $3, commission_level = $4
	WHERE id = $5 AND deleted_at IS NULL AND status = 'PENDING' AND approved_by IS NULL AND approved_at IS NULL`
	result, err := tx.Exec(query, approval.ApproverID, approval.ApprovedAt, recordUpdatedAt, commissionLevel, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found or cannot be approved"})
		return
	}

	var workflowHistoryRaw json.RawMessage
	err = tx.QueryRow("SELECT workflow_history FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL", id).Scan(&workflowHistoryRaw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get workflow history: " + err.Error()})
		return
	}

	var workflows []models.Workflow
	if len(workflowHistoryRaw) > 0 {
		if err := json.Unmarshal(workflowHistoryRaw, &workflows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal workflow history: " + err.Error()})
			return
		}
	}

	var approverName string
	err = tx.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", approval.ApproverID).Scan(&approverName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get approver name: " + err.Error()})
		return
	}

	approvalWorkflow := models.Workflow{
		ActionAt:  approval.ApprovedAt,
		Comment:   approval.Comment,
		ActorId:   approval.ApproverID,
		Step:      "تایید مالی",
		ActorName: approverName,
	}
	workflows = append(workflows, approvalWorkflow)

	updatedWorkflowHistory, err := json.Marshal(workflows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal updated workflow history: " + err.Error()})
		return
	}

	wfUpdateQuery := `UPDATE SalesLedger SET workflow_history = $1 WHERE id = $2`
	_, err = tx.Exec(wfUpdateQuery, updatedWorkflowHistory, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update workflow history: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales ledger approved"})
}

func (h *SalesLedgerHandler) RejectSalesLedger(c *gin.Context) {
	id := c.Param("id")
	var rejection models.SalesLedger2Rejection
	if err := c.ShouldBindJSON(&rejection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rejection.RejectedAt = time.Now()
	recordUpdatedAt := time.Now()

	tx, err := h.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction: " + err.Error()})
		return
	}
	defer tx.Rollback()

	var workflowHistoryRaw json.RawMessage
	err = tx.QueryRow("SELECT workflow_history FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL", id).Scan(&workflowHistoryRaw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get workflow history: " + err.Error()})
		return
	}

	var workflows []models.Workflow
	if len(workflowHistoryRaw) > 0 {
		if err := json.Unmarshal(workflowHistoryRaw, &workflows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal workflow history: " + err.Error()})
			return
		}
	}

	var rejectedBy string
	err = tx.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", rejection.RejectedBy).Scan(&rejectedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rejector name: " + err.Error()})
		return
	}

	rejectionWorkflow := models.Workflow{
		ActionAt:  rejection.RejectedAt,
		Comment:   rejection.Reason,
		ActorId:   rejection.RejectedBy,
		Step:      "رد مالی",
		ActorName: rejectedBy,
	}
	workflows = append(workflows, rejectionWorkflow)

	updatedWorkflowHistory, err := json.Marshal(workflows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal updated workflow history: " + err.Error()})
		return
	}

	query := `UPDATE SalesLedger SET status = 'REJECTED', workflow_history = $1, updated_at = $2 
              WHERE id = $3 AND deleted_at IS NULL AND status = 'PENDING' AND approved_by IS NULL AND approved_at IS NULL`
	result, err := tx.Exec(query, updatedWorkflowHistory, recordUpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found or cannot be rejected"})
		return
	}

	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales ledger rejected"})
}

func (h *SalesLedgerHandler) CancelSalesLedger(c *gin.Context) {
	id := c.Param("id")
	var cancellation models.SalesLedger2Cancellation
	if err := c.ShouldBindJSON(&cancellation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recordUpdatedAt := time.Now()
	cancellation.CancelledAt = time.Now()

	tx, err := h.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction: " + err.Error()})
		return
	}
	defer tx.Rollback()

	query := `UPDATE SalesLedger SET status = 'CANCELLED', cancelled_by = $1, cancelled_at = $2, updated_at = $3 
              WHERE id = $4 AND deleted_at IS NULL AND status <> 'APPROVED' AND status <> 'CANCELLED'`
	result, err := tx.Exec(query, cancellation.CancelledBy, cancellation.CancelledAt, recordUpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found or cannot be cancelled"})
		return
	}

	var workflowHistoryRaw json.RawMessage
	err = tx.QueryRow("SELECT workflow_history FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL", id).Scan(&workflowHistoryRaw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get workflow history: " + err.Error()})
		return
	}

	var workflows []models.Workflow
	if len(workflowHistoryRaw) > 0 {
		if err := json.Unmarshal(workflowHistoryRaw, &workflows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal workflow history: " + err.Error()})
			return
		}
	}

	var cancelledBy string
	err = tx.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", cancellation.CancelledBy).Scan(&cancelledBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get canceller name: " + err.Error()})
		return
	}

	cancellationWorkflow := models.Workflow{
		ActionAt:  cancellation.CancelledAt,
		Comment:   cancellation.Reason,
		Step:      "لغو بررسی",
		ActorId:   cancellation.CancelledBy,
		ActorName: cancelledBy,
	}
	workflows = append(workflows, cancellationWorkflow)

	updatedWorkflowHistory, err := json.Marshal(workflows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal updated workflow history: " + err.Error()})
		return
	}

	wfUpdateQuery := `UPDATE SalesLedger SET workflow_history = $1 WHERE id = $2`
	_, err = tx.Exec(wfUpdateQuery, updatedWorkflowHistory, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update workflow history: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales ledger cancelled"})
}

func (h *SalesLedgerHandler) ResendSalesLedger(c *gin.Context) {
	id := c.Param("id")
	var salesLedger models.SalesLedger2Resend
	var createdBy uuid.UUID
	recordUpdatedAt := time.Now()

	if err := c.ShouldBindJSON(&salesLedger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := h.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction: " + err.Error()})
		return
	}
	defer tx.Rollback()

	var workflowHistoryRaw json.RawMessage
	err = tx.QueryRow("SELECT workflow_history, created_by FROM SalesLedger WHERE id = $1 AND deleted_at IS NULL", id).Scan(&workflowHistoryRaw, &createdBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sales ledger data: " + err.Error()})
		return
	}

	var workflows []models.Workflow
	if len(workflowHistoryRaw) > 0 {
		if err := json.Unmarshal(workflowHistoryRaw, &workflows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal workflow history: " + err.Error()})
			return
		}
	}

	var resendBy string
	err = tx.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", createdBy).Scan(&resendBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get resender name: " + err.Error()})
		return
	}

	resendWorkflow := models.Workflow{
		ActionAt:  time.Now(),
		Comment:   salesLedger.Comment,
		Step:      "ارسال مجدد به مالی",
		ActorId:   createdBy,
		ActorName: resendBy,
	}
	workflows = append(workflows, resendWorkflow)

	updatedWorkflowHistory, err := json.Marshal(workflows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal updated workflow history: " + err.Error()})
		return
	}

	wfUpdateQuery := `UPDATE SalesLedger SET workflow_history = $1, status = 'PENDING', payment_method = $2, TRN = $3, updated_at = $4 
                      WHERE id = $5 AND deleted_at IS NULL AND status = 'REJECTED'`
	result, err := tx.Exec(wfUpdateQuery, updatedWorkflowHistory, salesLedger.PaymentMethod, salesLedger.TRN, recordUpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sales ledger: " + err.Error()})
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found or cannot be resent"})
		return
	}

	_, err = tx.Exec("DELETE FROM LedgerServiceItems WHERE ledger_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete existing service items: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales ledger resent"})
}

func (h *SalesLedgerHandler) GetSalesLedger(c *gin.Context) {
	id := c.Param("id")
	var salesLedger models.SalesLedger2

	query := `SELECT sl.id, sl.customer_product_id, sl.created_by, sl.approved_by, sl.cancelled_by, sl.marketer_id,
			sl.status, sl.payment_method, sl.TRN, sl.workflow_history, sl.commission_level, sl.approved_at,
			sl.cancelled_at, sl.created_at, sl.status, sl.updated_at, sl.deleted_at,
			u.first_name, u.last_name, u2.first_name as created_by_name, u2.last_name as created_by_last_name,
			u3.first_name as referer_name, u3.last_name as referer_last_name
	FROM SalesLedger sl
	INNER JOIN CompanyUserProducts cup ON sl.customer_product_id = cup.id
	INNER JOIN CompanyUsers cu ON cup.company_user_id = cu.id
	INNER JOIN Users u ON cu.user_id = u.id
	INNER JOIN Users u2 ON sl.created_by = u2.id
	LEFT JOIN Users u3 ON sl.marketer_id = u3.id
	WHERE sl.id = $1 AND sl.deleted_at IS NULL`

	err := h.DB.QueryRow(query, id).Scan(
		&salesLedger.ID, &salesLedger.CustomerProductID, &salesLedger.CreatedBy, &salesLedger.ApprovedBy,
		&salesLedger.CancelledBy, &salesLedger.MarketerID, &salesLedger.Status, &salesLedger.PaymentMethod,
		&salesLedger.TRN, &salesLedger.WorkflowHistory, &salesLedger.CommissionLevel, &salesLedger.ApprovedAt,
		&salesLedger.CancelledAt, &salesLedger.CreatedAt, &salesLedger.Status, &salesLedger.UpdatedAt, &salesLedger.DeletedAt,
		&salesLedger.CustomerFirstName, &salesLedger.CustomerLastName, &salesLedger.CreatedByName, &salesLedger.CreatedByLastName,
		&salesLedger.RefererName, &salesLedger.RefererLastName)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales ledger entry not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	serviceItemsQuery := `SELECT lsi.id, lsi.ledger_id, lsi.service_id, lsi.price, lsi.created_at,
	                       s.title as service_title, c.title as company_title
	                       FROM LedgerServiceItems lsi
	                       INNER JOIN Services s ON lsi.service_id = s.id
	                       INNER JOIN Products p ON s.product_id = p.id
	                       INNER JOIN Companies c ON p.company_id = c.id
	                       WHERE lsi.ledger_id = $1 AND lsi.deleted_at IS NULL`

	serviceRows, err := h.DB.Query(serviceItemsQuery, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get service items: " + err.Error()})
		return
	}
	defer serviceRows.Close()

	var serviceItems []models.LedgerServiceItems
	for serviceRows.Next() {
		var item models.LedgerServiceItems
		var serviceTitle, companyTitle string
		err := serviceRows.Scan(&item.ID, &item.LedgerID, &item.ServiceID, &item.Price, &item.CreatedAt, &serviceTitle, &companyTitle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan service item: " + err.Error()})
			return
		}
		serviceItems = append(serviceItems, item)
	}

	// Add service items to response
	response := gin.H{
		"sales_ledger":  salesLedger,
		"service_items": serviceItems,
	}

	c.JSON(http.StatusOK, response)
}

func (h *SalesLedgerHandler) GetSalesLedgers(c *gin.Context) {
	createdByParam := c.Query("created_by")
	var rows *sql.Rows
	var err error

	if createdByParam != "" {
		createdBy, err := uuid.Parse(createdByParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid created_by parameter"})
			return
		}
		rows, err = h.DB.Query(`SELECT sl.id, sl.customer_product_id, sl.created_by, sl.approved_by, sl.cancelled_by, sl.marketer_id,
			sl.status, sl.payment_method, sl.TRN, sl.workflow_history, sl.commission_level, sl.approved_at,
			sl.cancelled_at, sl.created_at, sl.status, sl.updated_at, sl.deleted_at,
			u.first_name, u.last_name, coalesce(total_price, 0)
		FROM SalesLedger sl
		INNER JOIN CompanyUserProducts cup ON sl.customer_product_id = cup.id
		INNER JOIN CompanyUsers cu ON cup.company_user_id = cu.id
		INNER JOIN Users u ON cu.user_id = u.id
		LEFT JOIN InvoiceTotalPrice itp ON sl.id = itp.ledger_id
		WHERE sl.deleted_at IS NULL AND sl.created_by = $1
		ORDER BY sl.updated_at DESC`, createdBy)
	} else {
		rows, err = h.DB.Query(`SELECT sl.id, sl.customer_product_id, sl.created_by, sl.approved_by, sl.cancelled_by, sl.marketer_id,
			sl.status, sl.payment_method, sl.TRN, sl.workflow_history, sl.commission_level, sl.approved_at,
			sl.cancelled_at, sl.created_at, sl.status, sl.updated_at, sl.deleted_at,
			u.first_name, u.last_name, coalesce(total_price, 0)
		FROM SalesLedger sl
		INNER JOIN CompanyUserProducts cup ON sl.customer_product_id = cup.id
		INNER JOIN CompanyUsers cu ON cup.company_user_id = cu.id
		INNER JOIN Users u ON cu.user_id = u.id
		LEFT JOIN InvoiceTotalPrice itp ON sl.id = itp.ledger_id
		WHERE sl.deleted_at IS NULL
		ORDER BY sl.updated_at DESC`)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var salesLedgers []models.SalesLedger2
	for rows.Next() {
		var salesLedger models.SalesLedger2
		err := rows.Scan(
			&salesLedger.ID, &salesLedger.CustomerProductID, &salesLedger.CreatedBy, &salesLedger.ApprovedBy,
			&salesLedger.CancelledBy, &salesLedger.MarketerID, &salesLedger.Status, &salesLedger.PaymentMethod,
			&salesLedger.TRN, &salesLedger.WorkflowHistory, &salesLedger.CommissionLevel, &salesLedger.ApprovedAt,
			&salesLedger.CancelledAt, &salesLedger.CreatedAt, &salesLedger.Status, &salesLedger.UpdatedAt, &salesLedger.DeletedAt,
			&salesLedger.CustomerFirstName, &salesLedger.CustomerLastName,
			&salesLedger.TotalPrice,
		)
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

func (h *SalesLedgerHandler) GetNewSalesLedgers(c *gin.Context) {
	statusParam := c.Query("status")
	operatorId := c.Query("operator_id")

	statuses := []string{}
	if statusParam != "" {
		statuses = strings.Split(statusParam, ",")
	}

	var count int
	var err error

	if operatorId != "" {
		query := `SELECT COUNT(*) FROM SalesLedger WHERE status = ANY($1) AND created_by = $2 AND deleted_at IS NULL`
		err = h.DB.QueryRow(query, pq.Array(statuses), operatorId).Scan(&count)
	} else {
		query := `SELECT COUNT(*) FROM SalesLedger WHERE status = ANY($1) AND deleted_at IS NULL`
		err = h.DB.QueryRow(query, pq.Array(statuses)).Scan(&count)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *SalesLedgerHandler) UpdateSalesLedger(c *gin.Context) {
	idStr := c.Param("id")

	salesLedgerID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sales ledger ID format"})
		return
	}

	var salesLedger models.SalesLedger2
	if err := c.ShouldBindJSON(&salesLedger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesLedger.ID = salesLedgerID
	now := time.Now()
	salesLedger.UpdatedAt = &now

	query := `UPDATE SalesLedger SET customer_product_id = $1, created_by = $2, marketer_id = $3, 
              payment_method = $4, TRN = $5, commission_level = $6, approved_by = $7, approved_at = $8, updated_at = $9
              WHERE id = $10 AND deleted_at IS NULL`
	result, err := h.DB.Exec(query, salesLedger.CustomerProductID, salesLedger.CreatedBy, salesLedger.MarketerID,
		salesLedger.PaymentMethod, salesLedger.TRN, salesLedger.CommissionLevel,
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

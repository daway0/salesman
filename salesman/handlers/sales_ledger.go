package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
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
	var salesLedger models.SalesLedger
	if err := c.ShouldBindJSON(&salesLedger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesLedger.ID = uuid.New()
	salesLedger.CreatedAt = time.Now()

	res := h.DB.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", salesLedger.CreatedBy).Scan(&salesLedger.CreatedByName)
	if res != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error()})
		return
	}

	initialWorkflow := models.Workflow{
		ActionAt:  salesLedger.CreatedAt,
		Comment:   "",
		ActorId:   salesLedger.CreatedBy,
		Step:      "ایجاد فاکتور و ارسال به مالی",
		ActorName: salesLedger.CreatedByName,
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

	log.Println(approval.TRN)

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

	var ApproverName string

	res := h.DB.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", approval.ApproverID).Scan(&ApproverName)
	if res != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error()})
		return
	}
	approvalWorkflow := models.Workflow{
		ActionAt:  approval.ApprovedAt,
		Comment:   approval.Comment,
		ActorId:   approval.ApproverID,
		Step:      "تایید مالی",
		ActorName: ApproverName,
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

	var RejectedBy string
	res := h.DB.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", rejection.RejectedBy).Scan(&RejectedBy)
	if res != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error()})
		return
	}
	rejectionWorkflow := models.Workflow{
		ActionAt:  rejection.RejectedAt,
		Comment:   rejection.Reason,
		ActorId:   rejection.RejectedBy,
		Step:      "رد مالی",
		ActorName: RejectedBy,
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

	var CancelledBy string
	res := h.DB.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", cancellation.CancelledBy).Scan(&CancelledBy)
	if res != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error()})
		return
	}
	cancellationWorkflow := models.Workflow{
		ActionAt:  cancellation.CancelledAt,
		Comment:   cancellation.Reason,
		Step:      "لغو بررسی",
		ActorId:   cancellation.CancelledBy,
		ActorName: CancelledBy,
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
	var ResendBy string
	res := h.DB.QueryRow("SELECT concat(first_name, ' ', last_name) as actor_name FROM Users WHERE id = $1", createdBy).Scan(&ResendBy)
	if res != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error()})
		return
	}
	resendWorkflow := models.Workflow{
		ActionAt:  time.Now(),
		Comment:   salesLedger.Comment,
		Step:      "ارسال مجدد به مالی",
		ActorId:   createdBy,
		ActorName: ResendBy,
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
	query := `SELECT sl.id,
			sl.customer_id,
			sl.service_id,
			sl.created_by,
			sl.referer_id,
			sl.price,
			sl.sales_price,
			sl.sales_discount,
			sl.trn,
			sl.workflow_history,
			sl.approved_by,
			sl.approved_at,
			sl.cancelled_by,
			sl.cancelled_at,
			sl.created_at,
			sl.status,
			sl.updated_at,
			sl.deleted_at,
			s.title,
			c.title,
			u.first_name,
			u.last_name,
			u2.first_name as created_by_name,
			u2.last_name as created_by_last_name,
			u3.first_name as referer_name,
			u3.last_name as referer_last_name
	FROM SalesLedger sl
	INNER JOIN services s on sl.service_id = s.id
	INNER JOIN companies c on c.id = s.company_id
	INNER JOIN users u on sl.customer_id = u.id
	INNER JOIN users u2 on sl.created_by = u2.id
	LEFT JOIN users u3 on sl.referer_id = u3.id
	WHERE sl.id = $1
  AND sl.deleted_at IS NULL`
	err := h.DB.QueryRow(query, id).Scan(&salesLedger.ID, &salesLedger.CustomerID, &salesLedger.ServiceID, &salesLedger.CreatedBy, &salesLedger.RefererID,
		&salesLedger.Price, &salesLedger.SalesPrice, &salesLedger.SalesDiscount, &salesLedger.TRN,
		&salesLedger.WorkflowHistory, &salesLedger.ApprovedBy,
		&salesLedger.ApprovedAt, &salesLedger.CancelledBy, &salesLedger.CancelledAt, &salesLedger.CreatedAt, &salesLedger.Status, &salesLedger.UpdatedAt, &salesLedger.DeletedAt,
		&salesLedger.ServiceTitle, &salesLedger.CompanyTitle, &salesLedger.CustomerFirstName, &salesLedger.CustomerLastName,
		&salesLedger.CreatedByName, &salesLedger.CreatedByLastName, &salesLedger.RefererName, &salesLedger.RefererLastName)
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
	createdByParam := c.Query("created_by")
	var rows *sql.Rows
	var err error

	if createdByParam != "" {
		createdBy, err := uuid.Parse(createdByParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid created_by parameter"})
			return
		}
		rows, err = h.DB.Query(`SELECT sl.id, sl.customer_id, sl.service_id, sl.created_by, sl.referer_id, sl.price, sl.sales_price, sl.sales_discount, sl.trn, sl.workflow_history, sl.approved_by, sl.approved_at, sl.cancelled_by, sl.cancelled_at, sl.created_at, sl.status, sl.updated_at, sl.deleted_at,
       s.title, c.title, u.first_name, u.last_name
FROM SalesLedger sl
INNER JOIN services s ON sl.service_id = s.id
INNER JOIN companies c ON c.id = s.company_id
INNER JOIN users u ON sl.customer_id = u.id
WHERE sl.deleted_at IS NULL
AND sl.created_by = $1
ORDER BY sl.updated_at DESC`, createdBy)
	} else {
		rows, err = h.DB.Query(`SELECT sl.id, sl.customer_id, sl.service_id, sl.created_by, sl.referer_id, sl.price, sl.sales_price, sl.sales_discount, sl.trn, sl.workflow_history, sl.approved_by, sl.approved_at, sl.cancelled_by, sl.cancelled_at, sl.created_at, sl.status, sl.updated_at, sl.deleted_at,
       s.title, c.title, u.first_name, u.last_name
FROM SalesLedger sl
INNER JOIN services s ON sl.service_id = s.id
INNER JOIN companies c ON c.id = s.company_id
INNER JOIN users u ON sl.customer_id = u.id
WHERE sl.deleted_at IS NULL
ORDER BY sl.updated_at DESC`)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var salesLedgers []models.SalesLedger
	for rows.Next() {
		var salesLedger models.SalesLedger
		err := rows.Scan(
			&salesLedger.ID, &salesLedger.CustomerID, &salesLedger.ServiceID, &salesLedger.CreatedBy, &salesLedger.RefererID,
			&salesLedger.Price, &salesLedger.SalesPrice, &salesLedger.SalesDiscount, &salesLedger.TRN,
			&salesLedger.WorkflowHistory, &salesLedger.ApprovedBy, &salesLedger.ApprovedAt, &salesLedger.CancelledBy, &salesLedger.CancelledAt,
			&salesLedger.CreatedAt, &salesLedger.Status, &salesLedger.UpdatedAt, &salesLedger.DeletedAt,
			&salesLedger.ServiceTitle, &salesLedger.CompanyTitle, &salesLedger.CustomerFirstName, &salesLedger.CustomerLastName,
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

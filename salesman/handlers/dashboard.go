package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"salesman/models"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	DB *sql.DB
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	username := c.Query("user_id")

	var dashboard models.Dashboard

	err := h.DB.QueryRow(`
		SELECT COUNT(*) FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at::date = CURRENT_DATE
	`).Scan(&dashboard.TotalApprovedSalesCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total approved sales count"})
		return
	}

	err = h.DB.QueryRow(`
		SELECT COUNT(*) FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at::date = CURRENT_DATE
		  AND created_by = $1
	`, username).Scan(&dashboard.UserApprovedSalesCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user approved sales count"})
		return
	}

	err = h.DB.QueryRow(`
		SELECT COALESCE(SUM(sales_price), 0) FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at::date = CURRENT_DATE
	`).Scan(&dashboard.TotalApprovedSalesAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total approved sales amount"})
		return
	}

	err = h.DB.QueryRow(`
		SELECT COALESCE(SUM(sales_price), 0) FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at::date = CURRENT_DATE
		  AND created_by = $1
	`, username).Scan(&dashboard.UserApprovedSalesAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user approved sales amount"})
		return
	}

	err = h.DB.QueryRow(`
		SELECT COUNT(*) FROM salesledger
		WHERE (status = 'REJECTED' OR status = 'PENDING')
	`).Scan(&dashboard.TotalUnapprovedSalesCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total unapproved sales count"})
		return
	}

	err = h.DB.QueryRow(`
		SELECT COUNT(*) FROM salesledger
		WHERE (status = 'REJECTED' OR status = 'PENDING')
		  AND created_by = $1
	`, username).Scan(&dashboard.UserUnapprovedSalesCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user unapproved sales count"})
		return
	}

	startDate := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")

	err = h.DB.QueryRow(`
		SELECT COUNT(*) FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at::date BETWEEN $1 AND $2
	`, startDate, endDate).Scan(&dashboard.TotalApprovedSalesCountInRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total approved sales count in range"})
		return
	}

	err = h.DB.QueryRow(`
		SELECT COALESCE(SUM(sales_price), 0) FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at::date BETWEEN $1 AND $2
	`, startDate, endDate).Scan(&dashboard.TotalApprovedSalesAmountInRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total approved sales amount in range"})
		return
	}

	err = h.DB.QueryRow(`
		SELECT COUNT(*) FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at::date BETWEEN $1 AND $2
		  AND created_by = $3
	`, startDate, endDate, username).Scan(&dashboard.UserApprovedSalesCountInRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user approved sales count in range"})
		return
	}

	err = h.DB.QueryRow(`
		SELECT COALESCE(SUM(sales_price), 0) FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at::date BETWEEN $1 AND $2
		  AND created_by = $3
	`, startDate, endDate, username).Scan(&dashboard.UserApprovedSalesAmountInRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user approved sales amount in range"})
		return
	}

	rows, err := h.DB.Query(`
		SELECT u.first_name, u.last_name, rs.total_sales_price, rs.count_sales
		FROM (SELECT created_by, SUM(sales_price) AS total_sales_price, COUNT(sales_price) AS count_sales
			FROM salesledger
			WHERE status = 'APPROVED'
			GROUP BY created_by
			ORDER BY SUM(sales_price) DESC
			LIMIT 5) rs
				INNER JOIN users u ON u.id = rs.created_by
		ORDER BY total_sales_price desc
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sales MVP"})
		return
	}
	defer rows.Close()
	var salesMVP []models.SalesMVP
	for rows.Next() {
		var mvp models.SalesMVP
		err := rows.Scan(&mvp.Firstname, &mvp.Lastname, &mvp.SalesAmount, &mvp.SalesCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan sales MVP"})
			return
		}
		salesMVP = append(salesMVP, mvp)
	}
	dashboard.SalesMVP = salesMVP

	startDate = time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	endDate = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	res, err := h.DB.Query(`
		SELECT count(*)
		FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at BETWEEN $1 AND $2
	`, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all-users chart data"})
		return
	}
	defer res.Close()

	rows, err = h.DB.Query(`
		SELECT approved_at::date, COALESCE(SUM(sales_price), 0)
		FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at BETWEEN $1 AND $2
		GROUP BY approved_at::date
		ORDER BY approved_at::date
	`, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all-users chart data"})
		return
	}
	defer rows.Close()
	var chartDataAll []models.ChartData
	for rows.Next() {
		var cd models.ChartData
		err := rows.Scan(&cd.Date, &cd.SalesAmount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan all-users chart data"})
			return
		}
		chartDataAll = append(chartDataAll, cd)
	}
	dashboard.ChartDataAll = chartDataAll

	rows, err = h.DB.Query(`
		SELECT approved_at::date, COALESCE(SUM(sales_price), 0)
		FROM salesledger
		WHERE status = 'APPROVED'
		  AND approved_at BETWEEN $1 AND $2
		  AND created_by = $3
		GROUP BY approved_at::date
		ORDER BY approved_at::date
	`, startDate, endDate, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user chart data"})
		return
	}
	defer rows.Close()
	var chartData []models.ChartData
	for rows.Next() {
		var cd models.ChartData
		err := rows.Scan(&cd.Date, &cd.SalesAmount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user chart data"})
			return
		}
		chartData = append(chartData, cd)
	}
	dashboard.ChartDataUser = chartData

	var todayMVP models.SalesMVP
	err = h.DB.QueryRow(`
		SELECT u.first_name, u.last_name, rs.total_sales_price, rs.count_sales
		FROM (
			SELECT created_by, SUM(sales_price) AS total_sales_price, COUNT(sales_price) AS count_sales
			FROM salesledger
			WHERE status = 'APPROVED' AND approved_at::date = CURRENT_DATE
			GROUP BY created_by
			ORDER BY SUM(sales_price) DESC
			LIMIT 1
		) rs
		INNER JOIN users u ON u.id = rs.created_by
	`).Scan(&todayMVP.Firstname, &todayMVP.Lastname, &todayMVP.SalesAmount, &todayMVP.SalesCount)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get today's MVP"})
		return
	}
	dashboard.TodayMVP = todayMVP

	c.JSON(http.StatusOK, dashboard)
}

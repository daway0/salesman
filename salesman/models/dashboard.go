package models

type SalesMVP struct {
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
	SalesCount  int     `json:"sales_count"`
	SalesAmount float64 `json:"sales_amount"`
}

type ChartData struct {
	Date        string  `json:"date"`
	SalesCount  int     `json:"sales_count"`
	SalesAmount float64 `json:"sales_amount"`
}

type Dashboard struct {
	// cards
	TotalApprovedSalesCount   int     `json:"total_approved_sales_count"`
	TotalApprovedSalesAmount  float64 `json:"total_approved_sales_amount"`
	TotalUnapprovedSalesCount int     `json:"total_unapproved_sales_count"`

	UserApprovedSalesCount   int     `json:"user_approved_sales_count"`
	UserApprovedSalesAmount  float64 `json:"user_approved_sales_amount"`
	UserUnapprovedSalesCount int     `json:"user_unapproved_sales_count"`

	// mvp table
	SalesMVP []SalesMVP `json:"sales_mvp"`

	// chart
	ChartData []ChartData `json:"chart_data"`

	// chart for all users
	ChartDataAll []ChartData `json:"chart_data_all"`

	// approved sales count in range
	UserApprovedSalesCountInRange   int     `json:"user_approved_sales_count_in_range"`
	UserApprovedSalesAmountInRange  float64 `json:"user_approved_sales_amount_in_range"`
	TotalApprovedSalesCountInRange  int     `json:"total_approved_sales_count_in_range"`
	TotalApprovedSalesAmountInRange float64 `json:"total_approved_sales_amount_in_range"`

	// today mvp operator
	TodayMVP SalesMVP `json:"today_mvp"`
}

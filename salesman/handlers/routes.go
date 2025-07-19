package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	userHandler := &UserHandler{DB: db}
	companyHandler := &CompanyHandler{DB: db}
	serviceHandler := &ServiceHandler{DB: db}
	roleHandler := &RoleHandler{DB: db}
	permissionHandler := &PermissionHandler{DB: db}
	salesLedgerHandler := &SalesLedgerHandler{DB: db}
	userRoleHandler := &UserRoleHandler{DB: db}
	rolePermissionHandler := &RolePermissionHandler{DB: db}

	api := r.Group("/api")
	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users", userHandler.GetUsers)
		api.GET("/users/:id", userHandler.GetUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
		api.POST("/users/:id/role/:role_id", userHandler.AddUserRole)
		api.DELETE("/users/:id/role/:role_id", userHandler.RemoveUserRole)

		api.POST("/companies", companyHandler.CreateCompany)
		api.GET("/companies", companyHandler.GetCompanies)
		api.GET("/companies/:id", companyHandler.GetCompany)
		api.PUT("/companies/:id", companyHandler.UpdateCompany)
		api.DELETE("/companies/:id", companyHandler.DeleteCompany)

		api.POST("/services", serviceHandler.CreateService)
		api.GET("/services", serviceHandler.GetServices)
		api.GET("/services/:id", serviceHandler.GetService)
		api.PUT("/services/:id", serviceHandler.UpdateService)
		api.DELETE("/services/:id", serviceHandler.DeleteService)

		api.POST("/sales-ledgers", salesLedgerHandler.CreateSalesLedger)
		api.GET("/sales-ledgers", salesLedgerHandler.GetSalesLedgers)
		api.GET("/sales-ledgers/:id", salesLedgerHandler.GetSalesLedger)
		api.PUT("/sales-ledgers/:id", salesLedgerHandler.UpdateSalesLedger)
		api.DELETE("/sales-ledgers/:id", salesLedgerHandler.DeleteSalesLedger)

		
		// rbac
		api.POST("/roles", roleHandler.CreateRole)
		api.GET("/roles", roleHandler.GetRoles)
		api.GET("/roles/:id", roleHandler.GetRole)
		api.PUT("/roles/:id", roleHandler.UpdateRole)
		api.DELETE("/roles/:id", roleHandler.DeleteRole)

		api.POST("/permissions", permissionHandler.CreatePermission)
		api.GET("/permissions", permissionHandler.GetPermissions)
		api.GET("/permissions/:id", permissionHandler.GetPermission)
		api.PUT("/permissions/:id", permissionHandler.UpdatePermission)
		api.DELETE("/permissions/:id", permissionHandler.DeletePermission)


		api.POST("/user-roles", userRoleHandler.CreateUserRole)
		api.GET("/user-roles", userRoleHandler.GetUserRoles)
		api.GET("/user-roles/:id", userRoleHandler.GetUserRole)
		api.GET("/users/:id/roles", roleHandler.GetUserRoles)
		// api.PUT("/user-roles/:id", userRoleHandler.UpdateUserRole)
		api.DELETE("/user-roles/:id", userRoleHandler.DeleteUserRole)

		api.POST("/role-permissions", rolePermissionHandler.CreateRolePermission)
		api.GET("/role-permissions", rolePermissionHandler.GetRolePermissions)
		api.GET("/role-permissions/:id", rolePermissionHandler.GetRolePermission)
		api.GET("/roles/:id/permissions", rolePermissionHandler.GetRolePermissionsByRoleID)
		api.PUT("/role-permissions/:id", rolePermissionHandler.UpdateRolePermission)
		api.DELETE("/role-permissions/:id", rolePermissionHandler.DeleteRolePermission)
	}
}

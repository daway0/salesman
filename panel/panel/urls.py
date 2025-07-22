from django.urls import path, re_path
from . import views

urlpatterns = [
    # Users RBAC
    path('users/create/', views.users_create_update, name='users_create_update'),
    path('users/edit/<uuid:id>/', views.users_create_update, name='users_create_update'),
    path('users/', views.users_list, name='users_list'),

    # Customers
    path('companies/create/', views.companies_create_update, name='companies_create_update'),
    path('companies/edit/<uuid:id>/', views.companies_create_update, name='companies_create_update'),
    path('companies/', views.companies_list, name='companies_list'),

    # Services
    path('services/create/', views.services_create_update, name='services_create_update'),
    path('services/edit/<uuid:id>/', views.services_create_update, name='services_create_update'),
    path('services/', views.services_list, name='services_list'),

    # Sales Ledgers
    path('sales-ledger/create/', views.sales_ledger_create_update, name='sales_ledger_create_update'),
    path('sales-ledger/edit/<uuid:id>/', views.sales_ledger_create_update, name='sales_ledger_create_update'),
    path('sales-ledger/', views.sales_ledger_list, name='sales_ledger_list'),

    # Roles
    path('roles/create/', views.roles_create_update, name='roles_create_update'),
    path('roles/edit/<uuid:id>/', views.roles_create_update, name='roles_create_update'),
    path('roles/', views.roles_list, name='roles_list'),

    # Dashboard
    path('', views.dashboard, name='dashboard'),

    # Login
    path('login/', views.login, name='login'),
    # API Gateway
    re_path(r'^api/salesman/(?P<path>.*)$', views.salesman_gateway, name='salesman_gateway'),
]

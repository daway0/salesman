from django.shortcuts import render
from django.http import JsonResponse
import requests


ADMIN_PERMISSIONS = [
        'CREATE_INVOICE',
        'VIEW_OWN_INVOICES',
        'VIEW_OWN_INVOICES_LIST',
        'VIEW_ALL_INVOICES',
        'VIEW_ALL_INVOICES_LIST',
        'DELETE_OWN_INVOICES',
        'DELETE_ALL_INVOICES',
        'APPROVE_ALL_INVOICES',
        'REJECT_ALL_INVOICES',
        'CANCEL_OWN_INVOICES',
        'CANCEL_ALL_INVOICES',
        
        'VIEW_ALL_USERS_LIST',
        'VIEW_ALL_USERS',
        'CREATE_ALL_USERS',
        'UPDATE_ALL_USERS',
        'DELETE_ALL_USERS',
        'VIEW_USER_ROLES',
        'ADD_USER_ROLE',
        'DELETE_USER_ROLE',
        
        'VIEW_ALL_SERVICES_LIST',
        'VIEW_ALL_SERVICES',
        'CREATE_ALL_SERVICES',
        'UPDATE_ALL_SERVICES',
        'DELETE_ALL_SERVICES',
        
        'VIEW_ALL_COMPANIES_LIST',
        'VIEW_ALL_COMPANIES',
        'CREATE_ALL_COMPANIES',
        'UPDATE_ALL_COMPANIES',
        'DELETE_ALL_COMPANIES',
        
        'VIEW_ALL_ROLES_LIST',
        'VIEW_ALL_ROLES',
        'CREATE_ALL_ROLES',
        'UPDATE_ALL_ROLES',
        'DELETE_ALL_ROLES',
        'VIEW_ROLE_PERMISSIONS',
        'ADD_ROLE_PERMISSION',
        'DELETE_ROLE_PERMISSION',
        
        'VIEW_ALL_LOGS',

        'VIEW_DASHBOARD',
        'VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_COUNT',
        'VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_AMMOUNT',
        'VIEW_DASHBOARD_CARD_ALL_TODAY_UNAPPROVED_COUNT',
        'VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_COUNT',
        'VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_AMMOUNT',
        'VIEW_DASHBOARD_CARD_USER_TODAY_UNAPPROVED_COUNT',
        'VIEW_TODAY_MVP_OPERATOR',
        'VIEW_SALES_MVP_LIST',
        'VIEW_SALES_HISTORY'
]

OPERATOR_PERMISSIONS = [
        'CREATE_INVOICE',
        'VIEW_OWN_INVOICES',
        'VIEW_OWN_INVOICES_LIST',
        'DELETE_OWN_INVOICES',
        'CANCEL_OWN_INVOICES',
        'VIEW_ALL_SERVICES_LIST',
        'VIEW_ALL_COMPANIES_LIST'
]
FINANANCE_PERMISSIONS = [
        'VIEW_ALL_INVOICES_LIST',

        'VIEW_ALL_INVOICES',
        'DELETE_ALL_INVOICES',
        'APPROVE_ALL_INVOICES',
        'REJECT_ALL_INVOICES',
        'CANCEL_ALL_INVOICES',
        'VIEW_ALL_SERVICES_LIST',
        'VIEW_ALL_COMPANIES_LIST'
]

USER_PERMISSIONS = ADMIN_PERMISSIONS

# proxy to salesman service
def salesman_gateway(request, path):
    SALESMAN_BASE_URL = "http://salesman:8080/api"

    target_path = path
    target_url = f"{SALESMAN_BASE_URL}/{target_path}"

    headers = {}
    for k, v in request.headers.items():
        if k.lower() not in ["host", "cookie", "content-length"]:
            headers[k] = v

    try:
        resp = requests.request(
            method=request.method,
            url=target_url,
            params=request.GET.dict(),
            data=request.body if request.body else None,
            headers=headers,
            cookies=request.COOKIES,
            allow_redirects=True,
            files=request.FILES if request.FILES else None,
        )
        try:
            data = resp.json()
        except ValueError:
            data = {"error": "Invalid JSON from upstream service"}
        return JsonResponse(data, status=resp.status_code, safe=False)
    except requests.RequestException as e:
        return JsonResponse({"error": f"Error forwarding request: {e}"}, status=502, safe=False)

# users
def users_create_update(request, id=None):
    context = {
        "edit": False,
        "user_id": None,
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    if id:
        context["edit"] = True
        context["user_id"] = id
    return render(request, 'pages/users_create_update.html', context)


def users_list(request):
    context = {
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    return render(request, 'pages/users_list.html', context)


# roles
def roles_create_update(request, id=None):
    context = {
        "edit": False,
        "role_id": None,
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    if id:
        context["edit"] = True
        context["role_id"] = id
    return render(request, 'pages/roles_create_update.html', context)


def roles_list(request):
    context = {
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    return render(request, 'pages/roles_list.html', context)


# companies
def companies_create_update(request, id=None):
    context = {
        "edit": False,
        "company_id": None,
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    if id:
        context["edit"] = True
        context["company_id"] = id
    return render(request, 'pages/companies_create_update.html', context)

def companies_list(request):
    context = {
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    return render(request, 'pages/companies_list.html', context)

# services
def services_create_update(request, id=None):
    context = {
        "edit": False,
        "service_id": None,
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    if id:
        context["edit"] = True
        context["service_id"] = id
    return render(request, 'pages/services_create_update.html', context)

def services_list(request):
    context = {
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    return render(request, 'pages/services_list.html', context)

# sales ledgers
def sales_ledger_create_update(request, id=None):
    context = {
        "edit": False,
        "sales_ledger_id": None,
        "current_user": "00000000-0000-0000-0000-000000000036",
        "permissions": USER_PERMISSIONS,
    }
    if id:
        context["edit"] = True
        context["sales_ledger_id"] = id
    return render(request, 'pages/sales_ledger_create_update.html', context)    

def sales_ledger_list(request):
    context = {
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    return render(request, 'pages/sales_ledger_list.html', context)

# dashboard
def dashboard(request):
    context = {
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    return render(request, 'pages/dashboard.html', context)


def login(request):
    context = {
        "current_user": None,
        "permissions": USER_PERMISSIONS,
    }
    return render(request, 'pages/login.html', context)

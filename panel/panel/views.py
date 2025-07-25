from django.shortcuts import render
from django.http import JsonResponse
import requests
from .decorators import jwt_required



@jwt_required
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

@jwt_required
def users_create_update(request, id=None):
    context = {
        "edit": False,
        "user_id": None,
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    if id:
        context["edit"] = True
        context["user_id"] = id
    return render(request, 'pages/users_create_update.html', context)


@jwt_required
def users_list(request):
    context = {
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    return render(request, 'pages/users_list.html', context)


@jwt_required
def roles_create_update(request, id=None):
    context = {
        "edit": False,
        "role_id": None,
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    if id:
        context["edit"] = True
        context["role_id"] = id
    return render(request, 'pages/roles_create_update.html', context)


@jwt_required
def roles_list(request):
    context = {
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    return render(request, 'pages/roles_list.html', context)


@jwt_required
def companies_create_update(request, id=None):
    context = {
        "edit": False,
        "company_id": None,
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    if id:
        context["edit"] = True
        context["company_id"] = id
    return render(request, 'pages/companies_create_update.html', context)

@jwt_required
def companies_list(request):
    context = {
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    return render(request, 'pages/companies_list.html', context)

@jwt_required
def services_create_update(request, id=None):
    context = {
        "edit": False,
        "service_id": None,
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    if id:
        context["edit"] = True
        context["service_id"] = id
    return render(request, 'pages/services_create_update.html', context)

@jwt_required
def services_list(request):
    context = {
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    return render(request, 'pages/services_list.html', context)

@jwt_required
def sales_ledger_create_update(request, id=None):
    context = {
        "edit": False,
        "sales_ledger_id": None,
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    if id:
        context["edit"] = True
        context["sales_ledger_id"] = id
    return render(request, 'pages/sales_ledger_create_update.html', context)    

@jwt_required
def sales_ledger_list(request):
    context = {
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    return render(request, 'pages/sales_ledger_list.html', context)


@jwt_required
def dashboard(request):
    context = {
        "current_user": request.user.user_uuid,
        "permissions": request.permissions,
    }
    return render(request, 'pages/dashboard.html', context)


def login(request):
    return render(request, 'pages/login.html')

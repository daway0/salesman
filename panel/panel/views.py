from django.shortcuts import render
from django.http import JsonResponse
import requests

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
    }
    if id:
        context["edit"] = True
        context["user_id"] = id
    return render(request, 'pages/users_create_update.html', context)


def users_list(request):
    context = {
        "current_user": None,
    }
    return render(request, 'pages/users_list.html', context)

# companies
def companies_create_update(request, id=None):
    context = {
        "initiate": True,
        "company_id": None,
        "current_user": None,
    }
    if id:
        context["initiate"] = False
        context["company_id"] = id
    return render(request, 'pages/companies_create_update.html', context)

def companies_list(request):
    context = {
        "current_user": None,
    }
    return render(request, 'pages/companies_list.html', context)

# services
def services_create_update(request, id=None):
    context = {
        "edit": False,
        "service_id": None,
        "current_user": None,
    }
    if id:
        context["edit"] = True
        context["service_id"] = id
    return render(request, 'pages/services_create_update.html', context)

def services_list(request):
    context = {
        "current_user": None,
    }
    return render(request, 'pages/services_list.html', context)

# sales ledgers
def sales_ledger_create_update(request, id=None):
    context = {
        "edit": False,
        "sales_ledger_id": None,
        "current_user": "0a7c4c35-d456-4fde-a4a8-475ef9f8a2cf",
        "user_role": "operator",
    }
    if id:
        context["edit"] = True
        context["sales_ledger_id"] = id
    return render(request, 'pages/sales_ledger_create_update.html', context)    

def sales_ledger_list(request):
    context = {
        "current_user": None,
    }
    return render(request, 'pages/sales_ledger_list.html', context)

# dashboard
def dashboard(request):
    context = {
        "current_user": None,
    }
    return render(request, 'pages/dashboard.html', context)


def login(request):
    context = {
        "current_user": None,
    }
    return render(request, 'pages/login.html', context)

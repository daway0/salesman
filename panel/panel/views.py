from django.shortcuts import render
from django.http import JsonResponse
import requests

def users_create_update(request, id=None):
    if id:
        return render(request, 'pages/users_create_update.html', {
            "edit": True,
            "user_id": id,
        })
    else:
        return render(request, 'pages/users_create_update.html', {
            "edit": False,
        })

def users_list(request):
    return render(request, 'pages/users_list.html')


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


def companies_create_update(request, id=None):
    if id:
        return render(request, 'pages/companies_create_update.html', {
            "edit": True,
            "company_id": id,
        })
    else:
        return render(request, 'pages/companies_create_update.html', {
            "edit": False,
        })

def companies_list(request):
    return render(request, 'pages/companies_list.html')


def services_create_update(request, id=None):
    if id:
        return render(request, 'pages/services_create_update.html', {
            "edit": True,
            "service_id": id,
        })
    else:
        return render(request, 'pages/services_create_update.html', {
            "edit": False,
        })

def services_list(request):
    return render(request, 'pages/services_list.html')


def sales_ledger_create_update(request, id=None):
    if id:
        return render(request, 'pages/sales_ledger_create_update.html', {
            "edit": True,
            "sales_ledger_id": id,
        })
    else:
        return render(request, 'pages/sales_ledger_create_update.html', {
            "edit": False,
        })

def sales_ledger_list(request):
    return render(request, 'pages/sales_ledger_list.html')

def dashboard(request):
    return render(request, 'pages/dashboard.html')

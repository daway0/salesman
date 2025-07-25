from functools import wraps
from django.shortcuts import redirect
from django.conf import settings
from django.contrib.auth import get_user_model
from rest_framework_simplejwt.tokens import UntypedToken
from rest_framework_simplejwt.exceptions import InvalidToken, ExpiredTokenError
from rest_framework_simplejwt.backends import TokenBackend

User = get_user_model()

def jwt_required(view_func):
    @wraps(view_func)
    def _wrapped_view(request, *args, **kwargs):
        access_token = request.COOKIES.get("access")
        token_backend = TokenBackend(
            algorithm='HS256',
            signing_key=settings.SECRET_KEY
        )
        if access_token:
            try:
                validated_token = UntypedToken(access_token)
                payload = token_backend.decode(access_token, verify=True)
                user_uuid = payload.get("user_id")
                user = User.objects.get(user_uuid=user_uuid)
                request.user = user
                request.permissions = payload.get("permissions", [])
            except (InvalidToken, User.DoesNotExist, ExpiredTokenError):
                return redirect('/login/')
        else:
            return redirect('/login/')
        return view_func(request, *args, **kwargs)
    return _wrapped_view

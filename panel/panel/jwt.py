from rest_framework_simplejwt.serializers import TokenObtainPairSerializer
from rest_framework_simplejwt.views import TokenObtainPairView
import requests
from django.contrib.auth import get_user_model

User = get_user_model()

class MyTokenObtainPairSerializer(TokenObtainPairSerializer):

    @classmethod
    def get_token(cls, user):
        token = super().get_token(user)

        token['user_id'] = str(user.user_uuid)

        response = requests.get(f"http://salesman:8080/api/users/{user.user_uuid}/permissions")
        if response.status_code == 200:
            token['permissions'] = response.json().get('permissions', [])
        else:
            token['permissions'] = []

        return token
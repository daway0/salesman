from pyexpat import model
from django.contrib.auth.models import AbstractUser
from django.db import models

class CustomUser(AbstractUser):
    email = models.EmailField(unique=True, null=False, blank=False)
    user_uuid = models.UUIDField()

    USERNAME_FIELD = 'email'  
    REQUIRED_FIELDS = [] 

    def __str__(self):
        return self.email
from django import template

register = template.Library()

@register.filter
def has_permission(permissions, perm_name):
    return perm_name in permissions

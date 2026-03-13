"""
URL configuration for farmec project.

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/6.0/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
from django.contrib.auth import views as auth_views
from django.shortcuts import render

def page_not_found(request, exception=None):
    return render(request, '404.html', status=404)

def server_error(request, exception=None):
    return render(request, 'error.html', status=500)

urlpatterns = [
    path('admin/', admin.site.urls),
    path('login/', auth_views.LoginView.as_view(template_name='pages/login.html', next_page='/'), name='login'),
    path('logout/', auth_views.LogoutView.as_view(next_page='/'), name='logout'),
    path('', include('catalog.urls')),
    path('', include('content.urls')),
    path('', include('legal.urls')),
    path('', include('team.urls')),
    path('', include('support.urls')),
]

# Error handlers
handler404 = page_not_found
handler500 = server_error

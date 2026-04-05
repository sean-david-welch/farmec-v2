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
from django.conf import settings
from django.contrib import admin
from django.urls import path, include
from django.contrib.auth import views as auth_views
from django.contrib.sitemaps.views import sitemap
from django.shortcuts import render
from django.views.generic import RedirectView, TemplateView

from farmec.sitemaps import sitemaps

def page_not_found(request, exception=None):
    return render(request, '404.html', status=404)

def server_error(request, exception=None):
    return render(request, 'error.html', status=500)

urlpatterns = [
    path('supplier', RedirectView.as_view(url='/suppliers/', permanent=True)),
    path('supplier/', RedirectView.as_view(url='/suppliers/', permanent=True)),
    path('about', RedirectView.as_view(url='/team/', permanent=True)),
    path('about/', RedirectView.as_view(url='/team/', permanent=True)),
    path('about/policies', RedirectView.as_view(url='/privacy/', permanent=True)),
    path('about/policies/', RedirectView.as_view(url='/privacy/', permanent=True)),
    path('blogs', RedirectView.as_view(url='/blog/', permanent=True)),
    path('blogs/', RedirectView.as_view(url='/blog/', permanent=True)),
    path('blog/exhibitions', RedirectView.as_view(url='/exhibitions/', permanent=True)),
    path('exhibitions', RedirectView.as_view(url='/exhibitions/', permanent=True)),
    path('spareparts', RedirectView.as_view(url='/spareparts/', permanent=True)),
    path('suppliers', RedirectView.as_view(url='/suppliers/', permanent=True)),

    path('sitemap.xml', sitemap, {'sitemaps': sitemaps}, name='django.contrib.sitemaps.views.sitemap'),
    path('robots.txt', TemplateView.as_view(template_name='robots.txt', content_type='text/plain')),
    path('favicon.ico', RedirectView.as_view(url='/static/favicon.svg', permanent=True)),
    path('admin/', admin.site.urls),
    path('login/', auth_views.LoginView.as_view(template_name='pages/login.html', next_page='/'), name='login'),
    path('logout/', auth_views.LogoutView.as_view(next_page='/'), name='logout'),
    path('', include('catalog.urls')),
    path('', include('content.urls')),
    path('', include('legal.urls')),
    path('', include('team.urls')),
]

if settings.DEBUG:
    import debug_toolbar
    urlpatterns += [
        path('__debug__/', include(debug_toolbar.urls)),
        path('test-404/', lambda r: page_not_found(r)),
        path('test-500/', lambda r: server_error(r)),
    ]

# Error handlers
handler404 = page_not_found
handler500 = server_error

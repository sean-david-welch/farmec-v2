from django.urls import path

from . import views

app_name: str = 'legal'

urlpatterns: list = [
    path('privacy/', views.PrivacyListView.as_view(), name='privacy_list'),
    path('terms/', views.TermsListView.as_view(), name='terms_list'),
]

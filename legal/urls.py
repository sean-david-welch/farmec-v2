from django.urls import path

from . import views

app_name: str = 'legal'

urlpatterns: list = [
    path('privacy/', views.TermsListView.as_view(), name='privacy_list'),
]

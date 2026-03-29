from django.urls import path

from . import views

app_name: str = 'team'

urlpatterns: list = [
    path('team/', views.EmployeeListView.as_view(), name='employee_list'),
]

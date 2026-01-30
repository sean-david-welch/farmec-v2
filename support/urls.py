from django.urls import path

from . import views

app_name: str = 'support'

urlpatterns: list = [
    path('warranty-claims/', views.WarrantyclaimListView.as_view(), name='warrantyclaim_list'),
    path('warranty-claims/<str:pk>/', views.WarrantyclaimDetailView.as_view(), name='warrantyclaim_detail'),
    path('parts-required/', views.PartsrequiredListView.as_view(), name='partsrequired_list'),
    path('machine-registrations/', views.MachineregistrationListView.as_view(), name='machineregistration_list'),
    path('machine-registrations/<str:pk>/', views.MachineregistrationDetailView.as_view(), name='machineregistration_detail'),
]

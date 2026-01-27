from django.urls import path

from . import views

app_name: str = 'catalog'

urlpatterns: list = [
    path('suppliers/', views.SupplierListView.as_view(), name='supplier_list'),
    path('suppliers/<slug:slug>/', views.SupplierDetailView.as_view(), name='supplier_detail'),
    path('machines/', views.MachineListView.as_view(), name='machine_list'),
    path('machines/<slug:slug>/', views.MachineDetailView.as_view(), name='machine_detail'),
]

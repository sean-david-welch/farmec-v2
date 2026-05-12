from django.urls import path

from . import views

app_name: str = 'catalog'

urlpatterns: list = [
    path('suppliers/', views.SupplierListView.as_view(), name='supplier_list'),
    path('suppliers/<slug:slug>/', views.SupplierDetailView.as_view(), name='supplier_detail'),
    path('machines/', views.MachineListView.as_view(), name='machine_list'),
    path('machines/<slug:slug>/', views.MachineDetailView.as_view(), name='machine_detail'),
    path('products/<slug:slug>/', views.ProductDetailView.as_view(), name='product_detail'),
    path('spareparts/', views.SparePartsIndexView.as_view(), name='spareparts'),
    path('spareparts/<slug:slug>/', views.SparePartsListView.as_view(), name='spareparts_detail'),
]

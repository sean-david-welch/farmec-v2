from django.urls import path

from . import views

app_name = 'support'

urlpatterns = [
    path('warranty-claim/', views.WarrantyclaimFormView.as_view(), name='warranty_claim'),
    path('machine-registration/', views.MachineregistrationFormView.as_view(), name='machine_registration'),
]

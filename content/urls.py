from django.urls import path

from . import views

app_name: str = 'content'

urlpatterns: list = [
    path('', views.HomeView.as_view(), name='home'),
    path('blog/', views.BlogListView.as_view(), name='blog_list'),
    path('blog/<str:id>/', views.BlogDetailView.as_view(), name='blog_detail'),
    path('exhibitions/', views.ExhibitionListView.as_view(), name='exhibition_list'),
    path('timeline/', views.TimelineListView.as_view(), name='timeline_list'),
]

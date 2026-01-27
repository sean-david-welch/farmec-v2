from django.urls import path

from . import views

app_name: str = 'content'

urlpatterns: list = [
    path('blog/', views.BlogListView.as_view(), name='blog_list'),
    path('blog/<slug:slug>/', views.BlogDetailView.as_view(), name='blog_detail'),
    path('carousel/', views.CarouselListView.as_view(), name='carousel_list'),
    path('exhibitions/', views.ExhibitionListView.as_view(), name='exhibition_list'),
    path('timeline/', views.TimelineListView.as_view(), name='timeline_list'),
]

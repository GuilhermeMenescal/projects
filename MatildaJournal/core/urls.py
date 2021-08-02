from django.urls import path
from .views import index, publications

from django.conf.urls.static import static
from django.conf import settings

urlpatterns = [
    path('', index),
    path('publications', publications),

    ]

urlpatterns += static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)

from django.urls import path
from counter.views import word_count_view

urlpatterns = [
    # for simplicity we use the base path
    path("", word_count_view, name="word_count"),
]

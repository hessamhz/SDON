from authentication.views import login_user, dashboard
from django.contrib.auth.views import LogoutView
from django.contrib import admin
from django.urls import path

urlpatterns = [
    # for simplicity we use the base path
    path("", login_user, name="login"),
    path("dashboard/", dashboard, name="dashboard"),
    path('logout/', LogoutView.as_view(next_page='login'), name='logout'),

]

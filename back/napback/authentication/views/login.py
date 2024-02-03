from django.contrib.auth import authenticate, login
from django.contrib.auth.forms import AuthenticationForm
from django.shortcuts import redirect, render
from django.urls import reverse
from django.contrib import messages

def login_user(request):
    if request.method == "POST":
        form = AuthenticationForm(request, data=request.POST)
        if form.is_valid():
            username = form.cleaned_data.get("username")
            password = form.cleaned_data.get("password")
            user = authenticate(username=username, password=password)
            if user is not None:
                login(request, user)
                return redirect(reverse('dashboard'))
            else:
                messages.error(request, "Username or password is incorrect.")
        else:
            messages.error(request, "Invalid form submission.")
            # Render the index.html template with the invalid form
            return render(request, "hpanel/index.html", {"form": form})
    else:
        form = AuthenticationForm()
    return render(request, "hpanel/index.html", {"form": form})

from django.contrib.auth import authenticate, login
from django.contrib.auth.forms import AuthenticationForm
from django.shortcuts import HttpResponse, redirect, render


def login_user(request):
    # Checks if form is submitted
    if request.method == "POST":
        #data = {
        #    "username": request.POST.get("username"),
        #    "password": request.POST.get("password"),
        #}
        form = AuthenticationForm(data=request.POST)
        print(form)
        if form.is_valid():
            username = form.cleaned_data.get("username")
            password = form.cleaned_data.get("password")
            user = authenticate(request, username=username, password=password)
            if user is not None:
                login(request, user)
                return HttpResponse("user logged in")
            else:
                return HttpResponse("user doesn't exist")
        else:
            # Render the counter.html template with the invalid form
            return render(request, "hpanel/index.html", {"form": form})
    else:
        # Create a new form instance for GET requests
        form = AuthenticationForm()

    # Render the counter.html template with the form assuming that it is a GET
    # return render(request, "counter.html", {"form": form})
    return render(request, "hpanel/index.html", {"form": form})

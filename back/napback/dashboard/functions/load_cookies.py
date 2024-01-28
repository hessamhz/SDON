import requests
from django.conf import settings


# Load cookies from the 'cookie.curl' file
def load_cookies(session):
    cookie_file = settings.COOKIE_FILE
    with open(cookie_file, 'r') as file:
        for line in file:
            parts = line.strip().split()
            if len(parts) == 7:
                domain, _, path, secure, _, name, value = parts
                cookie = requests.cookies.create_cookie(domain=domain, path=path, secure=secure, name=name, value=value)
                session.cookies.set_cookie(cookie)
        return session
    return None
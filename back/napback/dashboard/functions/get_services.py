import json

import requests
from load_cookies import load_cookies

import os
import django

# Set up Django environment
os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'napback.settings')
django.setup()

from django.conf import settings


def get_service():
    session = requests.Session()
    base_url = settings.BASE_URL
    url_service = (
        base_url
        + "onc/connection?hierarchicalLevel==service&view(connEndPoints.ltp.ne)"
    )
    session = load_cookies(session)
    response = session.get(url_service, verify=False)
    if response.status_code == 200:
        try:
            response_json = response.json()
        except requests.exceptions.JSONDecodeError as e:
            print("Error decoding JSON:", e)
            print("Response Content:", response.text)
    service_arr = []

    for i in range(len(response_json)):
        service_dict = {
            "ConnName": response_json[i]["name"],
            "ConnId": response_json[i]["id"],
            "HierarchicalLevel": response_json[i]["hierarchicalLevel"],
            "ConfigState": response_json[i]["configurationState"],
            "LayerRate": response_json[i]["topmostLayerRate"],
            "NeName1": response_json[i]["connEndPoints"][0]["ltp"]["ne"][
                "name"
            ],
            "NeName2": response_json[i]["connEndPoints"][1]["ltp"]["ne"][
                "name"
            ],
            "Port1": response_json[i]["connEndPoints"][0]["ltp"]["name"],
            "Port2": response_json[i]["connEndPoints"][1]["ltp"]["name"],
            "CreationTime": response_json[i]["log"]["creationTime"],
            "ModificationTime": response_json[i]["log"]["modificationTime"],
        }

        service_arr.append(service_dict)
    print("service_arr: ", service_arr)
    return service_arr

"""
def get_infra(base_url, cookie_file):
    url_infra = (
        base_url
        + "onc/connection?hierarchicalLevel==infrastructure&view(connEndPoints.ltp.ne)"
    )
    load_cookies(session)
    response = session.get(url_infra, verify=False)
    response_json = response.json()

    for i in range(len(response_json)):
        if response_json[i]["topmostLayerRate"] == "OTU_2X":
            infra_dict = {
                "ConnName": response_json[i]["name"],
                "ConnId": response_json[i]["id"],
                "HierarchicalLevel": response_json[i]["hierarchicalLevel"],
                "ConfigState": response_json[i]["configurationState"],
                "LayerRate": response_json[i]["topmostLayerRate"],
                "NeName1": response_json[i]["connEndPoints"][0]["ltp"]["ne"][
                    "name"
                ],
                "NeName2": response_json[i]["connEndPoints"][1]["ltp"]["ne"][
                    "name"
                ],
                "Port1": response_json[i]["connEndPoints"][0]["ltp"]["name"],
                "Port2": response_json[i]["connEndPoints"][1]["ltp"]["name"],
                "CreationTime": response_json[i]["log"]["creationTime"],
                "ModificationTime": response_json[i]["log"][
                    "modificationTime"
                ],
            }

    return infra_dict
"""

"""
infra_data = get_infra(base_url, cookie_file)
print("\nInfrastructure Data:")
print(json.dumps(infra_data, indent=4))
"""

service_data = get_service()
print("Service Data:")
print(json.dumps(service_data, indent=4))

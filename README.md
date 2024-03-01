# **Project**: "Scalable Microservice Architecture for Creation and Monitoring of Optical Service"
### Team Challengers: " Fatih Temiz, Hessam Hashemizadeh, Tugrul Kok, Nisanur Camuzcu"

## **Description of the project**

![image](https://github.com/hessamhz/SDON/assets/61333402/15c7de4d-ff1b-4fd2-882b-d64acc15e129)


The project aims to automatize the creation and monitoring of optical services. 
It hinders the underlying complexities of the controller. It is an intent-based SDN program where the user declares what to do and it solves how to do it.
The project is developed on top of the SMOptics Controller API.



User requests through UI for creating infrastructure, creating services and deleting connections(service/infrastructure) and actively user can visualize current services/infrastructure.
User requests proxied via Nginx and Django handle the request and publish to the related NATS topic. Golang Services subscribes to related topics and handles the request via HTTP request to the controller.
On the GO side current services and infrastructure are retrieving, and adding retrieved data into InfluxDB periodically.
On the Django side, InfluxDB queried and passed into UI for visualization.
The project is Dockerized and deployed via Nginx on the server.


## **Description of the repository**



- **back/_**: This folder contains the backend of the project 
    - **back/requirements/**: Python/Django requirements
    - **back/build/**: Dockerfile for backend
    - **back/napback/**: Django
        - **back/napback/authentication/**
            - **back/napback/authentication/migrations/**
            - **back/napback/authentication/static**
            - **back/napback/authentication/views/**
                - **back/napback/authentication/views/dashboard.py**:
                - **back/napback/authentication/views/login.py**:
                - **back/napback/authentication/views/nat_publisher.py**:
                - **back/napback/authentication/views/nat_subscriber.py**:testing subscribe message that send from Django
                - **back/napback/authentication/views/view_infraservice.py**:
            - **back/napback/authentication/views/apps.py**:
            - **back/napback/authentication/views/urls.py**:
    - **back/napback/core/**: Django core (settings, wsgi, asgi,URLs)
    - **back/napback/dashboard/**:
    - **back/napback/static/**:
    - **back/napback/templates/**:
            - **back/napback/templates/hpanel/user_panel.html**: Dashboard HTML file
            - **back/napback/templates/hpanel/index.html** : Login page HTML file
        - **back/napback/templates/login.html**: Template for login page
    - **back/napback/manage.py**
    - **back/napback/.gitignore**
    - **back/napback/pyproject.toml**
    - **back/napback/setup.cfg**
          
        - 
    - 
    - **_lib/program.py_**: ...
- **napcore/_**: This folder contains the backend of the project
    - **_napcore/intenal/client/**: This folder contains functions for HTTP request(GET,POST,DELETE) to SDN Controller which are used in functions (getting NE ID, posting for creation, deleting connection)
    - **_napcore/internal/functions/_**: This folder contains functions of the project (Create LP, Create Infra, DeleteConn, VisualizeService, VisualizeInfra)
        - **_napcore/internal/functions/create_LP.go** : Implementing Creating Services
        - **_napcore/internal/functions/create_infra.go** : Implementing Creating Infrastructure
        - **_napcore/internal/functions/delete_conn.go** : Implementing Deleting Services / Infrastructure
        - **_napcore/internal/functions/handle_nats.go** : Subscribing NATS messages and Modifying Message formats to be able to use in go functions(creations)
        - **_napcore/internal/functions/run_bash_script.go** : Running bash script(update_cookies.sh) for authentication
        - **_napcore/internal/functions/visualize_service.go** : Visualizing Services
        - **_napcore/internal/functions/visualize_infrastructure.go** : Visualizing Infrastructure
        - **_napcore/internal/functions/write_to_influxdb.go** : Getting visualization data from visualize_service and visualize_infrastructure functions and writing to influxdb periodically(every 10s)
    -**_napcore/go.mod**:
    -**_napcore/go.sum**:
    -**_napcore/main.go**:This is main for golang services
    -**_napcore/update_cookies.sh**: Bash script for updating local cookie.curl file via SSH.
          
 - **deploy/_**: This folder contains docker compose
 - **nginx/_**: This folder configuration file and docker file for nginx web server

## **How to run the project**

Describe how to clone the project and run it. Specify commands, etc.

## **Results**

Describe the main results of your project (max 250 words)
We have developed a microservice-based architecture for the creation and monitoring of optical services involving cutting-edge technologies (InfluxDB, NATS, Docker/Docker Compose, Nginx, Django, Golang)
Project automatizes Creating 10GB Infrastructure, Creating 10GB & 1GB Service, Deleting Connections and Visualizing Services & Infrastructure.
Login Page:
![image](https://github.com/hessamhz/SDON/assets/61333402/d99cee60-0882-4876-b409-94a0b7693711)
Dashboard Page:
![image](https://github.com/hessamhz/SDON/assets/61333402/d9d1d589-d9c6-4cc0-b4a5-1857108883aa)





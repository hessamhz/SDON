# **Project**: "Scalable Microservice Architecture for Creation and Monitoring of Optical Service"
### Team Challengers: " Fatih Temiz, Hessam Hashemizadeh, Tugrul Kok, Nisanur Camuzcu"

## **Description of the project**

![image](https://github.com/hessamhz/SDON/assets/61333402/15c7de4d-ff1b-4fd2-882b-d64acc15e129)


The project aims to create and monitor optical services. User requests through UI for creating infrastructure, creating services, deleting connections and visualizing current infrastructure. 


## **Description of the repository**


- **_main.py_**: 
This file contains ...


- **back/_**: This folder contains the backend of the project 

    - **napback/**: Django Framework 
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
          
 - **back/_**: This folder contains the backend of the project 

    - **_lib/utils.py_**: ...
    - **_lib/program.py_**: ...

## **How to run the project**

Describe how to clone the project and run it. Specify commands, etc.

## **Results**

Describe the main results of your project (max 250 words)


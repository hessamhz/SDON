
# **Project**: "Scalable Microservice Architecture for Creation and Monitoring of Optical Service"

#### Team Challengers: " Fatih Temiz, Hessam Hashemizadeh, Tugrul Kok, Nisanur Camuzcu"

## Table of Contents  
- [Introduction](#introduction)
- [Installation](#installation) 
- [Usage](#usage) 
- [Features](#features) 
- [Results](#results)
- [Project Structure](#project_structure)
  
## **Introduction**
This project is designed to streamline the control and management of infrastructure, services, and lightpaths for SMOptics controller. The application leverages a suite of technologies including Nginx, Django, PostgreSQL, InfluxDB, NATS messaging, and a custom Go service, all containerized using Docker for ease of deployment, developed on top of the SMOptics Controller API.

  
![image](https://github.com/hessamhz/SDON/assets/61333402/4a90cf33-4970-494a-8389-41d28fcc4f7b)

  
It hinders the underlying complexities of the controller. It is an intent-based SDN program where the user declares what to do and it solves how to do it.

## **Installation**
To set up the SMOptics Network Management App, ensure you have Docker and Docker Compose installed on your system. You also need to be provided with environment files which for security reasons will not be provided in the repository. 

Then, follow these steps: 
1. Clone the repository:
  ```
  git clone https://github.com/hessamhz/SDON.git
  cd SDON/deploy
  ```
2. Build the compose file:
```
docker-compose build
```
3. Run the compose: 
```
docker-compose up
```
4. Make sure all of the services are up:
``` 
docker-compose ps
```
Just a quick note that ```.env``` file should be present to complete the build and running the project.
Another thing to mention is that is it required to open the port of the nginx if you are planning to deploy on a server with.
```
sudo ufw allow 80
```
  
## **Usage** 
Once the services are up and running, you can access the Django admin panel to manage users and view the system's status. The application performs the following key functions: 
-  **Infrastructure and Service Management**: Create and delete network resources through the user-friendly dashboard. 
-  **Robust and Asynchronous Communication**: The connection between Django and SMOptics controller is completely Asynchronous, hence the user will not be delayed and we can manage retrying the failures. 
-  **Real-time Monitoring**: Go services updates the InfluxDB time-series database every 10 seconds to reflect the current status of services. 
-  **Automated Session Key Updates**: Ensures continuous and secure communication with the SMOptics controller.

## **Features**  
- User-friendly dashboard for managing optical network resources 
- Real-time service status monitoring 
- Secure message queuing with NATS for asynchronous communication 
- Automated session key management for enhanced security 
- Dockerized services for easy deployment and scalability

## **Results**

  

User requests through UI for creating infrastructure, creating services and deleting connections(service/infrastructure) and actively user can visualize current services/infrastructure.

We have developed a microservice-based architecture for the creation and monitoring of optical services involving cutting-edge technologies (InfluxDB, NATS, Docker/Docker Compose, Nginx, Django, Golang)

Project automatizes Creating 10GB Infrastructure, Creating 10GB & 1GB Service, Deleting Connections and Visualizing Services & Infrastructure.

Login Page:

![image](https://github.com/hessamhz/SDON/assets/61333402/d99cee60-0882-4876-b409-94a0b7693711)

Dashboard Page:

![image](https://github.com/hessamhz/SDON/assets/61333402/d9d1d589-d9c6-4cc0-b4a5-1857108883aa)

## **Project Structure**  
The structure of the project files and structure is [here](https://github.com/hessamhz/SDON/blob/main/repo_desc.md). 

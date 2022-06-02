# channel-list
Develop a service to list and organize channels.

Docker Container
```
# Starting Docker Container
$ docker-compose up -d

# Check Startup
$ docker ps -a

# Get into a Docker container
$ docker exec -it [CONTAINER ID] bash

# Connecting to mysql
$ mysql -h 127.0.0.1 -P 3306 -u root -p

# Create Database
mysql> create database channel_list;

# Go to Database
mysql> use channel_list;

# Add Table
mysql> source docker-entrypoint-initdb.d/create-tables.sql

# Check Table
mysql> show tables;

# Exit mysql
mysql> \q

# Exit From Container
$ exit
```
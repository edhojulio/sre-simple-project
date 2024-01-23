# sre-simple-project

Hi! This is my simple project for creating pipeline to deploy app.

There are 2 simple app in this project:
1. hello-go (using golang)
2. hello-node (using nodejs)

# hello-go
This app will print hello greeting and serving port 7777
1. To build
docker build -t hello-go . 
2. To run
docker run -p 7777:7777 hello-go

# hello-node
This app will print hello greeting and serving port 3000
1. To build
docker build -t hello-node . 
2. To run
docker run -p 3000:3000 hello-node
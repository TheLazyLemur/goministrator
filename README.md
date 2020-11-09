# Goministrator

A discord bot that lets you host temporary meeting rooms in your discord channel
these rooms have controlled access and are destroyed upon the last person leaving or
being kicked from the room

## How to run with Docker

 - Download the source code
 - Copy config.example.json to config.json
 - Replace the placeholder with your discord token
 - run ```docker build -t goministrator .```
 - run ```docker run goministrator```

## How to run with Go environment

 - Download and install golang
 - Install Make
 - Download the source code
 - Copy config.example.json to config.json
 - Replace the placeholder with your discord token
 - run ```make build```
 - run ```make run```

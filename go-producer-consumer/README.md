# Go Producer and Consumer for Kafka

This directory contains Go applications for producing and consuming messages to and from a Kafka topic named `game-events`.

## Prerequisites

- Go 1.16+
- Docker
- Docker Compose

## Setup

### 1. Start the Docker Compose Services

To start the Docker Compose services, run:

`docker-compose up -d`

This will start the Zookeeper, Kafka brokers, and Kafdrop services.

### 2. Update /etc/hosts

Add the following entries to your `/etc/hosts` file to ensure that the Kafka brokers can be resolved correctly:
```
127.0.0.1 kafka
127.0.0.1 kafka2 
127.0.0.1 kafka3
```


## Running the Producer

The producer application generates random messages and sends them to the `game-events` topic.

### Build and Run the Producer

Navigate to the `producer` directory and run:

```sh
cd producer
go run main.go
```

## Running the Consumer

The consumer application reads messages from the `game-events` topic and prints them to the console.

### Build and Run the Consumer

Navigate to the `consumer` directory and run:

```sh
cd consumer
go run main.go
```
# Kafka Docker Setup

This repository contains scripts and configurations to set up a Kafka environment using Docker Compose. It includes scripts faciliating the exploration of creating topics and producing and consuming messages.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### 1. Start the Docker Compose Services

To start the Docker Compose services, run:

`docker-compose up -d`

This will start the Zookeeper, Kafka brokers, and Kafdrop services.

### 2. Create a Kafka Topic

To create a Kafka topic named `game-events`, run:

`./create-topic.sh`

### 3. List Kafka Topics

To list all Kafka topics, run:

`./list-topics.sh`

### 4. Produce Messages to the Topic

To produce messages to the `game-events` topic, run:

`./produce-game-events-messages.sh`

### 5. Consume Messages from the Topic

To consume messages from the `game-events` topic, run:

`./consume-game-events-messages.sh`

## Services

### Zookeeper

- **Image:** `confluentinc/cp-zookeeper:latest`
- **Ports:** `2181:2181`

### Kafka Brokers

- **Image:** `confluentinc/cp-kafka:latest`
- **Ports:** 
  - `kafka`: `9092:9092`
  - `kafka2`: `9093:9093`
  - `kafka3`: `9094:9094`

We have three Kafka brokers to support a replication factor of 3. This means that each topic can have its data replicated across three brokers, providing higher availability and fault tolerance.

### Kafdrop

- **Image:** `obsidiandynamics/kafdrop`
- **Ports:** `9000:9000`

Kafdrop is a web UI for viewing Kafka topics and browsing consumer groups. To use Kafdrop, open your web browser and navigate to `http://localhost:9000`. You will be able to see the list of topics, inspect topic configurations, and view messages.

## Configuration

The Docker Compose configuration is defined in `docker-compose.yml`.

## License

This project is licensed under the MIT License.
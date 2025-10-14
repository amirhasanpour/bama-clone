# Car Sales Management Web API (Inspired by [bama.ir](https://bama.ir))
A scalable, observable, and modular RESTful API for managing a car sales platform. Built with Go using the Gin framework, the architecture follows clean coding practices and is containerized using Docker Compose.

---

## System Design Diagram

<p align="center"><img src='/files/images/system_diagram.png' alt='System Design Diagram' /></p>

---

## Features and Tools

- **[JWT-based](https://github.com/golang-jwt/jwt) authentication and authorization for securing protected routes**
- **Car sales management including creating, updating, deleting, and browsing listings**
- **Clean architecture using [Gin framework](https://github.com/gin-gonic/gin) (router → handler → service → contracts → infra)**
- **Input validation using [validator](https://github.com/go-playground/validator) for secure and strict endpoint validation**
- **Environment and configuration management via [viper](https://github.com/spf13/viper) with support for .env and YAML**
- **[PostgreSQL](https://github.com/postgres/postgres) as the main relational database engine**
- **[PgAdmin](https://github.com/pgadmin-org/pgadmin4) as a visual tool for inspecting and managing the [PostgreSQL](https://github.com/postgres/postgres) database**
- **[GORM](https://github.com/go-gorm/gorm) as the ORM layer for interacting with [PostgreSQL](https://github.com/postgres/postgres) using models and struct-based queries**
- **[Docker Compose](https://github.com/docker/compose) for orchestrating all services like DB, [Redis](https://github.com/redis/redis), [Elasticsearch](https://github.com/elastic/elasticsearch), [Prometheus](https://github.com/prometheus/prometheus), and more**
- **[Redis](https://github.com/redis/redis) caching to reduce database load and enhance performance for hot data**
- **[Prometheus](https://github.com/prometheus/prometheus) for real-time metrics collection and monitoring**
- **[Grafana](https://github.com/grafana/grafana) dashboards for visualizing performance and system metrics**
- **Centralized logging pipeline using [Elasticsearch](https://github.com/elastic/elasticsearch), [Filebeat](https://github.com/elastic/beats), and [Kibana](https://github.com/elastic/kibana)**
- **Structured logging using both [zap](https://github.com/uber-go/zap) and [zerolog](https://github.com/rs/zerolog) for performant, JSON-formatted logs**
- **Auto-generated API documentation with [Swagger UI](https://github.com/swaggo/swag) for easy testing and development**

---

## How to run?

### 1- Start Dependencies on Docker

cd to the docker directory and run this command:

```bash
docker compose -f "docker/docker-compose.yaml" up -d setup elasticsearch kibana filebeat postgres pgadmin redis prometheus node-exporter alertmanager grafana
```

### 2- Install Swagger CLI

on the root level of project run this command:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 3- Install Go Dependencies

on the root level of project run this command:

```bash
go mod download
```

### 4- Run the Application

cd to cmd directory:

```bash
cd src/cmd
```

and then run:

```bash
go run main.go
```

### 5- Visit the Application

#### Web API run on  [http://localhost:10000](http://localhost:10000)

#### Swagger on  [http://localhost:10000/swagger/index.html#/](http://localhost:10000/swagger/index.html#/)

```bash
Token Url: http://localhost:10000/api/v1/users/login-by-username
Username: admin
Password: 12345678
```

#### Kibana on  [http://localhost:5680](http://localhost:5680)

```bash
Username: elastic
Password: changeme
```

#### Prometheus on  [http://localhost:30090](http://localhost:30090)

#### Grafana on  [http://localhost:3000](http://localhost:3000)

```bash
Username: admin
Password: foobar
```

#### PgAdmin on  [http://localhost:8888](http://localhost:8888)

```bash
Username: amirhossinhp10@gmail.com
Password: admin
```

#### Postgres Server info:

```bash
Host: postgres_container
Port: 5432
Username: postgres
Password: admin
```

### 6- Stop all Services

```bash
docker compose -f "docker/docker-compose.yaml" down
```

---

## Project preview

### Swagger

<p align="center"><img src='/files/images/swagger.png' alt='swagger preview' /></p>

### Kibana

<p align="center"><img src='/files/images/kibana.png' alt='kibana preview' /></p>

### Grafana

<p align="center"><img src='/files/images/grafana.png' alt='grafana preview' /></p>

### PgAdmin

<p align="center"><img src='/files/images/pgadmin.png' alt='pgadmin preview' /></p>
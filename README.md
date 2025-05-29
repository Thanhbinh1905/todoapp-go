# TodoApp

Simple Todo REST API built with Go, PostgreSQL, and Docker for learning.

---

## Features

- Create, Read, Update, Delete (CRUD) todos
- PostgreSQL for data persistence
- Clean architecture with Repository, Service, Handler layers
- RESTful API endpoints
- Dockerized backend service for easy deployment

---

## Tech Stack

- Go (Golang)
- PostgreSQL
- Docker & Docker Compose
- Gorilla Mux (HTTP router)

---

## Setup & Run

### Prerequisites

- [Go](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker & Docker Compose](https://docs.docker.com/compose/install/)

### Step 1: Clone repo

```bash
git clone https://github.com/Thanhbinh1905/todoapp-go.git
cd todoapp-go
````

### Step 2: Run docker-compose

```bash
docker-compose up -d --build
```

### Step 3: Use the API

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | `/todos`      | List all todos    |
| POST   | `/todos`      | Create a new todo |
| PUT    | `/todos/{id}` | Update todo by ID |
| DELETE | `/todos/{id}` | Delete todo by ID |

---

## API Examples

Create todo:

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go"}'
```

Get todos:

```bash
curl http://localhost:8080/todos
```

Update todo:

```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go Fast", "completed":true}'
```

Delete todo:

```bash
curl -X DELETE http://localhost:8080/todos/1
```



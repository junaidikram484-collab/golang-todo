<div align="center">

# 📝 Go ToDo Service

**A lightweight, production-ready RESTful ToDo API built with Go**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-Framework-00ADD8?style=for-the-badge)](https://gin-gonic.com/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker)](https://www.docker.com/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Ready-326CE5?style=for-the-badge&logo=kubernetes)](https://kubernetes.io/)
[![Swagger](https://img.shields.io/badge/Swagger-Documented-85EA2D?style=for-the-badge&logo=swagger)](https://swagger.io/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)](LICENSE)

[Features](#-features) • [Tech Stack](#️-tech-stack) • [Quick Start](#-quick-start) • [Docker](#-docker) • [Kubernetes](#️-kubernetes) • [API Docs](#-api-documentation) • [Contributing](#-contributing)

</div>

---

## ✨ Features

| Feature | Description |
|--------|-------------|
| 🔄 **Full CRUD** | Create, Read, Update, and Delete ToDo items |
| 🆔 **UUID Support** | Auto-generated unique IDs for every item |
| ❤️ **Health Check** | Built-in `/health` endpoint for uptime monitoring |
| 🗃️ **In-Memory Repo** | Fast in-memory storage, easily extendable to a database |
| 📖 **Swagger Docs** | Auto-generated interactive API documentation |
| 🐳 **Docker Ready** | Single-command containerization |
| ☸️ **Kubernetes Ready** | Deployment and Service manifests included |

---

## 🛠️ Tech Stack

| Component | Technology |
|-----------|------------|
| **Language** | Go 1.21+ |
| **Framework** | [Gin](https://gin-gonic.com/) |
| **API Docs** | [Swaggo / gin-swagger](https://github.com/swaggo/gin-swagger) |
| **Container** | Docker |
| **Orchestration** | Kubernetes / Minikube |

---

## 🚀 Quick Start

### Prerequisites

- [Go 1.21+](https://golang.org/dl/) installed
- (Optional) [Docker](https://docs.docker.com/get-docker/) for containerization
- (Optional) [kubectl](https://kubernetes.io/docs/tasks/tools/) + [Minikube](https://minikube.sigs.k8s.io/docs/start/) for Kubernetes

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/junaidikram484-collab/golang-todo.git
cd golang-todo
```

### 2️⃣ Run Locally

```bash
go mod tidy
go run main.go
```

> 🌐 API is now available at `http://localhost:8080`

---

## 🐳 Docker

### Build the Image

```bash
docker build -t todo-service:latest .
```

### Run the Container

```bash
docker run -p 8080:8080 todo-service:latest
```

> 🌐 API is available at `http://localhost:8080`

### Cleanup

```bash
# List all containers
docker ps -a

# Stop and remove the container
docker stop <container_id>
docker rm <container_id>

# Remove the image
docker rmi todo-service:latest
```

---

## ☸️ Kubernetes

### Deploy to Cluster

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

### Check Status

```bash
kubectl get pods
kubectl get svc
```

### Access via Minikube

```bash
minikube service todo-service

### If running locally, need to run this command so that service can be accessed
kubectl port-forward svc/todo-service 8080:80
```

---

## 📖 API Documentation

Interactive Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/todos` | Retrieve all ToDo items |
| `POST` | `/api/v1/todos` | Create a new ToDo item |
| `PUT` | `/api/v1/todos/:id` | Update an existing ToDo item |
| `DELETE` | `/api/v1/todos/:id` | Delete a ToDo item |
| `GET` | `/api/v1/health` | Health check |

### Example Request

```bash
# Create a new ToDo
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy groceries", "completed": false}'
```

```bash
# Get all ToDos
curl http://localhost:8080/api/v1/todos
```

---

## 📁 Project Structure

```
golang-todo/
├── main.go               # Application entry point
├── Dockerfile            # Docker container config
├── go.mod / go.sum       # Go module files
├── k8s/
│   ├── deployment.yaml   # Kubernetes Deployment
│   └── service.yaml      # Kubernetes Service
└── docs/                 # Swagger generated docs
```

---

## 🤝 Contributing

Contributions are welcome! Here's how to get started:

1. **Fork** the repository
2. **Create** a feature branch
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Commit** your changes
   ```bash
   git commit -m "feat: add your feature description"
   ```
4. **Push** to your branch
   ```bash
   git push origin feature/your-feature-name
   ```
5. **Open** a Pull Request

Please make sure your code follows Go best practices and includes relevant tests.

---

<div align="center">

Made with ❤️ using Go & Gin

⭐ If you found this helpful, give it a star!

</div>

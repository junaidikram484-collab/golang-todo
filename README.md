📝 Go ToDo Service






A simple RESTful ToDo API in Go, using Gin, documented with Swagger, containerized with Docker, and ready for Kubernetes deployment.

🌟 Features

✅ Create, Read, Update, Delete (CRUD) ToDos

✅ Auto-generated IDs using UUID

✅ Health-check endpoint

✅ In-memory repository (can extend to database)

✅ Swagger API documentation

✅ Dockerized & Kubernetes-ready

🛠 Tech Stack
Component	Version / Tool
Language	Go 1.25
Framework	Gin
API Docs	Swagger (swaggo/gin-swagger)
Container	Docker
Orchestration	Kubernetes / Minikube
🚀 Quick Start
1️⃣ Clone the Repository
git clone https://github.com/junaidikram484-collab/golang-todo.git
cd golang-todo
2️⃣ Run Locally (Without Docker)
go mod tidy
go run main.go

Access API: http://localhost:8080

🐳 Docker
Build Docker Image
docker build -t todo-service:latest .
Run Container
docker run -p 8080:8080 todo-service:latest

API available at: http://localhost:8080

Optional: Clean Up
docker ps -a
docker stop <container_id>
docker rm <container_id>
docker rmi todo-service:latest
☸ Kubernetes
Apply Deployment & Service
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
Check Status
kubectl get pods
kubectl get svc
Access Service (Minikube)
minikube service todo-service
📖 API Documentation (Swagger)

Swagger UI:

http://localhost:8080/swagger/index.html
Method	Endpoint	Description
GET	/api/v1/todos	List all ToDos
POST	/api/v1/todos	Create a ToDo
PUT	/api/v1/todos/:id	Update a ToDo
DELETE	/api/v1/todos/:id	Delete a ToDo
GET	/api/v1/health	Health check
🤝 Contributing

Fork the repo

Create a branch: git checkout -b feature-name

Commit changes: git commit -m "Your message"

Push: git push origin feature-name

Open a Pull Request

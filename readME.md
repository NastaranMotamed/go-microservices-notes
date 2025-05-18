## 🧾 Project Overview: Go Microservices with gRPC, REST, and Docker

This project demonstrates a simple microservices architecture built with Go, utilizing **gRPC** for internal communication between services and **REST** for external API access.

It includes two main services:

- **user-service** – A gRPC server that returns mock user information based on an ID  
- **note-service** – A REST API server that consumes user-service via gRPC and exposes endpoints for external clients (e.g. Postman or a frontend)

---

## 🚀 Getting Started

To build and run the project:

```bash
docker-compose up --build
```

- `user-service` will be available on port **50052** (gRPC)  
- `note-service` will be available on port **50051** (REST)

---

## 🧪 Testing with Postman

Send a `GET` request to the following endpoint:

```
http://localhost:50051/notes/123
```

### ✅ Sample JSON Response:

```json
{
  "id": "123",
  "name": "User 123"
}
```

---

## ⚙️ Technologies Used

- Go (1.20+)  
- gRPC + Protocol Buffers  
- REST API (`net/http`)  
- Docker & Docker Compose

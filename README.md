# Go User API (Standard Library)

A simple REST API built using **Go standard library (`net/http`)** with in-memory storage.  
This project is made for learning backend development in Go step-by-step.

---

## ✅ Features
- Health check endpoint
- Create user (POST)
- List users (GET)
- Get user by ID (GET)
- Update user by ID (PUT)
- Delete user by ID (DELETE)

---

## 📁 Folder Structure

```text
go-user-api/
├── handler/
├── model/
├── store/
├── main.go
├── go.mod
└── README.md
```

---

## 🚀 Run Locally

### 1) Start the server
```bash
go run main.go
```
Server runs on: `http://localhost:8080`

---

## 🧪 API Testing (curl.exe)

### ✅ Health Check
```bash
curl.exe http://localhost:8080/health
```
**Expected response:** `OK`

### ✅ Create User (POST /users)
```bash
curl.exe -X POST http://localhost:8080/users -H "Content-Type: application/json" -d "{\"name\":\"Tejas\",\"email\":\"t@gmail.com\"}"
```
**Expected response:** `{"id":1,"name":"Tejas","email":"t@gmail.com"}`

### ✅ List Users (GET /users)
```bash
curl.exe http://localhost:8080/users
```
**Expected response:** `[{"id":1,"name":"Tejas","email":"t@gmail.com"}]`

### ✅ Get User By ID (GET /users/{id})
```bash
curl.exe http://localhost:8080/users/1
```
**Expected response:** `{"id":1,"name":"Tejas","email":"t@gmail.com"}`

### ✅ Update User (PUT /users/{id})
```bash
curl.exe -X PUT http://localhost:8080/users/1 -H "Content-Type: application/json" -d "{\"name\":\"Updated Name\",\"email\":\"updated@gmail.com\"}"
```
**Expected response:** `{"id":1,"name":"Updated Name","email":"updated@gmail.com"}`

### ✅ Delete User (DELETE /users/{id})
```bash
curl.exe -X DELETE http://localhost:8080/users/1
```
**Expected response:** `User deleted successfully`
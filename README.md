# Expense Tracker API

The **Expense Tracker API** is a RESTful service built with the Gin framework in Go. It provides endpoints to manage user credentials and track expenses, including features like JWT-based authentication, secure password hashing, and CRUD operations on expense records.

## Features

- **Authentication:**
  - User Signup and Login with JWT-based token generation.
- **Expense Management:**
  - Add, update, delete, and retrieve expenses.
  - Filter expenses by categories like Groceries, Electronics, etc.
- **Security:**
  - Passwords are hashed for secure storage.
  - Token validation with middleware for protected routes.

---

## Table of Contents

1. [Requirements](#requirements)
2. [Setup](#setup)
3. [API Endpoints](#api-endpoints)
4. [Environment Variables](#environment-variables)
5. [Run Locally](#run-locally)
6. [Testing with cURL](#testing-with-curl)
7. [License](#license)

---

## Requirements

- **Go** (v1.20+ recommended)
- **PostgreSQL/MySQL/SQLite** (as the database backend)
- **cURL** (for testing API endpoints)
- **Git** (for version control)

---

## Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/username/expense-tracker-api.git
   cd expense-tracker-api
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Set up the database:**
   - Create a database for the project.
   - Update the connection string in the `config.yml` file or as an environment variable.

4. **Run the server:**
   ```bash
   go run main.go
   ```

---

## API Endpoints

### **Auth Endpoints**
| Method | Endpoint        | Description            | Auth Required |
|--------|-----------------|------------------------|---------------|
| POST   | `/signup`       | User Signup           | No            |
| POST   | `/login`        | User Login (JWT Token)| No            |

### **Expense Endpoints**
| Method | Endpoint                | Description                 | Auth Required |
|--------|-------------------------|-----------------------------|---------------|
| POST   | `/expenseAPI/expense`   | Add a new expense           | Yes           |
| GET    | `/expenseAPI/expense`   | Get all expenses            | Yes           |
| GET    | `/expenseAPI/expense/:id` | Get expense by ID          | Yes           |
| PUT    | `/expenseAPI/expense/:id` | Update expense by ID       | Yes           |
| DELETE | `/expenseAPI/expense/:id` | Delete expense by ID       | Yes           |

---

## Environment Variables

| Variable       | Description                       |
|----------------|-----------------------------------|
| `PORT`         | Server listening port (default: 8080) |
| `DATABASE_URL` | Connection string for the database |
| `JWT_SECRET`   | Secret key for signing JWT tokens |

---

## Run Locally

1. **Start the server:**
   ```bash
   go run main.go
   ```

2. **Access the API:**
   Visit `http://localhost:8080` to access the API.

---

## Testing with cURL

### **1. Sign Up**
```bash
curl -X POST http://localhost:8080/signup \
-H "Content-Type: application/json" \
-d '{"username": "testuser", "password": "testpassword"}'
```

### **2. Log In**
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username": "testuser", "password": "testpassword"}'
```

### **3. Add an Expense**
```bash
curl -X POST http://localhost:8080/expenseAPI/expense \
-H "Authorization: Bearer <JWT_TOKEN>" \
-H "Content-Type: application/json" \
-d '{"category": "Groceries", "amount": 50.75, "description": "Lunch", "date": "2024-11-13"}'
```

### **4. Get All Expenses**
```bash
curl -X GET http://localhost:8080/expenseAPI/expense \
-H "Authorization: Bearer <JWT_TOKEN>"
```

---

## License

This project is licensed under the [MIT License](LICENSE).

Feel free to contribute to this project or use it as a reference for your own projects! 🚀

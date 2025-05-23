📘 Finance Tracker – Backend API

Modular RESTful API built in Golang + Gin + PostgreSQL to track income, expenses, transfers, taxes, and personal loans. Includes user authentication with JWT.
📦 Features

    ✅ JWT-based Authentication

    💼 Account Management (Bank, Cash, E-Wallet)

    🗂️ Custom Category System

    💸 Full Transaction Logging (income, expense, transfer, tax, loan)

    🧾 Personal Loan Tracker

    📊 Ready for frontend integration

🚀 Getting Started
📁 Project Structure

finance-tracker/
├── cmd/ # main.go (entry point)
├── config/ # DB config
├── controllers/ # HTTP handlers
├── models/ # Data structures
├── routes/ # Route grouping
├── services/ # Business logic
├── middlewares/ # JWT auth
├── utils/ # Helpers (hashing, response)
├── .env # Env variables
├── go.mod / go.sum # Dependencies

🧱 Prerequisites

    Golang 1.21+

    PostgreSQL

    Git

📦 Install Dependencies

go mod tidy

🧪 Run Server

go run cmd/main.go

🔐 Authentication
Register

POST /register

{
"name": "Raditya",
"email": "raditya@example.com",
"password": "test1234"
}

Login

POST /login

{
"email": "raditya@example.com",
"password": "test1234"
}

    Returns: token, user

🛡️ Protected Routes

    Include header:

Authorization: Bearer <JWT_TOKEN>

💼 Accounts
Method Endpoint Description
GET /accounts List user accounts
POST /accounts Create new account
PUT /accounts/:id Update account
DELETE /accounts/:id Delete account
🗂️ Categories
Method Endpoint Description
GET /categories List categories
POST /categories Create category
PUT /categories/:id Update category
DELETE /categories/:id Delete category
💸 Transactions
Method Endpoint Description
GET /transactions List all transactions
POST /transactions Create new transaction
PUT /transactions/:id Update transaction
DELETE /transactions/:id Delete transaction
🧾 Loans
Method Endpoint Description
GET /loans List all loans
POST /loans Create a new loan
PUT /loans/:id Update loan
DELETE /loans/:id Delete loan
📄 .env Example

PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=12345678
DB_NAME=finance
JWT_SECRET=supersecretkey123

🧪 Testing

Use Postman or curl with Bearer token:

curl -H "Authorization: Bearer <token>" http://localhost:8080/accounts

📌 License

MIT License

Jika kamu ingin file ini sebagai README.md nyata, cukup salin ke root proyek dan simpan. Saya juga bisa bantu buatkan Swagger UI jika diperlukan.

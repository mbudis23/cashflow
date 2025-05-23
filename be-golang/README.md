# 💰 Finance Tracker API (Golang + Gin + PostgreSQL)

Finance Tracker adalah aplikasi backend modular berbasis Golang yang memungkinkan pengguna untuk mencatat pengeluaran, pemasukan, transfer antar akun, pinjaman pribadi, dan pajak.

---

## 🔧 Teknologi

- **Golang** (Gin Framework)
- **PostgreSQL**
- **JWT Authentication**
- **Modular Architecture**
- **.env Configuration**

---

## 🗂️ Struktur Direktori

```markdown
finance-tracker/
├── cmd/ # main.go (entry point)
├── config/ # Database connection
├── controllers/ # Handler untuk semua endpoint
├── models/ # Struct model
├── routes/ # Routing konfigurasi
├── services/ # Business logic
├── middlewares/ # Middleware (JWT auth)
├── utils/ # Helper functions
├── go.mod / go.sum # Dependency manager
├── .env # Konfigurasi environment
└── README.md # Dokumentasi proyek
```

---

## 🚀 Cara Menjalankan

### 1. Clone dan Masuk ke Proyek

```bash
git clone https://github.com/yourname/finance-tracker.git
cd finance-tracker
```

### 2. Setup Environtment

Buat file .env

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=12345678
DB_NAME=finance
JWT_SECRET=supersecretkey123
```

### 3. Setup Database

Jalankan SQL berikut di PostgreSQL:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    balance NUMERIC(14, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    color VARCHAR(20),  -- opsional: HEX atau nama warna
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL CHECK (type IN ('income', 'expense', 'transfer', 'loan')),
    amount NUMERIC(14, 2) NOT NULL CHECK (amount >= 0),
    account_from INTEGER REFERENCES accounts(id),
    account_to INTEGER REFERENCES accounts(id),
    category_id INTEGER REFERENCES categories(id),
    description TEXT,
    is_tax_applied BOOLEAN DEFAULT FALSE,
    tax_amount NUMERIC(14, 2) DEFAULT 0,
    note TEXT,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    contact_name VARCHAR(100) NOT NULL,
    amount NUMERIC(14, 2) NOT NULL CHECK (amount >= 0),
    loan_type VARCHAR(10) CHECK (loan_type IN ('given', 'received')) NOT NULL,
    due_date DATE,
    interest NUMERIC(5, 2) DEFAULT 0,
    is_paid BOOLEAN DEFAULT FALSE,
    note TEXT,
    related_transaction_id INTEGER REFERENCES transactions(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 4. Install Dependency & Run

```bash
go mod tidy
go run cmd/main.go
```

# 🔐 Finance Tracker API

API backend modular untuk mencatat pengeluaran, pemasukan, transfer antar akun, pajak, dan pinjaman pribadi menggunakan Golang, Gin, dan PostgreSQL.

---

## 🔐 Autentikasi

### 📥 Register

**POST** `/register`

**Body**

```json
{
  "name": "Raditya",
  "email": "raditya@example.com",
  "password": "test1234"
}
```

## 🔑 Login

**POST** `/login`
**Body**

```json
{
  "email": "raditya@example.com",
  "password": "test1234"
}
```

**Response**

```json
{
  "token": "JWT_TOKEN",
  "user": {
    "id": 1,
    "name": "Raditya",
    "email": "..."
  }
}
```

Gunakan token sebagai header:

```makerfile
Authorization: Bearer JWT_TOKEN
```

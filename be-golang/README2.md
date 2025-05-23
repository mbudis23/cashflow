# 📊 Finance Tracker API

Finance Tracker adalah RESTful API yang dirancang untuk mengelola akun finansial pribadi, kategori pengeluaran, transaksi, dan pinjaman. Sistem ini menggunakan:

- Go (Golang)
- Gin Framework
- PostgreSQL
- JSON Web Token (JWT) untuk autentikasi

## 🔧 Cara Menjalankan

1. **Clone repositori**
   ```bash
   git clone https://github.com/namamu/finance-tracker.git
   cd finance-tracker
   ```
2. **Buat file `.env`**

```env
JWT_SECRET=your_secret_key
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=finance_tracker
DB_HOST=localhost
DB_PORT=5432

```

3. **Jalankan Aplikasi**

```bash
go run cmd/main.go
```

## 🔐 Autentikasi

API ini menggunakan JWT (Bearer token). Anda harus login terlebih dahulu dan menyertakan token dalam header `Authorization`.

```makefile
Authorization: Bearer <your-token>
```

## 🛠️ Endpoints

### 🧍 Auth

`POST /register`
Mendaftarkan user baru.
**Request Body**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "yourpassword"
}
```

`POST /login`
Login user dan mendapatkan token JWT.
**Request Body**

```json
{
  "email": "john@example.com",
  "password": "yourpassword"
}
```

### 💼 Accounts

`GET /accounts`
Mengambil semua akun user.

`POST /accounts`
Membuat akun baru.

```json
{
  "name": "Bank BCA",
  "type": "cash",
  "balance": 1000000
}
```

`PUT /accounts/:id`
Mengubah informasi akun.

`DELETE /accounts/:id`
Menghapus akun.

### 🏷️ Categories

`GET /categories`
Mengambil semua kategori user.

`POST /categories`
Menambah kategori baru.

```json
{
  "name": "Makan",
  "color": "#FF5733"
}
```

`PUT /categories/:id`
Mengubah kategori.
`DELETE /categories/:id`
Menghapus kategori.

### 💸 Transactions

`GET /transactions`
Mengambil semua transaksi user.

`POST /transactions`
Membuat transaksi baru.

```json
{
  "type": "expense",
  "amount": 50000,
  "account_from": 1,
  "account_to": null,
  "category_id": 2,
  "description": "Makan siang",
  "is_tax_applied": false,
  "tax_amount": 0,
  "note": "dengan teman",
  "date": "2025-05-23T00:00:00Z"
}
```

### 🏦 Loans

`GET /loans`
Mengambil semua data pinjaman.

`POST /loans`
Menambahkan pinjaman.

```json
{
  "contact_name": "Budi",
  "amount": 200000,
  "loan_type": "debt",
  "due_date": "2025-06-15",
  "interest": 5.0,
  "is_paid": false,
  "note": "Pinjam uang beli motor",
  "related_transaction_id": null
}
```

`PUT /loans/:id`
Mengubah informasi pinjaman.

`DELETE /loans/:id`
Menghapus pinjaman.

## 📄 Struktur Folder

```bash
finance-tracker/
├── cmd/
│ └── main.go # Entry point aplikasi
│
├── config/
│ └── database.go # Setup koneksi database
│
├── controllers/
│ └── transaction_controller.go
│ └── account_controller.go
│ └── user_controller.go
│ └── loan_controller.go
│ └── category_controller.go
│
├── models/
│ └── transaction.go
│ └── account.go
│ └── user.go
│ └── loan.go
│ └── category.go
│
├── routes/
│ └── routes.go # Mengatur semua routing
│
├── services/
│ └── transaction_service.go
│ └── loan_service.go
│ └── user_service.go
│ └── account_service.go
│
├── middlewares/
│ └── auth.go # Middleware JWT/authentikasi
│
├── utils/
│ └── response.go # Standard response helper
│ └── hash.go # Password hashing utils
│
├── .env # Environment variables (dsn, port, secret, etc)
├── go.mod
└── go.sum
```

© 2025 Finance Tracke

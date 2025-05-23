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

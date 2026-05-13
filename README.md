# Go Example Project — CRUD + JWT Auth

## Project Structure
```
go-example/
├── main.go               ← entry point, wire dependencies
├── .env.example          ← copy เป็น .env
├── config/
│   └── config.go         ← load env, connect DB
├── model/
│   └── user.go           ← struct + DTO
├── repository/
│   └── user_repository.go ← DB queries (interface + impl)
├── service/
│   └── user_service.go   ← business logic
├── handler/
│   └── user_handler.go   ← HTTP handlers
├── middleware/
│   └── auth.go           ← JWT middleware
└── router/
    └── router.go         ← route definitions
```

## วิธีรัน

```bash
# 1. copy env
cp .env.example .env

# 2. install dependencies
go mod tidy

# 3. run
go run main.go
```

## API Endpoints

| Method | Path               | Auth | Description        |
|--------|--------------------|------|--------------------|
| GET    | /health            | -    | health check       |
| POST   | /api/auth/register | -    | สมัครสมาชิก        |
| POST   | /api/auth/login    | -    | login → JWT token  |
| GET    | /api/users         | -    | ดู user ทั้งหมด    |
| GET    | /api/users/:id     | -    | ดู user ตาม ID    |
| GET    | /api/me            | JWT  | ดูข้อมูลตัวเอง     |
| PUT    | /api/users/:id     | JWT  | แก้ไข user         |
| DELETE | /api/users/:id     | JWT  | ลบ user            |

## ตัวอย่าง Request

```bash
# Register
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Nuthipong","email":"n@mail.com","password":"123456"}'

# Login → copy token ที่ได้
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"n@mail.com","password":"123456"}'

# Get me (ใส่ token)
curl http://localhost:8080/api/me \
  -H "Authorization: Bearer <token>"
```

## Dependencies
- [Fiber](https://gofiber.io/) — HTTP framework
- [GORM](https://gorm.io/) — ORM
- [SQLite driver](https://github.com/gorm-io/sqlite) — local DB ไม่ต้องติดตั้งอะไร
- [jwt](https://github.com/golang-jwt/jwt) — JWT
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — hash password
- [godotenv](https://github.com/joho/godotenv) — .env loader

# ERP Fullstack

ระบบ ERP แบบ Fullstack ประกอบด้วย Go (Fiber) backend และ React (TypeScript) frontend

## Project Structure

```
erp-fullstack/
├── backend/          ← Go API server (Fiber + GORM + SQLite)
│   ├── cmd/server/
│   │   └── main.go
│   ├── internal/
│   │   ├── config/
│   │   ├── handler/
│   │   ├── middleware/
│   │   ├── model/
│   │   ├── repository/
│   │   ├── router/
│   │   └── service/
│   ├── docs/         ← Swagger generated docs
│   ├── go.mod
│   └── .env.example
└── frontend/         ← React + TypeScript (Vite)
    ├── src/
    │   ├── api/
    │   ├── components/
    │   ├── context/
    │   ├── hooks/
    │   ├── pages/
    │   └── types/
    ├── package.json
    └── vite.config.ts
```

## Tech Stack

| Layer    | Technology                          |
|----------|-------------------------------------|
| Backend  | Go, Fiber v2, GORM, SQLite, JWT     |
| Frontend | React 18, TypeScript, Vite, Axios   |
| Docs     | Swagger / swaggo                    |
| Test     | testify (backend), Vitest (frontend)|

## วิธีรัน

### Backend

```bash
cd backend

# 1. copy env
cp .env.example .env

# 2. install dependencies
go mod tidy

# 3. run
go run cmd/server/main.go
```

Server จะรันที่ `http://localhost:8080`
Swagger UI: `http://localhost:8080/swagger/index.html`

### Frontend

```bash
cd frontend

# 1. install dependencies
npm install

# 2. run dev server
npm run dev
```

App จะรันที่ `http://localhost:5173`

## API Endpoints

| Method | Path                      | Auth | Description           |
|--------|---------------------------|------|-----------------------|
| GET    | /health                   | -    | health check          |
| POST   | /api/auth/register        | -    | สมัครสมาชิก           |
| POST   | /api/auth/login           | -    | login → JWT token     |
| GET    | /api/users                | JWT  | ดู user ทั้งหมด       |
| GET    | /api/users/:id            | JWT  | ดู user ตาม ID        |
| GET    | /api/me                   | JWT  | ดูข้อมูลตัวเอง        |
| PUT    | /api/users/:id            | JWT  | แก้ไข user            |
| DELETE | /api/users/:id            | JWT  | ลบ user               |
| GET    | /api/employees            | JWT  | จัดการพนักงาน         |
| GET    | /api/inventory            | JWT  | จัดการสินค้าคงคลัง    |
| GET    | /api/sales                | JWT  | จัดการการขาย          |
| GET    | /api/purchases            | JWT  | จัดการการซื้อ         |

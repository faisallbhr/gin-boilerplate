# 🧩 Golang Gin Boilerplate

Go Gin boilerplate with JWT authentication, environment-based configuration, structured responses, pagination, search, and sorting.

---

## ✨ Features

- ✅ Multiple environment support (`.env` files)
- 🔐 JWT Authentication with Refresh Token
- 📄 Structured JSON response with metadata
- 🔎 Pagination, Search, and Sorting
- 🧪 Scalable modular project structure

---

## 📁 Project Structure

```
.
├── config/            # Application configuration
├── controllers/       # HTTP request handlers
├── database/          # DB initialization & connection
├── helpers/           # Utility functions
├── middlewares/       # HTTP middlewares
├── models/            # GORM models
├── presenters/        # API response formatting
├── requests/          # Request validation
├── routes/            # Route definitions
├── structs/           # Common shared structures
├── .env.example       # Environment variables template
├── .env.development   # Development environment config
├── .env.staging       # Staging environment config
├── .env.production    # Production environment config
├── go.mod             # Go module dependencies
├── go.sum             # Go module checksums
└── main.go            # Application entry point
```

---

## 🚀 Getting Started

### 📥 Installation

```bash
git clone https://github.com/faisallbhr/gin-boilerplate.git
cd gin-boilerplate
go mod download
cp .env.example .env.development
```

Edit `.env.development` as needed.

---

### 🏃 Running the Application

```bash
# Default (.env.development)
go run main.go

# Or use other environments:
APP_ENV=staging go run main.go
APP_ENV=production go run main.go
```

---

## 🔐 API Usage

### Authentication

| Method | Endpoint             | Description          |
| ------ | -------------------- | -------------------- |
| POST   | `/api/auth/register` | Register a new user  |
| POST   | `/api/auth/login`    | Login and get tokens |
| POST   | `/api/auth/refresh`  | Refresh JWT token    |

### Users

| Method | Endpoint                  | Description                         |
| ------ | ------------------------- | ----------------------------------- |
| GET    | `/api/users`              | List users (search, sort, paginate) |
| GET    | `/api/users/me`           | Get current user profile            |
| GET    | `/api/users/:id`          | Get user by ID                      |
| PATCH  | `/api/users/:id`          | Update user profile                 |
| PATCH  | `/api/users/:id/password` | Change user password                |
| DELETE | `/api/users/:id`          | Delete user                         |

---

## 📦 API Response Format

### ✅ Success

```json
{
  "success": true,
  "message": "Success message",
  "data": <object|array|null>,
  "meta": {
    "search": "<query>",
    "sort": {
      "by": "<field>",
      "order": "asc|desc"
    },
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 100,
      "total_pages": 10
    }
  }
}
```

- `data` is flexible: object, array, or `null`.
- `meta` is optional and mainly used in list endpoints.

### ❌ Error

#### General Error

```json
{
  "success": false,
  "message": "Something went wrong"
}
```

#### Validation Error

```json
{
  "success": false,
  "message": "Validation failed",
  "errors": {
    "field": "Field is required"
  }
}
```

- `errors` contains field-specific feedback (optional).

---

## 🔍 Search, Sort & Pagination

You can pass the following query parameters to list endpoints (e.g., `/api/users`):

| Parameter | Description            | Default | Example         |
| --------- | ---------------------- | ------- | --------------- |
| `search`  | Filter data by keyword | `""`    | `john`          |
| `page`    | Page number            | `1`     | `2`             |
| `limit`   | Items per page         | `10`    | `5`             |
| `sort_by` | Field to sort by       | `id`    | `name`, `email` |
| `order`   | Sort direction         | `desc`  | `asc`, `desc`   |

---

## 🧩 Customization

If you're cloning this for a new project, update your `go.mod`:

```go
module github.com/your-username/your-project
```

Then replace import paths:

```bash
# macOS
grep -rl 'github.com/faisallbhr/gin-boilerplate' . | xargs sed -i '' 's|github.com/faisallbhr/gin-boilerplate|github.com/your-username/your-project|g'

# Linux
grep -rl 'github.com/faisallbhr/gin-boilerplate' . | xargs sed -i 's|github.com/faisallbhr/gin-boilerplate|github.com/your-username/your-project|g'
```

Lastly, run:

```bash
go mod tidy
```

---

## 🤝 Contributing

1. Fork this repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to your branch (`git push origin feature/your-feature`)
5. Create a Pull Request

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

**Happy Coding!** 🚀

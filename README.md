# REST API Application for Item Management

This is a RESTful API built with Go and MySQL for managing item data.  
The application provides basic **CRUD operations** along with **JWT-based authentication**.

---

## ✨ Features

- Get all item data
- Get item data by ID
- Create a new item
- Update an existing item
- Delete an item by ID
- JWT-based user authentication

---

## 📦 Technologies & Libraries Used

| Library      | Purpose                                                   |
|--------------|-----------------------------------------------------------|
| [Fiber](https://github.com/gofiber/fiber)       | Web framework for building REST APIs       |
| [uuid](https://github.com/google/uuid)         | Generates unique IDs for items             |
| [validator](https://github.com/go-playground/validator)   | Request validation                         |
| [GORM](https://gorm.io/)           | ORM (Object Relational Mapping) for MySQL  |
| [GoDotEnv](https://github.com/joho/godotenv)     | Loads environment variables from `.env`    |
| [JWT](https://github.com/golang-jwt/jwt)          | JWT authentication                         |
| [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)       | Password hashing                           |
| [apitest](https://github.com/steinfletcher/apitest)       | API testing                                |
| [Faker](https://github.com/bxcodec/faker)         | Generate fake data for testing             |

---

## 🚀 Running the Application

Make sure you have **Go 1.23.0** or later installed and `MySQL` is running.

To start the application, run:
```bash
make run
make test


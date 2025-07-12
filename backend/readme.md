# 📚 Classroom Management Website

> 🚧 **Status**: In Progress
> 👤 **Role**: Full-stack Developer (Individual Project)

## 📖 Description

A full -feature web application to manage classrooms, built with modern auxiliary technologies by using ** Golang ** and ** Gozero Framework **. The system allows users (teachers) to create, update and manage:
* Class schedule
* Exercise
* Student information
* Score of the classroom
Allow users (students) to monitor points, monitor notifications, monitor lessons

It is designed with secure **JWT-based authentication/authorization**, and uses **MySQL** for data persistence.

## 🌐 Demo & Source

GitHub: [NguyenDuyHungIT17/classroom](https://github.com/NguyenDuyHungIT17/classroom)

---

## 🛠️ Technologies Used

| Layer        | Technology                     |
| ------------ | -----------------------------  |
| Backend      | Golang + GoZero Framework      |
| Auth         | JWT (JSON Web Token)           |
| Database     | MySQL                          |
| Version Ctrl | Git + GitHub                   |

---

## 📁 Project Structure

```
CLASSROOM/
├── backend/
│   ├── api/                   # API route definitions (.api files)
│   ├── common/                # Shared constants, types
│   ├── config/                # Config helpers
│   ├── etc/                   # YAML config files
│   ├── schema/                # SQL schema and explanation
│   │   ├── classroom_gen.sql  # SQL file to generate DB
│   │   └── explain_Sql.txt    # Explanation of schema
│   └── service/
│       └── classroom/
│           ├── api/           # API routing handlers
│           ├── etc/           # Local YAML config
│           ├── internal/
│           │   ├── config/    # Service config structure
│           │   ├── handler/   # HTTP handler functions
│           │   ├── logic/     # Business logic layer
│           │   ├── svc/       # Service context
│           │   └── types/     # Request/Response types
│           ├── model/         # DB models and queries
│           └── utils/         # Utilities (e.g., JWT, password)
├── templates/                 # (Optional) HTML templates or frontend placeholders
├── docker-compose.yml         # Docker configuration for services
├── go.mod                     # Go module file
├── go.sum                     # Go dependencies checksum
├── main.go                    # App entry point
├── test.txt                   # Temporary or test file
└── readme.md                  # Project documentation
```

---

## ⚙️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/NguyenDuyHungIT17/classroom.git
cd backend
```

### 2. Configure MySQL

Make sure you have MySQL running, and create a database:

```sql
CREATE DATABASE classroom CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

Update DataSource at`backend/etc/classroom.yaml` with your DB credentials.

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the server

```bash
go run main.go
```

The API server should start at `http://localhost:8888`.

---


## 🤝 Contribution

This is a personal learning project. Feel free to fork and suggest improvements.

---

## 📬 Contact

For any inquiries or feedback, please reach out to [nguyenduyhungit17@gmail.com](mailto:nguyenduyhungit17@gmail.com).

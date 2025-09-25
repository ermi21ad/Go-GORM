# GO-GORM

Welcome to the **GO-GORM** repository!  
This project is a step-by-step guide to learning **GORM**, a popular ORM (Object-Relational Mapping) library for Go, with a focus on **PostgreSQL integration**.  

It covers essential concepts from setup to advanced features, culminating in a **mini Blog App project**.  
Whether you're a beginner or refining your skills, this repo provides **hands-on examples and a practical application**.

---

## 🚀 Overview

This repository documents a learning journey through **13 lessons** and a **mini project**, exploring GORM's capabilities for backend development.  
Each lesson includes **code examples, explanations, and exercises**, with the final project tying everything together into a functional **blog backend with users and posts**.

- **Language**: Go (Golang)  
- **Database**: PostgreSQL  
- **ORM**: GORM  
- **Project Structure**: Organized with models, migrations, and scopes for modularity  

---

## 📑 Table of Contents

1. [Setup](#-setup)  
2. [Lessons](#-lessons)  
   - Lesson 1: Overview & Purpose of GORM  
   - Lesson 2: Install & Connect to PostgreSQL  
   - Lesson 3: Models & Tags  
   - Lesson 4: AutoMigrate vs Migrations  
   - Lesson 5: CRUD Examples  
   - Lesson 6: Querying (Where, Select, Order, Limit, Offset)  
   - Lesson 7: Scopes & Reusable Queries  
   - Lesson 8: Associations (One-to-Many, Many-to-Many, Preload)  
   - Lesson 9: Transactions with db.Transaction  
   - Lesson 10: Hooks & Password Hashing  
   - Lesson 11: Error Handling & SQL Logging  
   - Lesson 12: Postgres-Specific Tips (UUID, JSONB, Indexes)  
   - Lesson 13: Mini Project (Blog App)  
3. [Contributing](#-contributing)  
4. [License](#-license)  

---

## 🛠 Setup

### Prerequisites

- **Go**: v1.24.5 or later → [Download](https://golang.org)  
- **PostgreSQL**: Installed & running → [Download](https://www.postgresql.org/)  
- **Git**: For cloning the repo  
- **VS Code**: Recommended, with the Go extension  

### Installation

Clone the repository:
git clone https://github.com/ermi21ad/GO-GORM.git

Initialize the Go module:
cd GO-GORM
go mod init GO-GORM

Install dependencies:

go get gorm.io/gorm
go get gorm.io/driver/postgres
go get gorm.io/gorm/logger
go mod tidy

Configure PostgreSQL

Create a database (e.g., mydb) with a user (e.g., postgres) and password (e.g., 1289).

Update the DSN in main.go with your credentials:

dsn := "host=localhost user=postgres password=1289 dbname=mydb port=5432 sslmode=disable"

▶️ Running the Project

Navigate to the desired lesson directory or the blog-app folder, then run:

go run main.go


Check your database output using pgAdmin or CLI.

📚 Lessons
Lesson 1: Overview & Purpose of GORM

Purpose: Intro to ORM concepts & GORM basics

Code: main.go

Lesson 2: Install & Connect to PostgreSQL

Purpose: Setup DSN & gorm.Open

Code: main.go

Lesson 3: Models & Tags

Purpose: Define structs with GORM tags

Code: Model/model.go

Lesson 4: AutoMigrate vs Migrations

Purpose: Compare automatic schema vs manual migrations

Code: main.go

Lesson 5: CRUD Examples

Purpose: Basic Create, Read, Update, Delete

Code: main.go

Lesson 6: Querying

Purpose: Filters, sorting, pagination

Code: main.go

Lesson 7: Scopes

Purpose: Reusable queries for DRY code

Code: scopes/scopes.go

Lesson 8: Associations

Purpose: One-to-Many, Many-to-Many, Preload

Code: main.go

Lesson 9: Transactions

Purpose: Use db.Transaction for atomic ops

Code: main.go

Lesson 10: Hooks & Password Hashing

Purpose: Use BeforeCreate for preprocessing (e.g., hashing)

Code: Model/model.go

Lesson 11: Error Handling & SQL Logging

Purpose: Add robust logging & error handling

Code: main.go

Lesson 12: Postgres-Specific Tips

Purpose: Leverage UUID, JSONB, Indexes

Code: Model/model.go

Lesson 13: Mini Project (Blog App)

Purpose: Build a functional blog backend

Directory: blog-app/

Features:

User & Post models with associations

AutoMigrate schema setup

CRUD for posts

Scopes for querying posts by user

Transaction handling for creating posts

📂 Directory Structure
blog-app/
├── go.mod
├── main.go
├── Model/
│   └── model.go
├── migrations/
│   ├── create_users_table.go
│   └── create_posts_table.go
└── scopes/
    └── scopes.go


Root: go.mod, main.go (entry point)

Model/: Contains User & Post structs

migrations/: Manual migration files

scopes/: Reusable queries

🤝 Contributing

Contributions are welcome!

Fork the repo

Make your changes

Open a Pull Request with a clear description

You can also open issues for bugs, ideas, or feedback.


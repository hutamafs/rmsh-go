# 📝 Personal Blog API (Go + Gin + PostgreSQL)

This project is a backend-only REST API for managing a personal blog. It supports full CRUD operations for blog posts using Go, Gin, and PostgreSQL — without using any ORM. Perfect as a foundational backend project.

---

## 🚀 Features

- Get all blog posts
- Get a single post by ID
- Create a new post
- Update an existing post
- Delete a post
- Connected to PostgreSQL with raw SQL queries

---

## 📁 Project Structure

personal-blog/
├── main.go # Entry point
├── db.go # DB connection logic
├── /handlers # Route handlers
│ └── posts.go # CRUD logic for posts
├── go.mod # Module file

---

## 🔧 Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Database:** PostgreSQL
- **Database Driver:** `lib/pq`
- **Query Type:** Raw SQL

---

## 🛠 Setup Instructions

### 1. Clone the repo

git clone https://github.com/hutamafs/personal-blog.git
cd personal-blog

### 2. Install dependencies

go mod tidy

### 3. Set up PostgreSQL

CREATE TABLE posts (
id SERIAL PRIMARY KEY,
title TEXT NOT NULL,
content TEXT NOT NULL,
published_at TIMESTAMP DEFAULT NOW()
);

### 4. Run the server

go run main.go

Server will run on:
http://localhost:8080

### Project URL: https://github.com/hutamafs/rmsh-go

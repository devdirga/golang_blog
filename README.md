## Details on the structure of the service

### Authentication (`auth`)
Handles user registration, login, and session management. Utilizes JWT for secure authentication.

**Endpoints:**
- `POST /signup`: User registration.
- `POST /signin`: User login.
- `POST /google`: User login using google.

### User Profile (`profile`)
Manages user profiles, including viewing and editing personal information.

**Endpoints:**
- `GET /me`: Retrieve user profile.
- `PATCH /updateprofile`: Update user profile.

### Blog Post
Implements create, gets, get, update, delete blog post feature, allowing users to get profile.

**Endpoints:**
- `GET /post`: Load all data post.
- `GET /post/:id`: Load one data post.
- `PUT /post/:id`: Update one data post.
- `DELETE /post/:id`: Delete one data post.


## Instructions on how to run the service

### Prerequisites
Ensure you have the following installed:
- **Go** (1.2 or later)
- **PostgreSQL** (v13 or later)
- **Redis**
- **Git**
- **Docker** (optional, for redis)

## Installation

### Clone the Repository
```bash
git clone https://github.com/devdirga/golang_blog.git
cd golang_blog
```
---

### Environment Configuration
Create a `.config.json` file in the project root and populate it with the following variables:
```bash
{
  "DB": "host=localhost port=5432 user=postgres password=mysecretpassword dbname=tinder sslmode=disable",
  "Redis": "localhost:6379",
  "IsDebug": true,
  "IsConcurrent":true,
  "Secret": "secret",
  "GoogleSmtpKey": "azdg rkiv wnqe vuil ",
  "URLFront": "http://localhost:3000/"
}
```
Or if you want to quick start, i provide url DB postgres and Redis 
```bash
{
  "DB": "postgres://default:Wwo2qJFu4gkt@ep-solitary-fire-a4wfh8bx.us-east-1.aws.neon.tech:5432/verceldb?sslmode=require",
  "Redis": "redis-10534.c278.us-east-1-4.ec2.redns.redis-cloud.com:10534",
  "RedisPassword":"55t66s1x67yWQ2Kgz8Erb51fJAmCXKzZ",
  "IsDebug": true,
  "IsConcurrent":true,
  "Secret": "secret",
  "GoogleSmtpKey": "azdg rkiv wnqe vuil ",
  "URLFront": "http://localhost:3000/"
}
```

### Backend Setup
Install Go dependencies:
```bash
go mod tidy
```

Run the backend service:
```bash
go run cmd/main.go
```
The service will run on `http://localhost:5000` by default.



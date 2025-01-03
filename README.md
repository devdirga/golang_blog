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
Implements create, gets, get, update, delete blog post feature.

**Endpoints:**
- `GET /post`: Load all data post.
- `GET /post/:id`: Load one data post.
- `PUT /post/:id`: Update one data post.
- `DELETE /post/:id`: Delete one data post.

### Blog Comment
Implements create, gets, get, update, delete blog comment feature.

**Endpoints:**
- `GET /commentall`: Load all data comment.
- `GET /comment/:id`: Load all data comment by post id.
- `PUT /comment/:id`: Update one data comment.
- `DELETE /comment/:id`: Delete one data comment (deactivated).


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

### Alasan penggunaan pattern

**-Modular**, Dengan memisahkan setiap lapisan aplikasi (handler, service, repository, dll.), kita dapat menambahkan fitur baru atau mengubah logika tanpa mengganggu bagian lain dari aplikasi

**-Clean Architecture**, Layer-layer seperti /service/ untuk logika bisnis dan /repository/ untuk akses database mencerminkan prinsip Clean Architecture atau Hexagonal Architecture, yang membuat kode lebih bersih dan mudah diuji

**-Testability**, Layer service dapat diuji dengan mengganti dependency repository menggunakan mock.

**-Penerapan Prinsip SOLID**

**-Kemudahan Maintenance dan Kolaborasi Tim**

Pemahaman Kode yang Mudah,Dengan struktur folder yang jelas, anggota tim baru dapat dengan cepat memahami tanggung jawab setiap bagian aplikasi.

Perubahan yang Aman, Karena kode terpisah dengan baik, perubahan di satu bagian kecil (misalnya, di repository) tidak memengaruhi bagian lain.

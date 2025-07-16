# Go CRUD API (User, Post, Comment)

Bu loyiha Go dasturlash tilida yozilgan RESTful CRUD API bo‘lib, unda:
- `user`, `post`, `comment` entitylari
- `JWT token` asosida autentifikatsiya
- `PostgreSQL` ma'lumotlar bazasi
- `Docker + docker-compose`
- `Chi` routing framework ishlatilgan.

---

##  Loyihaning tuzilishi

```bash
project/
├── cmd/                      # main.go – ilovani ishga tushuradi
│   └── main.go
├── config/                   # konfiguratsiyalar
│   └── config.go
├── internal/
│   ├── handler/              # HTTP handlerlar (Auth, User, Post, Comment)
│   ├── service/              # Biznes logika (Service layer)
│   ├── store/                # Ma'lumotlar bazasi (Repository layer)
│   │   └── postgres/
│   └── model/                # Struct/model fayllari
├── migrations/               # SQL fayllar (up/down)
├── pkg/
│   ├── jwt/                  # JWT yaratish va tekshirish
│   └── hash/                 # Parolni hash qilish
├── Dockerfile
├── docker-compose.yml
├── go.mod                    # Go modul fayli
└── README.md
```

---

##  Texnologiyalar

- **Go 1.22+**
- **PostgreSQL**
- **Chi** – marshrutlash (routing)
- **Docker** va **Docker Compose**
- **JWT (JSON Web Token)** – foydalanuvchi autentifikatsiyasi

---

##  O‘rnatish va ishga tushirish

### 1. `go.mod` va `go.sum` nima uchun kerak?
- `go.mod` — loyiha moduli, bog‘liqliklar, versiyalarni belgilaydi
- `go.sum` — aniqlik uchun bog‘liqliklar hash qiymatlarini saqlaydi

 Yaratish uchun:
```bash
go mod init github.com/yourusername/go-crud-api
```

 Paketlarni tozalab yuklash uchun:
```bash
go mod tidy
```

### 2. Docker orqali ishga tushirish:
```bash
docker-compose build    # konteynerlarni build qiladi
docker-compose up       # loyihani ishga tushiradi
```

---

##  Asosiy buyruqlar va nima uchun kerak:

| Buyruq                           | Nima uchun ishlatiladi                                      |
|----------------------------------|---------------------------------------------------------------|
| `go mod init`                   | Go loyihasini boshlaydi va `go.mod` faylini yaratadi         |
| `go mod tidy`                   | Kerakli paketlarni yuklaydi va `go.mod/go.sum` ni tozalaydi  |
| `docker-compose build`         | Docker konteynerlarini quradi (Go app + Postgres)            |
| `docker-compose up`            | Loyihani ishga tushiradi                                     |
| `go build ./cmd/main.go`       | Go serverni lokal build qilish                               |

---

| Buyruq                     | Nima uchun ishlatiladi                                      |
| -------------------------- | ----------------------------------------------------------- |
| `go mod init <modul_nomi>` | Go loyihasini boshlaydi va `go.mod` faylini yaratadi        |
| `go mod tidy`              | Kerakli paketlarni yuklaydi va `go.mod/go.sum` ni tozalaydi |
| `docker-compose build`     | Docker konteynerlarini quradi (Go app + Postgres)           |
| `docker-compose up`        | Loyihani ishga tushiradi                                    |
| `go build ./cmd/main.go`   | Go serverni lokal build qilish                              |


---

O‘rnatish va ishga tushirish

Loyihani klonlash:

git clone https://github.com/yourusername/go-crud-api.git
cd go-crud-api

Modullarni yuklash:

go mod tidy

Docker orqali ishga tushurish:

docker-compose up --build  -->    Bu: Go serverni build qiladi va PostgreSQL konteynerini ishga tushiradi

---
Docker bilan qayta build qilish:

docker-compose down --> containerni o'chiradi
docker-compose up --build --> build qiladi yani container yaratilinadi


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


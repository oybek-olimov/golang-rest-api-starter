
##  Loyihaning tuzilishi

```bash
project/
├── cmd/                      # main.go – ilovani ishga tushuradi
│   └── main.go
├── config/                   # konfiguratsiyalar
│   └── config.go
├── internal/
│   ├── handler/              # HTTP handlerlar
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

##  Technologies

- **Go 1.24.4**
- **PostgreSQL**
- **Chi Framework**
- **Docker** & **Docker Compose**
- **JWT (JSON Web Token)**

---


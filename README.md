# 📊 Executive Information System (EIS) - Undiksha (Backend API)

Proyek ini merupakan **RESTful API** untuk mendukung website **Executive Information System (EIS) Undiksha**, dibangun menggunakan **Golang** dengan framework **Gin**.
API ini mengambil data dari **MongoDB** (hasil migrasi & ETL dari SQL) serta memanfaatkan **Redis** untuk caching.
Sistem juga telah diintegrasikan dengan **SSO Undiksha** menggunakan **JWT Authentication**.

---

## 🚀 Tech Stack

- **Backend:** Golang + Gin
- **Database:** MongoDB
- **Cache:** Redis
- **ETL Process:** Python (Extract, Transform, Load)
- **Auth:** SSO + JWT
- **Deployment:** Docker & Docker Compose

---

## 🛠️ Fitur Utama

- 🔑 **Autentikasi SSO + JWT** untuk keamanan akses
- ⚡ **Caching dengan Redis** untuk mempercepat akses data
- 📡 **RESTful API** untuk konsumsi data di frontend lama maupun hasil redesign
- 📊 **Dukungan analitik & informasi eksekutif** untuk EIS Undiksha

---

## 📐 Arsitektur Sistem

Berikut gambaran arsitektur sistem yang sedang dibangun:

![Arsitektur Sistem](assets/architecture.png)

---

## 📂 Struktur Project

```bash
.
├── assets/
├── config/
├── constants/
├── docs/
├── handlers/
├── interfaces/
├── middlewares/
├── models/
├── repositories/
├── routes/
├── services/
├── tests/
  ├── mocks/
├── utils/
├── main.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

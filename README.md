
# Noreast CMS

Noreast CMS is a lightweight customer and event management system built with Go and MySQL. It is designed to manage clients, events, and associated business processes for Noreast Entertainment, a Maine-based DJ and event company.

---

## 📁 Project Structure

```
noreast_cms/
├── cmd/
│   └── client/           # HTTP client to connect to server
│   └── server/           # Go HTTP server
├── internal/
│    └── api/
│    └── auth/
│    └── config/
│    └── customer/
│    └── event/
│    └── referral/
│    └── storage/
├── data/
│   └── schema.sql        # MySQL schema for clients and events
├── ui/
│   └── assets/
│   └── components
│   └── windows                    # Placeholder for future web interface
├── test/
├── .env                 # Placeholder for unit/integration tests
├── go.mod
├── go.sum
├── README.md
```

---

## ✅ Current Features

- [x] Simple Go-based HTTP server and client setup
- [x] .env support for configurable server IPs
- [x] MySQL database with:
  - Clients table
  - Events table with foreign key to clients
- [x] Timestamp tracking for clients and events
- [x] Boolean and decimal support for payment tracking

---

## 📝 Current TODO List

- [ ] Connect Go backend to MySQL (via `database/sql` and `go-sql-driver/mysql`)
- [ ] Implement CRUD operations:
  - [ ] Create new client
  - [ ] Edit existing client
  - [ ] Delete client
  - [ ] View client list/details
  - [ ] Create new event (linked to client)
  - [ ] Edit existing event
  - [ ] Delete event
  - [ ] View event list/details
- [ ] API endpoints for all database operations
- [ ] Validate inputs before writing to DB

---

## 🚀 Planned Features

- [ ] Web UI using Go templates or React/Vue
- [ ] Authentication system (admin login)
- [ ] Event filtering (by date, type, etc.)
- [ ] Payment reminders and logs
- [ ] Export client/event data to CSV or PDF
- [ ] Docker support for simplified deployment

---

## 🛠 Technologies

- **Go** - Backend server/client logic
- **MySQL** - Data persistence
- **dotenv** - Environment variable management (`github.com/joho/godotenv`)
- **Future**: Frontend web interface

---

## 📌 Notes

- MySQL schema is stored in `/data/schema.sql`
- You must manually `SOURCE` this file in MySQL to update the schema
- The Go client currently assumes server is available at `http://<LAN_IP>:8080`

---

© 2025 Noreast Entertainment. All rights reserved.

# noreast_cms

Noreast CMS is a simple customer and event management system built with Go and MySQL.  
It provides a web-based interface for managing clients and events for your business.

## Features

- Client management: add, view, and search clients  
- Event management: add, view, and search events  
- MySQL database for persistent data storage  
- Simple and clean web interface with HTML templates  
- Easily extendable and customizable  

## Tech Stack

- [Go](https://golang.org/)  
- [MySQL](https://www.mysql.com/)  
- HTML, CSS (served as static files)  

## Getting Started

### Prerequisites

- Go 1.18+ installed  
- MySQL server running  
- Git (optional, for cloning repo)  

### Installation

1. Clone this repository:

```bash
git clone https://github.com/yourusername/noreast_cms.git
cd noreast_cms
```

2. Create a MySQL database and import the schema:

```bash
mysql -u your_user -p your_database < data/schema.sql
```

3. Update the database connection string in `main.go`:

```go
dsn := "username:password@tcp(localhost:3306)/noreast_cms?parseTime=true"
```

4. Build and run the server:

```bash
go build -o noreast_cms
./noreast_cms
```

5. Open your browser and navigate to:

```
http://localhost:8080
```

## Project Structure

```
/data
  └── schema.sql        # MySQL database schema
/templates
  ├── index.html        # Homepage
  ├── clients.html      # Clients main page
  ├── add_client.html   # Add client form
  ├── view_clients.html # View clients table
  ├── search_clients.html # Search clients form
  ├── events.html       # Events main page
  ├── add_event.html    # Add event form (to be implemented)
  ├── view_events.html  # View events table
  └── search_events.html # Search events form
/static
  └── styles.css        # CSS stylesheet
main.go                 # Go web server
README.md               # This file
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests for bug fixes or new features.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Made with ❤️ by Noreast Entertainment

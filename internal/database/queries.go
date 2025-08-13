package database

import (
	"database/sql"
)

// Client CRUD operations

func GetAllClients(db *sql.DB) ([]Client, error) {
	query := `SELECT id, firstname, lastname, company_name, email, phone, 
              st_address, city, state, zip, date_added FROM clients ORDER BY date_added DESC`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.FirstName, &client.LastName,
			&client.CompanyName, &client.Email, &client.Phone,
			&client.Address, &client.City, &client.State, &client.Zip,
			&client.DateAdded)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func GetClientByID(db *sql.DB, id int) (*Client, error) {
	query := `SELECT id, firstname, lastname, company_name, email, phone, 
              st_address, city, state, zip, date_added FROM clients WHERE id = ?`
	
	var client Client
	err := db.QueryRow(query, id).Scan(&client.ID, &client.FirstName, &client.LastName,
		&client.CompanyName, &client.Email, &client.Phone,
		&client.Address, &client.City, &client.State, &client.Zip,
		&client.DateAdded)
	
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func CreateClient(db *sql.DB, client Client) (int64, error) {
	query := `INSERT INTO clients (firstname, lastname, company_name, email, phone, 
              st_address, city, state, zip) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	result, err := db.Exec(query, client.FirstName, client.LastName,
		client.CompanyName, client.Email, client.Phone,
		client.Address, client.City, client.State, client.Zip)
	
	if err != nil {
		return 0, err
	}
	
	return result.LastInsertId()
}

// Event CRUD operations

func GetAllEvents(db *sql.DB) ([]Event, error) {
	query := `SELECT e.id, e.event_date, e.event_name, e.event_type, e.start_time, 
              e.end_time, e.client_id, e.event_location, e.package, e.guest_count,
              e.deposit_amount, e.deposit_received, e.total_price, e.payment_received,
              e.payment_date, e.notes, c.firstname, c.lastname
              FROM events e 
              JOIN clients c ON e.client_id = c.id 
              ORDER BY e.event_date DESC`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		var clientFirstName, clientLastName string
		var paymentDate sql.NullTime
		
		err := rows.Scan(&event.ID, &event.EventDate, &event.EventName,
			&event.EventType, &event.StartTime, &event.EndTime,
			&event.ClientID, &event.Location, &event.PackageType,
			&event.GuestCount, &event.DepositAmount, &event.DepositReceived,
			&event.TotalPrice, &event.PaymentReceived, &paymentDate,
			&event.Notes, &clientFirstName, &clientLastName)
		
		if err != nil {
			return nil, err
		}
		
		if paymentDate.Valid {
			event.PaymentDate = &paymentDate.Time
		}
		
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(db *sql.DB, id int) (*Event, error) {
	query := `SELECT id, event_date, event_name, event_type, start_time, end_time,
              client_id, event_location, package, guest_count, deposit_amount,
              deposit_received, total_price, payment_received, payment_date, notes
              FROM events WHERE id = ?`
	
	var event Event
	var paymentDate sql.NullTime
	
	err := db.QueryRow(query, id).Scan(&event.ID, &event.EventDate,
		&event.EventName, &event.EventType, &event.StartTime,
		&event.EndTime, &event.ClientID, &event.Location,
		&event.PackageType, &event.GuestCount, &event.DepositAmount,
		&event.DepositReceived, &event.TotalPrice, &event.PaymentReceived,
		&paymentDate, &event.Notes)
	
	if err != nil {
		return nil, err
	}
	
	if paymentDate.Valid {
		event.PaymentDate = &paymentDate.Time
	}
	
	return &event, nil
}

func CreateEvent(db *sql.DB, event Event) (int64, error) {
	query := `INSERT INTO events (event_date, event_name, event_type, start_time,
              end_time, client_id, event_location, package, guest_count,
              deposit_amount, deposit_received, total_price, payment_received,
              payment_date, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	result, err := db.Exec(query, event.EventDate, event.EventName,
		event.EventType, event.StartTime, event.EndTime, event.ClientID,
		event.Location, event.PackageType, event.GuestCount,
		event.DepositAmount, event.DepositReceived, event.TotalPrice,
		event.PaymentReceived, event.PaymentDate, event.Notes)
	
	if err != nil {
		return 0, err
	}
	
	return result.LastInsertId()
}

// Search functions

func SearchClients(db *sql.DB, searchParams map[string]string) ([]Client, error) {
	query := `SELECT id, firstname, lastname, company_name, email, phone, 
              st_address, city, state, zip, date_added FROM clients WHERE 1=1`
	
	var args []interface{}
	
	if firstname := searchParams["firstname"]; firstname != "" {
		query += " AND firstname LIKE ?"
		args = append(args, "%"+firstname+"%")
	}
	if lastname := searchParams["lastname"]; lastname != "" {
		query += " AND lastname LIKE ?"
		args = append(args, "%"+lastname+"%")
	}
	if email := searchParams["email"]; email != "" {
		query += " AND email LIKE ?"
		args = append(args, "%"+email+"%")
	}
	if city := searchParams["city"]; city != "" {
		query += " AND city LIKE ?"
		args = append(args, "%"+city+"%")
	}
	
	query += " ORDER BY date_added DESC"
	
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.FirstName, &client.LastName,
			&client.CompanyName, &client.Email, &client.Phone,
			&client.Address, &client.City, &client.State, &client.Zip,
			&client.DateAdded)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}
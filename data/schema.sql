-- Create clients table
CREATE TABLE clients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(100),
    lastname VARCHAR(100),
    company_name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(20),
    st_address VARCHAR(100),
    city VARCHAR(100),
    state VARCHAR(2),
    zip VARCHAR(5),
    date_added TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create events table
CREATE TABLE events (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    event_date DATE,
    event_name VARCHAR(100),
    event_type VARCHAR(100),
    start_time TIME,
    end_time TIME,
    client_id INT,
    event_location VARCHAR(255),
    package VARCHAR(50),
    ceremony TINYINT(1) DEFAULT 0,
    deposit_amount DECIMAL(10,2) NOT NULL,
    deposit_received TINYINT(1) DEFAULT 0,
    total_price DECIMAL(10,2) NOT NULL,
    payment_received TINYINT(1) DEFAULT 0,
    payment_date DATE,
    notes VARCHAR(255),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

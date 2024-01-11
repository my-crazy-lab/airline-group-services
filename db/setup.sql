CREATE TABLE Passenger (
  id serial PRIMARY KEY,
  phone VARCHAR(20) UNIQUE NOT NULL,
  full_name VARCHAR(50) NOT NULL,
  email VARCHAR(50),
  created_at DATE,
  updated_at DATE
)

CREATE TABLE Setting (
  id serial PRIMARY KEY,
  unique_time VARCHAR(20) NOT NULL,
  created_at DATE,
  updated_at DATE
)

CREATE TABLE Flight (
  id serial PRIMARY KEY,
  FOREIGN KEY (aircraft_id) REFERENCES Aircraft (id),
  depart_time DATE NOT NULL,
  duration float8 NOT NULL,
  port float8 NOT NULL,
  price float8 NOT NULL,
  created_at DATE,
  updated_at DATE
)

CREATE TABLE Seat (
  id serial PRIMARY KEY,
  FOREIGN KEY (passenger_id) REFERENCES Passenger (id),
  FOREIGN KEY (flight_id) REFERENCES Flight (id),
  created_at DATE,
  updated_at DATE
)

CREATE TABLE Aircraft(
  id serial PRIMARY KEY,
  name VARCHAR(50) NOT NULL, 
  maintain_time DATE,
  created_at DATE,
  updated_at DATE
)


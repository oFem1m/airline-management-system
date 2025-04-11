CREATE TABLE Aircrafts
(
    id               SERIAL PRIMARY KEY,
    tail_number      VARCHAR(50) UNIQUE NOT NULL,
    model            VARCHAR(100)       NOT NULL,
    capacity         INT                NOT NULL,
    manufacture_year INT
);

CREATE TABLE Roles
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE Employees
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name  VARCHAR(50) NOT NULL,
    role_id    INT,
    hire_date  DATE,
    salary     NUMERIC(10, 2),
    email      VARCHAR(100) UNIQUE,
    phone      VARCHAR(20),
    CONSTRAINT email_check CHECK (email ~ E'^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$'),
    CONSTRAINT phone_number_check CHECK (phone ~ E'^\\+?[0-9]{10,15}$'),
    CONSTRAINT fk_employee_role FOREIGN KEY (role_id)
        REFERENCES Roles (id)
        ON DELETE SET NULL
);

CREATE TABLE Airports
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(100)   NOT NULL,
    code     CHAR(3) UNIQUE NOT NULL,
    city     VARCHAR(50),
    country  VARCHAR(50),
    timezone VARCHAR(50)
);

CREATE TABLE Route
(
    id                   SERIAL PRIMARY KEY,
    departure_airport_id INT NOT NULL,
    arrival_airport_id   INT NOT NULL,
    distance             INT, -- расстояние в км
    duration_minutes     INT, -- длительность полёта в минутах
    CONSTRAINT fk_departure_airport FOREIGN KEY (departure_airport_id)
        REFERENCES Airports (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_arrival_airport FOREIGN KEY (arrival_airport_id)
        REFERENCES Airports (id)
        ON DELETE CASCADE
);

CREATE TABLE Flights
(
    id             SERIAL PRIMARY KEY,
    flight_number  VARCHAR(20) UNIQUE NOT NULL,
    aircraft_id    INT                NOT NULL,
    route_id       INT                NOT NULL,
    departure_time TIMESTAMP          NOT NULL,
    arrival_time   TIMESTAMP          NOT NULL,
    status         VARCHAR(20), -- например: scheduled, departed, cancelled
    CONSTRAINT fk_aircraft FOREIGN KEY (aircraft_id)
        REFERENCES Aircrafts (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_route FOREIGN KEY (route_id)
        REFERENCES Route (id)
        ON DELETE CASCADE
);

CREATE TABLE Passengers
(
    id              SERIAL PRIMARY KEY,
    first_name      VARCHAR(50) NOT NULL,
    last_name       VARCHAR(50) NOT NULL,
    email           VARCHAR(100) UNIQUE,
    phone           VARCHAR(20),
    passport_number VARCHAR(50)
    CONSTRAINT email_check CHECK (email ~ E'^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$'),
    CONSTRAINT phone_number_check CHECK (phone ~ E'^\\+?[0-9]{10,15}$')
);

CREATE TABLE Booking
(
    id           SERIAL PRIMARY KEY,
    passenger_id INT       NOT NULL,
    booking_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status       VARCHAR(20), -- например: confirmed, pending, cancelled
    CONSTRAINT fk_booking_passenger FOREIGN KEY (passenger_id)
        REFERENCES Passengers (id)
        ON DELETE CASCADE
);

CREATE TABLE Tickets
(
    id           SERIAL PRIMARY KEY,
    flight_id    INT            NOT NULL,
    passenger_id INT            NOT NULL,
    booking_id   INT, -- может быть NULL, если билет не привязан к бронированию
    seat_number  VARCHAR(10),
    price        NUMERIC(10, 2) NOT NULL,
    issue_date   TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_ticket_flight FOREIGN KEY (flight_id)
        REFERENCES Flights (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_ticket_passenger FOREIGN KEY (passenger_id)
        REFERENCES Passengers (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_ticket_booking FOREIGN KEY (booking_id)
        REFERENCES Booking (id)
        ON DELETE SET NULL
);

CREATE TABLE Maintenance
(
    id                    SERIAL PRIMARY KEY,
    aircraft_id           INT       NOT NULL,
    maintenance_date      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description           TEXT,
    performed_by          INT, -- сотрудник, выполнивший обслуживание
    next_maintenance_date TIMESTAMP,
    CONSTRAINT fk_maintenance_aircraft FOREIGN KEY (aircraft_id)
        REFERENCES Aircrafts (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_maintenance_employee FOREIGN KEY (performed_by)
        REFERENCES Employees (id)
        ON DELETE SET NULL
);

CREATE TABLE FlightCrew
(
    flight_id   INT NOT NULL,
    employee_id INT NOT NULL,
    PRIMARY KEY (flight_id, employee_id),
    CONSTRAINT fk_flight FOREIGN KEY (flight_id)
        REFERENCES Flights (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_employee FOREIGN KEY (employee_id)
        REFERENCES Employees (id)
        ON DELETE CASCADE
);

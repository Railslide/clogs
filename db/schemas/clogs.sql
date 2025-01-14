CREATE DATABASE IF NOT EXISTS clogs CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE clogs;

CREATE TABLE IF NOT EXISTS gyms (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  location VARCHAR(255) NOT NULL,
  country VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS routes (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  gym INTEGER NOT NULL,
  grade VARCHAR(10) NOT NULL,
  route_type VARCHAR(255) NOT NULL,
  notes TEXT NOT NULL,
  sent BOOLEAN NOT NULL DEFAULT 0,
  archived BOOLEAN NOT NULL DEFAULT 0,
  FOREIGN KEY (gym) REFERENCES gyms(id)
);

CREATE INDEX idx_routes_grade ON routes(grade);

CREATE TABLE IF NOT EXISTS route_attempts (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  route INTEGER NOT NULL,
  created DATETIME NOT NULL,
  send BOOLEAN NOT NULL DEFAULT 0,
  all_moves BOOLEAN NOT NULL DEFAULT 0,
  FOREIGN KEY (route) REFERENCES routes(id)
);

CREATE INDEX idx_attempts_route ON attempts(route);

-- Gym data --
INSERT INTO gyms (name, location, country) VALUES (
  "Karbin",
  "Stockholm",
  "Sweden"
);

INSERT INTO gyms (name, location, country) VALUES (
  "Telefonplan",
  "Stockholm",
  "Sweden"
);

INSERT INTO gyms (name, location, country) VALUES (
  "Akalla",
  "Stockholm",
  "Sweden"
);

CREATE TABLE actors (
    actor_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    gender VARCHAR(255),
    date_of_birth DATE
);

CREATE TABLE movies (
    movie_id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    date DATE,
    rating DECIMAL(3,1) CHECK (rating >= 0 AND rating <= 10)
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    isAdmin BOOLEAN
);

INSERT INTO users (username, password_hash, isAdmin) VALUES('userAdmin', 'password', true);
INSERT INTO users (username, password_hash, isAdmin) VALUES('user', 'password', false);

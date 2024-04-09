CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    surname VARCHAR(255),
    patronymic VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS car (
    id SERIAL PRIMARY KEY,
    reg_num VARCHAR(255) NOT NULL UNIQUE,
    mark VARCHAR(255),
    model VARCHAR(255),
    year INT,
    owner_id INT REFERENCES people(id)
);

CREATE INDEX idx_cars_fulltext ON car USING GIN (to_tsvector('russian', reg_num || ' ' || mark || ' ' || model || ' ' || year::text));

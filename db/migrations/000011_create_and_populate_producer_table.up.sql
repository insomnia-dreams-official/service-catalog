CREATE TABLE IF NOT EXISTS producer
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

INSERT INTO producer (id, name) VALUES (1, 'УПАКС-ЮНИТИ');
INSERT INTO producer (id, name) VALUES (2, 'Юпласт');
INSERT INTO producer (id, name) VALUES (3, 'Покровский полимер');
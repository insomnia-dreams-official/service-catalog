CREATE TABLE IF NOT EXISTS country
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

INSERT INTO country (id, name) VALUES (1, 'Россия');
INSERT INTO country (id, name) VALUES (2, 'Китай');
INSERT INTO country (id, name) VALUES (3, 'Германия');
INSERT INTO country (id, name) VALUES (4, 'Малайзия');
INSERT INTO country (id, name) VALUES (5, 'Финляндия');
INSERT INTO country (id, name) VALUES (6, 'Беларусь');
INSERT INTO country (id, name) VALUES (7, 'Бельгия');
INSERT INTO country (id, name) VALUES (8, 'Польша');
INSERT INTO country (id, name) VALUES (9, 'Венгрия');
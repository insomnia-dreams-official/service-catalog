CREATE TABLE IF NOT EXISTS delivery_quant
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

INSERT INTO delivery_quant (id, name) VALUES (1, 'штука');
INSERT INTO delivery_quant (id, name) VALUES (3, 'коробка');
INSERT INTO delivery_quant (id, name) VALUES (2, 'упаковка');
INSERT INTO delivery_quant (id, name) VALUES (4, 'пара');
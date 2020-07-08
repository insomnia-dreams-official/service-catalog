CREATE TABLE IF NOT EXISTS product
(
    articul          CHAR(8) PRIMARY KEY,
    name             VARCHAR(200) NOT NULL,
    category_articul CHAR(8) REFERENCES category (articul),
    status_id        SMALLINT
);

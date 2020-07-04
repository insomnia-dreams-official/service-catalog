CREATE TABLE IF NOT EXISTS category
(
    articul   CHAR(8) PRIMARY KEY,
    name      VARCHAR(200) NOT NULL,
    path      VARCHAR(10),
    link      VARCHAR(200),
    full_link VARCHAR(300)
);

CREATE TABLE IF NOT EXISTS category_seo
(
    category_articul CHAR(8) PRIMARY KEY,
    trans_link       VARCHAR(200),
    meta_description varchar(500),
    meta_keywords    varchar(500),
    description      varchar(2000)
);


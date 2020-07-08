CREATE TABLE IF NOT EXISTS product_base_attributes
(
    product_articul   VARCHAR(8) PRIMARY KEY,
    producer_id       INT REFERENCES producer (id),
    country_id        INT REFERENCES country (id),
    minimum_delivery  INT,
    pieces_in_pack    INT,
    pieces_in_box     INT,
    packs_in_box      INT,
    volume            FLOAT,
    weight            FLOAT,
    delivery_quant_id INT REFERENCES delivery_quant (id)
);

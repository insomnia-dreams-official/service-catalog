CREATE TABLE product_price
(
    product_articul             CHAR(8) PRIMARY KEY REFERENCES product (articul),
    stock_price                 FLOAT,
    small_wholesale_price       FLOAT,
    small_wholesale_delay_price FLOAT,
    wholesale_delay_price       FLOAT,
    wholesale_price             FLOAT,
    wholesale_artificial_price  FLOAT,
    sale_price                  FLOAT,
    retail_warehouse_price      FLOAT,
    special_price               FLOAT,
    special_offer_price         FLOAT
);

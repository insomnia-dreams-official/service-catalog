DO
$$
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_constraint WHERE conname = 'product_base_attributes_product_articul_fkey') THEN
            ALTER TABLE product_base_attributes
                ADD CONSTRAINT product_base_attributes_product_articul_fkey
                    FOREIGN KEY (product_articul) REFERENCES product (articul)
                        ON DELETE NO ACTION
                        ON UPDATE NO ACTION
                    NOT VALID;
        END IF;
    END;
$$;
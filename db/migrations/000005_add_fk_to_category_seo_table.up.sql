DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_category_seo_to_category') THEN
            ALTER TABLE category_seo
                ADD CONSTRAINT fk_category_seo_to_category
                    FOREIGN KEY (category_articul) REFERENCES category (articul)
                        ON DELETE NO ACTION
                        ON UPDATE NO ACTION
                    NOT VALID;
        END IF;
    END;
$$;
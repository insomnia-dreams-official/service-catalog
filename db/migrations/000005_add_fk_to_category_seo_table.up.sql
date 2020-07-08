DO
$$
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_constraint WHERE conname = 'category_seo_category_articul_fkey') THEN
            ALTER TABLE category_seo
                ADD CONSTRAINT category_seo_category_articul_fkey
                    FOREIGN KEY (category_articul) REFERENCES category (articul)
                        ON DELETE NO ACTION
                        ON UPDATE NO ACTION
                    NOT VALID;
        END IF;
    END;
$$;
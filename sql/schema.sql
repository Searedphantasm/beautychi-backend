CREATE TABLE IF NOT EXISTS category
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    slug                   VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    image       TEXT NOT NULL,
    image_key   VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sub_category
(
    id                 SERIAL PRIMARY KEY,
    parent_category_id INT REFERENCES category (id) ON DELETE RESTRICT ,
    name               VARCHAR(255) NOT NULL,
    slug                   VARCHAR(255) NOT NULL,
    description        TEXT         NOT NULL,
    image              TEXT NOT NULL,
    image_key          VARCHAR(255) NOT NULL,
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS brand
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255),
    slug                   VARCHAR(255) NOT NULL,
    description TEXT,
    country     VARCHAR(100),
    logo        TEXT,
    logo_key    VARCHAR(255),
    website_url VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_status') THEN
            CREATE TYPE product_status AS ENUM ('Active', 'Inactive');
        END IF;
    END $$;

CREATE TABLE IF NOT EXISTS product
(
    id                     SERIAL PRIMARY KEY,
    title                  VARCHAR(255) NOT NULL,
    slug                   VARCHAR(255) NOT NULL,
    description            TEXT,
    poster                 TEXT NOT NULL,
    poster_key             VARCHAR(255) NOT NULL,
    price                  INT          NOT NULL,
    category_id            INT REFERENCES category (id) ON DELETE RESTRICT ,
    brand_id               INT REFERENCES brand (id) ON DELETE RESTRICT ,
    product_stock          INT          NOT NULL,
    product_discount_price INT,
    sub_category_id        INT REFERENCES sub_category (id) ON DELETE RESTRICT ,
    consumer_guide         TEXT,
    contact                VARCHAR(11),
    status                 product_status,
    created_at             TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at             TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE IF NOT EXISTS product_specifications
(
    id                SERIAL PRIMARY KEY,
    product_id        INT REFERENCES product (id) ON DELETE CASCADE ,
    specs_title       VARCHAR(255) NOT NULL,
    specs_description VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS product_image
(
    id         SERIAL PRIMARY KEY,
    product_id INT REFERENCES product (id) ON DELETE CASCADE ,
    url TEXT ,
    url_key VARCHAR(255),
    alt_text VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user (
    id  SERIAL PRIMARY KEY ,
    first_name VARCHAR(100) NOT NULL ,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(11) NOT NULL UNIQUE,
    is_admin BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_address (
    id  SERIAL PRIMARY KEY ,
    user_id INT REFERENCES user (id) ON DELETE CASCADE ,
    city VARCHAR(255) NOT NULL ,
    state VARCHAR(255) NOT NULL ,
    address TEXT NOT NULL ,
    postal_code VARCHAR(11) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX idx_product_title ON product(slug);
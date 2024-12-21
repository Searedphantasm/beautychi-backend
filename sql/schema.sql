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

CREATE TYPE product_status AS ENUM ('Active','Inactive');

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
)
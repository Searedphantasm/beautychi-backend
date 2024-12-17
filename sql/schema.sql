CREATE TABLE IF NOT EXISTS category
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    image       VARCHAR(255) NOT NULL,
    image_key   VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sub_category
(
    id                 SERIAL PRIMARY KEY,
    parent_category_id INT REFERENCES category (id),
    name               VARCHAR(255) NOT NULL,
    description        TEXT         NOT NULL,
    image              VARCHAR(255) NOT NULL,
    image_key          VARCHAR(255) NOT NULL,
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS brand
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255),
    description TEXT,
    country     VARCHAR(100),
    logo        VARCHAR(255),
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
    poster                 VARCHAR(255) NOT NULL,
    poster_key             VARCHAR(255) NOT NULL,
    price                  INT          NOT NULL,
    category_id            INT REFERENCES category (id),
    brand_id               INT REFERENCES brand (id),
    product_stock          INT          NOT NULL,
    product_discount_price INT,
    sub_category_id        INT REFERENCES sub_category (id),
    consumer_guide         TEXT,
    contact                VARCHAR(11),
    status                 product_status,
    created_at             TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at             TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE IF NOT EXISTS product_specifications
(
    id                SERIAL PRIMARY KEY,
    product_id        INT REFERENCES product (id),
    specs_title       VARCHAR(255) NOT NULL,
    specs_description VARCHAR(255) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS product_image
(
    id         SERIAL PRIMARY KEY,
    product_id INT REFERENCES product (id),
    url TEXT ,
    alt_text VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
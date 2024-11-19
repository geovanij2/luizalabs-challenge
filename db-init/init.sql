CREATE TABLE IF NOT EXISTS clients (
    id TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS favorite_products (
    client_id TEXT NOT NULL,
    product_id TEXT NOT NULL,
    price BIGINT NOT NULL,
    image TEXT NOT NULL,
    brand TEXT NOT NULL,
    title TEXT NOT NULL,
    review_score FLOAT NOT NULL,
    PRIMARY KEY (client_id, product_id),
    FOREIGN KEY (client_id) REFERENCES clients (id) ON DELETE CASCADE
);

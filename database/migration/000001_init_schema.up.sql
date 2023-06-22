CREATE TABLE "orders"
(
    "id"         bigserial PRIMARY KEY,
    "user_id"    int UNIQUE NOT NULL,
    "status"     varchar    NOT NULL,
    "created_at" date       NOT NULL,
    "updated_at" date       NOT NULL
);

CREATE TABLE "order_items"
(
    "order_id"   bigserial NOT NULL,
    "product_id" bigserial NOT NULL,
    "quantity"   int       NOT NULL
);

CREATE TABLE "products"
(
    "id"          bigserial PRIMARY KEY,
    "name"        varchar NOT NULL,
    "merchant_id" int     NOT NULL,
    "price"       int     NOT NULL,
    "status"      varchar NOT NULL,
    "created_at"  date    NOT NULL,
    "updated_at"  date    NOT NULL,
    "category_id" int     NOT NULL
);

CREATE TABLE "users"
(
    "id"            bigserial PRIMARY KEY,
    "full_name"     varchar        NOT NULL,
    "email"         varchar UNIQUE NOT NULL,
    "gender"        varchar        NOT NULL,
    "date_of_birth" date           NOT NULL,
    "created_at"    date           NOT NULL,
    "updated_at"    date           NOT NULL,
    "country_code"  varchar        NOT NULL
);

CREATE TABLE "merchants"
(
    "id"            bigserial PRIMARY KEY,
    "admin_id"      int     NOT NULL,
    "merchant_name" varchar NOT NULL,
    "country_code"  varchar NOT NULL,
    "created_at"    date    NOT NULL,
    "updated_at"    date    NOT NULL

);

CREATE TABLE "categories"
(
    "id"       bigserial PRIMARY KEY,
    "cat_name" varchar NOT NULL
);

CREATE TABLE "countries"
(
    "code"           varchar PRIMARY KEY,
    "name"           varchar NOT NULL,
    "continent_name" varchar NOT NULL
);

ALTER TABLE "orders"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products"
    ADD FOREIGN KEY ("merchant_id") REFERENCES "merchants" ("id");

ALTER TABLE "products"
    ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "users"
    ADD FOREIGN KEY ("country_code") REFERENCES "countries" ("code");

ALTER TABLE "merchants"
    ADD FOREIGN KEY ("admin_id") REFERENCES "users" ("id");

ALTER TABLE "merchants"
    ADD FOREIGN KEY ("country_code") REFERENCES "countries" ("code");

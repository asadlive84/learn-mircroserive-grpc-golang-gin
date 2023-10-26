-- +goose Up

CREATE TABLE IF NOT EXISTS product(
   id uuid DEFAULT uuid_generate_v4 (),
   name VARCHAR (600) NOT NULL,
   description VARCHAR (20) NOT NULL DEFAULT '',
   is_active BOOLEAN DEFAULT FALSE,
   created timestamp default current_timestamp,
   updated timestamp NULL,
   created_by uuid NOT NULL
);

-- +goose Down

DROP TABLE product;


CREATE TABLE "customers" (
  "customer_id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "date_of_birth" timestamp DEFAULT 'now()',
  "city" varchar NOT NULL,
  "zipcode" varchar NOT NULL,
  "status" SMALLINT NOT NULL DEFAULT 1
);

ALTER SEQUENCE customers_customer_id_seq RESTART WITH 2006
CREATE INDEX ON "customers" ("status");
CREATE INDEX ON "customers" ("customer_id");
CREATE INDEX ON "customers" ("name");


CREATE TABLE "accounts" (
  "account_id" bigserial PRIMARY KEY NOT NULL,
  "customer_id" bigint NOT NULL,
  "opening_date" timestamp NOT NULL DEFAULT 'now()',
  "account_type" varchar NOT NULL,
  "amount" decimal NOT NULL,
  "status" SMALLINT NOT NULL
);

CREATE INDEX ON "accounts" ("customer_id");
CREATE INDEX ON "accounts" ("account_id");
CREATE INDEX ON "accounts" ("account_type");
CREATE INDEX ON "accounts" ("status");

ALTER TABLE "accounts" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");
ALTER SEQUENCE accounts_account_id_seq RESTART WITH 95476;


CREATE TABLE "transactions" (
  "transaction_id" bigserial PRIMARY KEY NOT NULL,
  "account_id" bigint NOT NULL,
  "amount" decimal NOT NULL,
  "transaction_type" varchar NOT NULL,
  "transaction_date" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "customers" ("customer_id");
CREATE INDEX ON "transactions" ("account_id");
CREATE INDEX ON "transactions" ("transaction_type");

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "customer_id" bigint NOT NULL,
  "created_on" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "users" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("role");

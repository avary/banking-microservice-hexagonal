CREATE TABLE "customers" (
  "customer_id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "date_of_birth" timestamp DEFAULT (now()),
  "city" varchar NOT NULL,
  "zipcode" varchar NOT NULL,
  "status" SMALLINT NOT NULL DEFAULT 1
);

CREATE TABLE "accounts" (
  "account_id" bigserial PRIMARY KEY NOT NULL,
  "customer_id" bigint NOT NULL,
  "opening_date" timestamp NOT NULL DEFAULT 'now()',
  "account_type" varchar NOT NULL,
  "amount" decimal NOT NULL,
  "status" SMALLINT NOT NULL DEFAULT 1
);

CREATE TABLE "transactions" (
  "transaction_id" bigserial PRIMARY KEY NOT NULL,
  "account_id" bigint NOT NULL,
  "amount" decimal NOT NULL,
  "transaction_type" varchar NOT NULL,
  "transaction_date" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "customer_id" bigint,
  "created_on" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "users" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");

CREATE INDEX ON "customers" ("status");

CREATE INDEX ON "customers" ("customer_id");

CREATE INDEX ON "customers" ("name");

CREATE INDEX ON "accounts" ("customer_id");

CREATE INDEX ON "accounts" ("account_id");

CREATE INDEX ON "accounts" ("account_type");

CREATE INDEX ON "accounts" ("status");

CREATE INDEX ON "transactions" ("account_id");

CREATE INDEX ON "transactions" ("transaction_type");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("role");


INSERT INTO customers VALUES
                            (2000,'Steve','1978-12-15','Delhi','110075',1),
                            (2001,'Arian','1988-05-21','Newburgh, NY','12550',1),
                            (2002,'Hadley','1988-04-30','Inglewood, NJ','07631',1),
                            (2003,'Ben','1988-01-04','Manchester, NH','03102',0),
                            (2004,'Nina','1988-05-14','Blackstone, MI','48348',1),
                            (2005,'Osman','1988-11-08','Hyattsville, MD','20782',0);
ALTER SEQUENCE customers_customer_id_seq RESTART WITH 2006;

INSERT INTO accounts VALUES
                           (95470,2000,'2020-08-22 10:20:06', 'saving', 6823.23, 1),
                           (95471,2002,'2020-08-09 10:27:22', 'checking', 3342.96, 1),
                           (95472,2001,'2020-08-09 10:35:22', 'saving', 7000, 1),
                           (95473,2001,'2020-08-09 10:38:22', 'saving', 5861.86, 1);

ALTER SEQUENCE accounts_account_id_seq RESTART WITH 95474;

INSERT INTO users VALUES
                        ('admin','abc123','admin', NULL, '2020-08-09 10:27:22'),
                        ('2001','abc123','user', 2001, '2020-08-09 10:27:22'),
                        ('2000','abc123','user', 2002, '2020-08-09 10:27:22');
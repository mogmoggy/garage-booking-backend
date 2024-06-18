CREATE TABLE
  "user" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" VARCHAR(50) NOT NULL,
    "email" VARCHAR(50) UNIQUE NOT NULL,
    "phone_number" VARCHAR(20) UNIQUE NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  "car" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "car_registration" VARCHAR(20) UNIQUE NOT NULL,
    "make" VARCHAR(20) NOT NULL,
    "fuel_type" VARCHAR(20) NOT NULL,
    "year_manufacture" INT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  "booking" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "car_id" BIGINT NOT NULL,
    "booking_date" TIMESTAMP NOT NULL DEFAULT (now ()),
    "problem_description" TEXT,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now ())
  );

ALTER TABLE "car" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "booking" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "booking" ADD FOREIGN KEY ("car_id") REFERENCES "car" ("id");
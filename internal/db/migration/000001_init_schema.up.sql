CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
    "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "name" varchar unique NOT NULL,
    "phone" varchar unique NOT NULL,
    "email" varchar unique NOT NULL,
    "hashed_password" varchar NOT NULL,
    "role" varchar DEFAULT '"user"',
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT NOW() + INTERVAL '7 hours',
    "updated_at" timestamptz NOT NULL DEFAULT NOW() + INTERVAL '7 hours'
);

-- CREATE INDEX ON "users" ("phone", "email");
ALTER TABLE "users" ADD CONSTRAINT "phone_email_key" UNIQUE ("phone", "email");
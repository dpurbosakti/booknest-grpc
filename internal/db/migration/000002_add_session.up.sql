CREATE TABLE "sessions" (
    "id" uuid PRIMARY KEY,
    "name" varchar NOT NULL,
    "refresh_token" varchar NOT NULL,
    "user_agent" varchar NOT NULL,
    "client_ip" varchar NOT NULL,
    "is_blocked" boolean NOT NULL DEFAULT false,
    "expired_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW() + INTERVAL '7 hours'
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("name") REFERENCES "users" ("name");
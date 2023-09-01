CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "urls" (
  "id" serial PRIMARY KEY,
  "user_id" integer NOT NULL,
  "original_url" varchar NOT NULL,
  "short_url" varchar,
  "created_at" timestamptz DEFAULT (now())
);


CREATE INDEX ON "urls" ("short_url", "user_id");

CREATE INDEX ON "urls" ("short_url");

ALTER TABLE "urls" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

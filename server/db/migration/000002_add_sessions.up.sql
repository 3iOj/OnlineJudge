CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "username" varchar(100) NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "refresh_token" varchar NOT NULL,
  "profileimg" text,
  "motto" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "dob" timestamptz NOT NULL,
  "rating" integer,
  "problem_solved" integer,
  "admin_id" bigint,
  "is_setter" boolean NOT NULL DEFAULT false
);
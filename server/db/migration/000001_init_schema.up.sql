CREATE TABLE "admin" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "username" varchar(100) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "requests" (
  "username" varchar(100),
  "admin_name" varchar(100),
  "permission_asked" varchar(100),
  "current_status" varchar(100),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("username", "admin_name")
);

CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "username" varchar(100) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "profileimg" text,
  "motto" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "dob" timestamptz NOT NULL,
  "rating" integer,
  "problem_solved" integer,
  "admin_id" bigint,
  "is_setter" boolean NOT NULL DEFAULT false
);

CREATE TABLE "problems" (
  "id" BIGSERIAL PRIMARY KEY,
  "problem_name" varchar(255) NOT NULL,
  "description" text NOT NULL,
  "sample_input" text NOT NULL,
  "sample_output" text NOT NULL,
  "ideal_solution" text NOT NULL,
  "time_limit" integer NOT NULL,
  "memory_limit" integer NOT NULL,
  "code_size" integer NOT NULL,
  "rating" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "contest_id" bigint NOT NULL
);

CREATE TABLE "problem_tests" (
  "id" bigint PRIMARY KEY,
  "problem_id" bigint NOT NULL,
  "input" text NOT NULL,
  "output" text NOT NULL
);

CREATE TABLE "problem_creators" (
  "problem_id" bigint,
  "created_by" varchar(100),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("problem_id", "created_by")
);

CREATE TABLE "prob_tags" (
  "problem_id" bigint,
  "tag" varchar(255),
  PRIMARY KEY ("problem_id", "tag")
);

CREATE TABLE "blogs" (
  "id" BIGSERIAL PRIMARY KEY,
  "blog_title" text NOT NULL,
  "blog_content" text NOT NULL,
  "created_by" varchar(100) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "publish_at" timestamptz NOT NULL DEFAULT (now()),
  "votes_count" integer
);

CREATE TABLE "blog_tags" (
  "blog_id" bigint,
  "tag" varchar(255),
  PRIMARY KEY ("blog_id", "tag")
);

CREATE TABLE "blog_comments" (
  "id" BIGSERIAL PRIMARY KEY,
  "blog_id" bigint NOT NULL,
  "message" text NOT NULL,
  "commented_by" varchar(100) NOT NULL,
  "child_comment" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "blog_likes" (
  "blog_id" bigint,
  "action_by" varchar(100),
  "is_liked" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("blog_id", "action_by")
);

CREATE TABLE "contests" (
  "id" BIGSERIAL PRIMARY KEY,
  "contest_name" text NOT NULL,
  "start_time" timestamptz,
  "end_time" timestamptz,
  "duration" bigint NOT NULL,
  "registration_start" timestamptz,
  "registration_end" timestamptz,
  "announcement_blog" bigint,
  "editorial_blog" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "ispublish"  boolean DEFAULT false
);

CREATE TABLE "contest_creators" (
  "contest_id" bigint,
  "creator_name" varchar(255),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("contest_id", "creator_name")
);

CREATE TABLE "submissions" (
  "id" BIGSERIAL PRIMARY KEY,
  "submitted_at" timestamptz NOT NULL DEFAULT (now()),
  "problem_id" bigint NOT NULL,
  "username" varchar(100) NOT NULL,
  "user_id" bigint NOT NULL,
  "contest_id" bigint NOT NULL,
  "language" varchar(255) NOT NULL,
  "verdict" varchar(40) NOT NULL,
  "code" text NOT NULL,
  "exec_time" integer NOT NULL,
  "memory_consumed" integer NOT NULL,
  "score" integer NOT NULL
);

CREATE TABLE "contest_registered" (
  "contest_id" bigint,
  "username" varchar(100),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("contest_id", "username")
);

CREATE TABLE "submission_tests" (
  "id" bigint PRIMARY KEY,
  "submission_id" bigint,
  "input" text NOT NULL,
  "output" text NOT NULL
);

CREATE TABLE "community" (
  "id" BIGSERIAL PRIMARY KEY,
  "community_name" varchar(255) NOT NULL,
  "community_admin" varchar(100) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "community_users" (
  "id" bigint PRIMARY KEY,
  "community_id" bigint NOT NULL,
  "username" varchar(100) NOT NULL
);

CREATE INDEX ON "problems" ("problem_name");

CREATE INDEX ON "blogs" ("blog_title");

CREATE INDEX ON "contests" ("contest_name");

CREATE INDEX ON "submissions" ("username");

CREATE INDEX ON "submissions" ("language");

COMMENT ON COLUMN "problems"."ideal_solution" IS 'to generate the output files in problemTests';

COMMENT ON COLUMN "problems"."time_limit" IS 'should be in seconds';

COMMENT ON COLUMN "problems"."memory_limit" IS 'should be in MB';

COMMENT ON COLUMN "problems"."code_size" IS 'should be in KB';

COMMENT ON COLUMN "contests"."end_time" IS 'must be greater than start time';

COMMENT ON COLUMN "contests"."duration" IS 'must be equal to difference between end time and start time';

COMMENT ON COLUMN "contests"."registration_end" IS 'must be greater than registration_start';

COMMENT ON COLUMN "contests"."announcement_blog" IS 'should be created automatically';

COMMENT ON COLUMN "contests"."editorial_blog" IS 'should be created automatically';

COMMENT ON COLUMN "submissions"."submitted_at" IS 'fetch solutions based on this time during a contest';

COMMENT ON COLUMN "submissions"."problem_id" IS 'will come handy in contest score calculation';

COMMENT ON COLUMN "submissions"."user_id" IS 'will come handy in contest score calculation';

COMMENT ON COLUMN "submissions"."contest_id" IS 'will come handy in contest score calculation';

COMMENT ON COLUMN "submissions"."exec_time" IS 'should be in seconds';

COMMENT ON COLUMN "submissions"."memory_consumed" IS 'should be in MB';

COMMENT ON COLUMN "submissions"."score" IS 'must be calculated in application logic';

ALTER TABLE "users" ADD FOREIGN KEY ("admin_id") REFERENCES "admin" ("id");

ALTER TABLE "problems" ADD FOREIGN KEY ("contest_id") REFERENCES "contests" ("id");

ALTER TABLE "blogs" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("username");

ALTER TABLE "blog_comments" ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id");

ALTER TABLE "blog_comments" ADD FOREIGN KEY ("commented_by") REFERENCES "users" ("username");

ALTER TABLE "blog_comments" ADD FOREIGN KEY ("child_comment") REFERENCES "blog_comments" ("id");

ALTER TABLE "contests" ADD FOREIGN KEY ("announcement_blog") REFERENCES "blogs" ("id");

ALTER TABLE "contests" ADD FOREIGN KEY ("editorial_blog") REFERENCES "blogs" ("id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("problem_id") REFERENCES "problems" ("id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "submissions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("contest_id") REFERENCES "contests" ("id");

ALTER TABLE "community" ADD FOREIGN KEY ("community_admin") REFERENCES "users" ("username");

ALTER TABLE "community_users" ADD FOREIGN KEY ("community_id") REFERENCES "community" ("id");

ALTER TABLE "community_users" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "prob_tags" ADD FOREIGN KEY ("problem_id") REFERENCES "problems" ("id");

ALTER TABLE "blog_tags" ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id");

ALTER TABLE "blog_likes" ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id");

ALTER TABLE "blog_likes" ADD FOREIGN KEY ("action_by") REFERENCES "users" ("username");

ALTER TABLE "contest_creators" ADD FOREIGN KEY ("contest_id") REFERENCES "contests" ("id");

ALTER TABLE "contest_creators" ADD FOREIGN KEY ("creator_name") REFERENCES "users" ("username");

ALTER TABLE "problem_creators" ADD FOREIGN KEY ("problem_id") REFERENCES "problems" ("id");

ALTER TABLE "problem_creators" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("username");

ALTER TABLE "problem_tests" ADD FOREIGN KEY ("problem_id") REFERENCES "problems" ("id");

ALTER TABLE "submission_tests" ADD FOREIGN KEY ("submission_id") REFERENCES "submissions" ("id");

ALTER TABLE "contest_registered" ADD FOREIGN KEY ("contest_id") REFERENCES "contests" ("id");

ALTER TABLE "contest_registered" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "requests" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "requests" ADD FOREIGN KEY ("admin_name") REFERENCES "admin" ("username");

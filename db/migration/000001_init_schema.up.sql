CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "profile_picture" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_media" (
  "media_id" bigserial PRIMARY KEY,
  "post_id" bigserial NOT NULL,
  "media_file" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "post_id" bigserial PRIMARY KEY,
  "created_by_user_id" bigserial NOT NULL,
  "caption" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_likes" (
  "post_id" bigserial NOT NULL,
  "post_liker" varchar NOT NULL
);

CREATE TABLE "comments" (
  "comment_id" bigserial PRIMARY KEY,
  "created_by_user_id" bigserial NOT NULL,
  "post_id" bigserial NOT NULL,
  "comment" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "follower" (
  "user_id" bigserial NOT NULL,
  "followed_user_id" bigserial NOT NULL
);

ALTER TABLE "post_media" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id");

ALTER TABLE "posts" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "post_likes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id");

ALTER TABLE "post_likes" ADD FOREIGN KEY ("post_liker") REFERENCES "users" ("username");

ALTER TABLE "comments" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id");

ALTER TABLE "follower" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "follower" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("user_id");

CREATE TABLE users (
 id  uuid primary key,
 auth_id text not null,
 email varchar,
 name varchar,
 role varchar,
 image_url text,
 created_at  timestamp,
 updated_at  timestamp,
 deleted_at  timestamp
);
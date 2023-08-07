CREATE TABLE users (
	id uuid PRIMARY KEY,
  username VARCHAR NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
	deleted_at TIMESTAMPTZ NOT NULL
);

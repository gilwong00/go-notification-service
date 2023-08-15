CREATE TABLE followers (
	id uuid PRIMARY KEY,
	user_id uuid NOT NULL REFERENCES users (id),
	follower_id uuid NOT NULL REFERENCES users (id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
	deleted_at TIMESTAMPTZ NOT NULL
);

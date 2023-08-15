CREATE TABLE notification_queue (
	id uuid PRIMARY KEY,
	message VARCHAR NOT NULL,
	follower_id uuid NOT NULL REFERENCES users (id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
	state_id uuid NOT NULL,
	attempts integer DEFAULT 0,
	updated_at TIMESTAMPTZ NOT NULL
);

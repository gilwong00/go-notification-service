CREATE TABLE notification_queue (
	id uuid PRIMARY KEY,
	message VARCHAR NOT NULL,
	follower_id uuid NOT NULL REFERENCES users (id),
	state_id uuid NOT NULL,
	attempts integer DEFAULT 0,
	created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
	updated_at TIMESTAMPTZ NOT NULL
);

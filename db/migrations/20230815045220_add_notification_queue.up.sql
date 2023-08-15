CREATE TABLE notification_queue (
	id uuid PRIMARY KEY,
  message VARCHAR NOT NULL,
	follower_id VARCHAR NOT NULL REFERENCES users (id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
	state_id varchar NOT NULL REFERENCES notification_state (id),
  attempts integer DEFAULT 0,
	updated_at TIMESTAMPTZ NOT NULL
);

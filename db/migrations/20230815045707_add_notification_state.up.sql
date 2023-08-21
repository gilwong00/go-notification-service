CREATE TYPE state AS ENUM ('pending','success','failed');

CREATE TABLE notification_state (
	id uuid PRIMARY KEY,
	state state NOT NULL,
	message VARCHAR,
	requested_at_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
	completed_at TIMESTAMPTZ
);

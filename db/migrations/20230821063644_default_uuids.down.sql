ALTER TABLE users ALTER COLUMN id DROP DEFAULT;

ALTER TABLE followers ALTER COLUMN id DROP DEFAULT;

ALTER TABLE notification_state ALTER COLUMN id DROP DEFAULT;

ALTER TABLE notification_queue ALTER COLUMN id DROP DEFAULT;
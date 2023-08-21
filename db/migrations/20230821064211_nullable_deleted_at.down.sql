ALTER TABLE users ALTER COLUMN deleted_at SET NOT NULL;

ALTER TABLE followers ALTER COLUMN deleted_at SET NOT NULL;

ALTER TABLE notification_queue ALTER COLUMN updated_at SET NOT NULL;

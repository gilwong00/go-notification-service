ALTER TABLE users ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE followers ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE notification_state ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE notification_queue ALTER COLUMN id SET DEFAULT gen_random_uuid();

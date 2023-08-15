ALTER TABLE notification_queue
ADD CONSTRAINT fk_notification_state FOREIGN KEY (state_id)
REFERENCES notification_state (id);
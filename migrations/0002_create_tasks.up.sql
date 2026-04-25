ALTER TABLE tasks
ADD COLUMN recurrence_type TEXT NOT NULL DEFAULT 'none',
ADD COLUMN recurrence_value TEXT NOT NULL DEFAULT '';
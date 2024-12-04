CREATE TABLE IF NOT EXISTS groupps (
	groupp UUID,
	group_name VARCHAR(32),
	rep_id INT
);
CREATE INDEX IF NOT EXISTS idx_id ON groupps (group_name);

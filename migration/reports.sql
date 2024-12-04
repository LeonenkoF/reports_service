CREATE TABLE IF NOT EXISTS reports (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	uuid UUID,
	user_uuid UUID,
	description TEXT,
	priority PRIORITY_TYPE,
	stage STAGE_TYPE,
	category INT,
	groupp UUID,
	groupp_name VARCHAR(32),
	adm_uuid UUID,
	closing_comment TEXT,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);
CREATE INDEX IF NOT EXISTS idx_created_at ON reports (created_at);

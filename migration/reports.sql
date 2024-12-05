	/*
	reports - таблица жалоб
	--------------------------------------------------
		id				порядковый номер жалобы
		uuid 			идентификация жалобы (по ТЗ)
		user_uuid 		индентификатор автора-пользователя
		description 	текст жалобы
		priority		приоритет ('high','medium','low')
		stage			стадия обработки ('new','inprogress','done', 'canceled')
		category		категория (тема) жалобы.
		groupp 			группа жалобы. Используется при объединении жалоб в группы.
	*/
CREATE TABLE IF NOT EXISTS reports (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	uuid UUID,
	user_uuid UUID,
	description TEXT,
	priority PRIORITY_TYPE,
	stage STAGE_TYPE,
	category INT,
	groupp UUID,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);
CREATE INDEX IF NOT EXISTS idx_created_at ON reports (created_at);

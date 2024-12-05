	/*
	reports - таблица жалоб
	--------------------------------------------------
		id				порядковый номер жалобы
		uuid 			идентификация жалобы (по ТЗ)
		user_uuid 		индентификатор автора-пользователя
		description 	текст жалобы
		priority		приоритет ('high','medium','low')
		stage			стадия обработки ('new','inprogress','done', 'canceled')
		created_at 		дата создания
		updated_at 		дата обновления
	*/
CREATE TABLE IF NOT EXISTS reports (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	uuid UUID,
	user_uuid UUID,
	description TEXT,
	priority PRIORITY_TYPE,
	stage STAGE_TYPE,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);
CREATE INDEX IF NOT EXISTS idx_created_at ON reports (created_at);
CREATE INDEX IF NOT EXISTS idx_id ON reports (id);


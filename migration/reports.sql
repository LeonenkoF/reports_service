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
						*** Пока реализована таблица, т к в ТЗ не указаны категории
		groupp 			группа жалобы. Используется при объединении жалоб в группы.
						*** Варианты:
						1. Каждой группе своя таблица с номерами жалоб, как и в случае хранения диалога по жалобе (comments)
						2. Таблица одна с записями (группа, имя группы, жалоба). Тогда будет дублирование данных по группам. Возможно легче обработка (поиск).
		adm_uuid		идентификатор админа, который закрывал жалобу и писал closing_comment
						*** Возможно поле избыточно т к закрывает во всех случаях админ, но их может быть несколько
						    Для истории лучше сохранять
		closing_comment	коментарий админа при закрытии жалобы
						*** Возможно поле избыточно, но было бы интересно чем в итоге закончилась жалоба
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
	groupp_name VARCHAR(32),
	adm_uuid UUID,
	closing_comment TEXT,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);
CREATE INDEX IF NOT EXISTS idx_created_at ON reports (created_at);

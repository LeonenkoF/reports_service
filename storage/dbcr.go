package storage

// запросы создания таблиц сервица
var (

	// report - таблица жалоб
	/*
		id				порядковый номер жалобы
		uuid 			идентификация жалобы (по ТЗ)
		user_uuid 		индентификатор автора-пользователя
		description 	текст жалобы
		priority		приоритет [Высокий, Средний, Низкий]
		stage			стадия обработки.
						*** Может быть enum но более гибко завести справочник под редакцией админа. Возможно пополнение стадий.
		topic			тема жалобы.
						*** Рекомендации те же чтои у stage
		group 			группа жалобы. может отсутсвовать
						*** Варианты:
						1. Каждой группе своя таблица с номерами жалоб, как и в случае хранения диалога по жалобе (comments)
						2. Таблица одна с парами (группа, имя группы, жалоба). Тогда будет дублирование данных по группам. Возможно легче обработка (поиск).

		adm_uuid		идентификатор админа, который закрывал жалобу и писал closing_comment
		closing_comment	коментарий админа при закрытии жалобы
						*** Возможно поле избыточно, но было бы интересно чем в итоге закончилась жалоба
	*/
	schReport string = `
CREATE TABLE IF NOT EXISTS reports (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	uuid UUID,
	user_uuid UUID,
	description TEXT,
	priority INT,
	stage INT,
	category INT,
	groupp UUID,
	groupp_name VARCHAR(32),
	adm_uuid UUID,
	closing_comment TEXT,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);
CREATE INDEX IF NOT EXISTS idx_created_at ON reports (created_at);
`
	// таблица группы (под каждую группу своя таблица с уникальным именем)
	// !!! %s имя уникально НЕ groupp !!!
	schGroup string = `
CREATE TABLE IF NOT EXISTS groupp (
	rep_id INT
)
`
	// или таблица групп (под все группы одна таблица)
	// *** недостаток - дублирование данных. преимущество - предположительно более легкий поиск
	schGroups string = `
CREATE TABLE IF NOT EXISTS groups (
	id INT,
	group_name VARCHAR(32)
	rep_id INT
)
CREATE INDEX IF NOT EXISTS idx_id ON groups (id); 
`
	// таблица диалога по жалобе. каждому диалогу своя таблица
	/*
			id 			идентификатор
			user_uuid 	автор
			comment 	собтвенно текст
		    created_at 	дата создания
		    updated_at 	дата редакции *** возможно лишнее поле
	*/
	// !!! %s имя уникально НЕ comments !!!
	schComments string = `
CREATE TABLE IF NOT EXISTS comments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_uuid UUID,
	comment TEXT,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,	
    updated_at DATE NOT NULL DEFAULT CURRENT_DATE,	
);
`
	// таблица пользователей
	/*
	   	id 			идентификатор
	   	user_uuid 	UUID,
	   	mail 		почта
	   	phone 		номер телефона (*** не обязательно)
	   	nick 		ник (*** необязательно)
	   	role INT	роль (user admin *** superuser *** superadmin)
	       created_at 	дата регистрации
	       updated_at 	дата инзменения (*** не обязательно)
	*/
	schUsers string = `
CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_uuid UUID,
	mail VARCHAR(64),
	phone VARCHAR(16),
	nick VARCHAR(16),
	role INT,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,	
    updated_at DATE NOT NULL DEFAULT CURRENT_DATE,	
);
`
	// список стадий обработки *** возможно нужно просто enum, но возможное появление новых стадий будет вызывать переписывание кода
	schStage string = `
CREATE TABLE IF NOT EXISTS stages (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	stage VARCHAR(10)
)	
`
	// список тем *** возможно нужно просто enum, но возможное появление новых тем будет вызывать переписывание кода
	schCategory string = `
CREATE TABLE IF NOT EXISTS categoryes (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	category VARCHAR(32)
)	
`

	// список ролей участников процесса *** возможно нужно просто enum, но возможное появление новых ролей будет вызывать переписывание кода
	schRole string = `
CREATE TABLE IF NOT EXISTS roles (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	role VARCHAR(10)
)	
`
	schPriority string = `
CREATE TABLE IF NOT EXISTS roles (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	priority VARCHAR(10)
)	
`
)
